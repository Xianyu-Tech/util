package encryptutil

import (
	"encoding/hex"
	"testing"
)

func TestMd5(t *testing.T) {
	type Input struct {
		Raw []byte
	}
	type Expect struct {
		Result string
	}
	type TestCase struct {
		Input  Input
		Expect Expect
	}

	tcs := []TestCase{
		{Input: Input{[]byte("x233")}, Expect: Expect{"8c8b09e7e420cfeb6f2875ba07de603b"}},
		{Input: Input{[]byte("lifetrip")}, Expect: Expect{"d7ec2cfef7ee7f86b86e4c2c9bf29652"}},
	}

	for _, v := range tcs {
		got := hex.EncodeToString(Md5(v.Input.Raw))
		if got != v.Expect.Result {
			t.Errorf("got=%s,expect=%s", got, v.Expect.Result)
		}
	}
}

func TestMd5String(t *testing.T) {
	type Input struct {
		Raw []byte
	}
	type Expect struct {
		Result string
	}
	type TestCase struct {
		Input  Input
		Expect Expect
	}

	tcs := []TestCase{
		{Input: Input{[]byte("x233")}, Expect: Expect{"8c8b09e7e420cfeb6f2875ba07de603b"}},
		{Input: Input{[]byte("lifetrip")}, Expect: Expect{"d7ec2cfef7ee7f86b86e4c2c9bf29652"}},
	}

	for _, v := range tcs {
		got := Md5Str(v.Input.Raw)
		if got != v.Expect.Result {
			t.Errorf("got=%s,expect=%s", got, v.Expect.Result)
		}
	}
}
