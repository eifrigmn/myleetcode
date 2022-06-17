package _018

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	Convey("fourSum", t, func() {
		result := [][]int{{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}}
		So( reflect.DeepEqual(result, fourSum([]int{1,0,-1,0,-2,2}, 0)), ShouldBeTrue)
		result = [][]int{{2,2,2,2}}
		So(reflect.DeepEqual(result, fourSum([]int{2,2,2,2}, 8)), ShouldBeTrue)
		result = [][]int{{2,2,2,2}}
		So(reflect.DeepEqual(result, fourSum([]int{2,2,2,2,2}, 8)), ShouldBeTrue)
	})
}
