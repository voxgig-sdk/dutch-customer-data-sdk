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

func TestEuApIEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EuApI(nil)
		if ent == nil {
			t.Fatal("expected non-nil EuApIEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := eu_ap_iBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "eu_ap_i." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		euApIRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.eu_ap_i", setup.data)))
		var euApIRef01Data map[string]any
		if len(euApIRef01DataRaw) > 0 {
			euApIRef01Data = core.ToMapAny(euApIRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = euApIRef01Data

		// LIST
		euApIRef01Ent := client.EuApI(nil)
		euApIRef01Match := map[string]any{}

		euApIRef01ListResult, err := euApIRef01Ent.List(euApIRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, euApIRef01ListOk := euApIRef01ListResult.([]any)
		if !euApIRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", euApIRef01ListResult)
		}

		// LOAD
		euApIRef01MatchDt0 := map[string]any{
			"id": euApIRef01Data["id"],
		}
		euApIRef01DataDt0Loaded, err := euApIRef01Ent.Load(euApIRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		euApIRef01DataDt0LoadResult := core.ToMapAny(euApIRef01DataDt0Loaded)
		if euApIRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if euApIRef01DataDt0LoadResult["id"] != euApIRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func eu_ap_iBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "eu_ap_i", "EuApITestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read eu_ap_i test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse eu_ap_i test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"eu_ap_i01", "eu_ap_i02", "eu_ap_i03"},
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
	entidEnvRaw := os.Getenv("DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID": idmap,
		"DUTCHCUSTOMERDATA_TEST_LIVE":      "FALSE",
		"DUTCHCUSTOMERDATA_TEST_EXPLAIN":   "FALSE",
		"DUTCHCUSTOMERDATA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["DUTCHCUSTOMERDATA_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["DUTCHCUSTOMERDATA_APIKEY"],
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
