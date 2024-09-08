load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

export amc="./amc -u artemis -p artemis"

@test "initialize ANYCAST address" {
    run $amc get queue1
    assert_success
}

@test "send/receive text payload" {
    run $amc put queue1 HelloWorld
    assert_success

    run $amc get queue1
    assert_success

    assert_output "HelloWorld"
}

@test "send HelloWorld from STDIN and receive" {
    run bash -c 'echo "HelloWorld" | $amc put queue1'
    assert_success

    run $amc get queue1
    assert_success

    assert_output "HelloWorld"
}
