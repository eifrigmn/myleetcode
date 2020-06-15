package _0007

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReverse(t *testing.T) {
	Convey("reverse", t, func() {
		num := 1534236469
		num = 120
		fmt.Println("@@@@@@@@", reverse(num))
	})
}
