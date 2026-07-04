# EuApI entity test

require "minitest/autorun"
require "json"
require_relative "../DutchCustomerData_sdk"
require_relative "runner"

class EuApIEntityTest < Minitest::Test
  def test_create_instance
    testsdk = DutchCustomerDataSDK.test(nil, nil)
    ent = testsdk.EuApI(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = eu_ap_i_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "eu_ap_i." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    eu_ap_i_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.eu_ap_i")))
    eu_ap_i_ref01_data = nil
    if eu_ap_i_ref01_data_raw.length > 0
      eu_ap_i_ref01_data = Helpers.to_map(eu_ap_i_ref01_data_raw[0][1])
    end

    # LIST
    eu_ap_i_ref01_ent = client.EuApI(nil)
    eu_ap_i_ref01_match = {}

    eu_ap_i_ref01_list_result = eu_ap_i_ref01_ent.list(eu_ap_i_ref01_match, nil)
    assert eu_ap_i_ref01_list_result.is_a?(Array)

    # LOAD
    eu_ap_i_ref01_match_dt0 = {
      "id" => eu_ap_i_ref01_data["id"],
    }
    eu_ap_i_ref01_data_dt0_loaded = eu_ap_i_ref01_ent.load(eu_ap_i_ref01_match_dt0, nil)
    eu_ap_i_ref01_data_dt0_load_result = Helpers.to_map(eu_ap_i_ref01_data_dt0_loaded)
    assert !eu_ap_i_ref01_data_dt0_load_result.nil?
    assert_equal eu_ap_i_ref01_data_dt0_load_result["id"], eu_ap_i_ref01_data["id"]

  end
end

def eu_ap_i_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "eu_ap_i", "EuApITestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = DutchCustomerDataSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["eu_ap_i01", "eu_ap_i02", "eu_ap_i03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID" => idmap,
    "DUTCHCUSTOMERDATA_TEST_LIVE" => "FALSE",
    "DUTCHCUSTOMERDATA_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["DUTCHCUSTOMERDATA_TEST_EU_AP_I_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["DUTCHCUSTOMERDATA_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
      },
      extra || {},
    ])
    client = DutchCustomerDataSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["DUTCHCUSTOMERDATA_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["DUTCHCUSTOMERDATA_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
