load 'test_helper/bats-support/load'
load 'test_helper/bats-assert/load'

export AMC_USER="artemis"
export AMC_PASSWORD="artemis"
export amc="./amc"

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

@test "send/receive binary file" {
    run dd if=/dev/urandom of=./test.bin bs=1024 count=1
    assert_success

    run bash -c '$amc put queue1 < ./test.bin'
    assert_success

    run bash -c '$amc get queue1 > ./test.out.bin'
    assert_success

    run diff ./test.bin ./test.out.bin
    assert_success
}
