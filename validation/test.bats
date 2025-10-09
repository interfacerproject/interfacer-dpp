setup() {
  load bats/setup
}

@test "zenroom exists and is executable" {
  run which zenroom
  assert_success
}

@test "string dictionary to dpp (lua)" {
  run luaexec strdict_to_dpp.lua dpp_examples.json datatypes.json
  assert_success
  >&3 echo $output
}
