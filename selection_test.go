package selection

import (
	"reflect"
	"testing"
)
func Test_MinIdx(t *testing.T) {
	nums, expect := []int{4,1,3,2}, 1
	got := minIdx(nums)
	if got != expect {
		t.Errorf("error getting minimum index of slice, got: %v, expect: %v", got, expect)
	}
}

func Test_Sort(t *testing.T) {
	unsorted := []int{5,4,3,1,2}
	sorted := []int{1,2,3,4,5}

	if !reflect.DeepEqual(sorted, Sort(unsorted)) {
		t.Errorf("error sorting slice, got: %+v, expect: %+v", Sort(unsorted), sorted)
	}
}