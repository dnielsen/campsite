package testUtil

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"testing"
)

// Compares values other than errors. Otherwise it might throw "cannot handle unexported field at {*errors.errorString}.:..."
func Cmp(t *testing.T, val1 interface{}, val2 interface{})  {
	if diff := cmp.Diff(val1, val2); diff != "" {
		t.Errorf("mismatch (-wantCode, +gotCode): \n%s", diff)
	}
}

// We must remember that the pointers to the want and got value must be the same.
// If pointers won't be the same, the test will fail.
func CmpErr(t *testing.T, val1 error, val2 error)  {
	if diff := cmp.Diff(val1, val2, cmp.Comparer(func(x, y error) bool {
		return errors.Is(x,y) || errors.Is(y, x)
	})); diff != "" {
		t.Errorf("mismatch (-wantCode, +gotCode): \n%s", diff)
	}
}
