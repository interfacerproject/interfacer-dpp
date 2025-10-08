setup() {
  load bats/setup
}

@test "zencode-exec exists and is executable" {
  run which zencode-exec
  assert_success
}

@test "string dictionary to dpp (lua)" {
  run luaexec strdict_to_dpp.lua example_drill.json
  assert_success
  >&3 echo $output
}
