package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/dutch-customer-data-sdk/go"
	"github.com/voxgig-sdk/dutch-customer-data-sdk/go/core"

	vs "github.com/voxgig-sdk/dutch-customer-data-sdk/go/utility/struct"
)

func TestGlobalApIEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.GlobalApI(nil)
		if ent == nil {
			t.Fatal("expected non-nil GlobalApIEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := global_ap_iBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "global_ap_i." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		globalApIRef01Ent := client.GlobalApI(nil)
		globalApIRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "global_ap_i"}, setup.data), "global_ap_i_ref01"))

		globalApIRef01DataResult, err := globalApIRef01Ent.Create(globalApIRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		globalApIRef01Data = core.ToMapAny(globalApIRef01DataResult)
		if globalApIRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// LIST
		globalApIRef01Match := map[string]any{}

		globalApIRef01ListResult, err := globalApIRef01Ent.List(globalApIRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		globalApIRef01List, globalApIRef01ListOk := globalApIRef01ListResult.([]any)
		if !globalApIRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", globalApIRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(globalApIRef01List), map[string]any{"id": globalApIRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// LOAD
		globalApIRef01MatchDt0 := map[string]any{}
		globalApIRef01DataDt0Loaded, err := globalApIRef01Ent.Load(globalApIRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if globalApIRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func global_ap_iBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "global_ap_i", "GlobalApITestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read global_ap_i test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse global_ap_i test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"global_ap_i01", "global_ap_i02", "global_ap_i03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID": idmap,
		"DUTCHCUSTOMERDATA_TEST_LIVE":      "FALSE",
		"DUTCHCUSTOMERDATA_TEST_EXPLAIN":   "FALSE",
	})

	idmapResolved := core.ToMapAny(env["DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["DUTCHCUSTOMERDATA_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
			},
			extra,
		})
		client = sdk.NewDutchCustomerDataSDK(core.ToMapAny(mergedOpts))
	}

	live := env["DUTCHCUSTOMERDATA_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["DUTCHCUSTOMERDATA_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
