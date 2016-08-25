package foo

import "testing"

func TestEqualsFoo(t *testing.T) {
	if EqualsFoo("bar") {
		t.Fail()
	}
	if !EqualsFoo("foo") {
		t.Fail()
	}
}

func TestEqualsFooTable(t *testing.T) {
	tests := []struct {
		args   string
		result bool
	}{
		{
			args:   "bar",
			result: false,
		},
		{
			args:   "foo",
			result: true,
		},
	}

	for _, tt := range tests {
		if EqualsFoo(tt.args) != tt.result {
			t.Fail()
		}
	}
}
