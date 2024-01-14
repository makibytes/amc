load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

export amc="go run main.go -u artemis -p artemis"

@test "send and receive hello world" {
    run $amc put queue1 Hello World
    assert_success

    run $amc get queue1
    assert_success

    assert_output "Hello World"
}

@test "send Hello World2 from STDIN and receive" {
    run echo "Hello World2" | $amc put queue1
    assert_success

    run $amc get queue1
    assert_success

    assert_output "Hello World2"
}
