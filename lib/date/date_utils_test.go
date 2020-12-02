package date

import "testing"

func TestToYYYYMMDD(t *testing.T) {
	result := len(ToYYYYMMDD())
	expect := 8
	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}
