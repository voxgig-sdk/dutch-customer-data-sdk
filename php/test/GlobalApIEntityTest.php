<?php
declare(strict_types=1);

// GlobalApI entity test

require_once __DIR__ . '/../dutchcustomerdata_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class GlobalApIEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = DutchCustomerDataSDK::test(null, null);
        $ent = $testsdk->GlobalApI(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "global_ap_i" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = DutchCustomerDataSDK::test($seed, null);
        $seen = iterator_to_array($base->GlobalApI(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = DutchCustomerDataConfig::make_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = DutchCustomerDataSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->GlobalApI(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = global_ap_i_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "global_ap_i." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $global_ap_i_ref01_ent = $client->GlobalApI(null);
        $global_ap_i_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.global_ap_i"), "global_ap_i_ref01"));

        $global_ap_i_ref01_data_result = $global_ap_i_ref01_ent->create($global_ap_i_ref01_data, null);
        $global_ap_i_ref01_data = Helpers::to_map($global_ap_i_ref01_data_result);
        $this->assertNotNull($global_ap_i_ref01_data);

        // LIST
        $global_ap_i_ref01_match = [];

        $global_ap_i_ref01_list_result = $global_ap_i_ref01_ent->list($global_ap_i_ref01_match, null);
        $this->assertIsArray($global_ap_i_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($global_ap_i_ref01_list_result),
            ["id" => $global_ap_i_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // LOAD
        $global_ap_i_ref01_match_dt0 = [];
        $global_ap_i_ref01_data_dt0_loaded = $global_ap_i_ref01_ent->load($global_ap_i_ref01_match_dt0, null);
        $this->assertNotNull($global_ap_i_ref01_data_dt0_loaded);

    }
}

function global_ap_i_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/global_ap_i/GlobalApITestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = DutchCustomerDataSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["global_ap_i01", "global_ap_i02", "global_ap_i03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID" => $idmap,
        "DUTCHCUSTOMERDATA_TEST_LIVE" => "FALSE",
        "DUTCHCUSTOMERDATA_TEST_EXPLAIN" => "FALSE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["DUTCHCUSTOMERDATA_TEST_GLOBAL_AP_I_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["DUTCHCUSTOMERDATA_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
            ],
            $extra ?? [],
        ]);
        $client = new DutchCustomerDataSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["DUTCHCUSTOMERDATA_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["DUTCHCUSTOMERDATA_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
