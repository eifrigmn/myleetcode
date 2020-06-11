package _0155

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinStack(t *testing.T) {
	Convey("min stack", t, func() {
		// -2, 0, -3
		minStack := Constructor2()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		So(minStack.GetMin(), ShouldEqual, -3)
		minStack.Pop()
		So(minStack.Top(), ShouldEqual, 0)
		So(minStack.GetMin(), ShouldEqual, -2)
		// ["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
		// [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
		minStack = Constructor2()
		minStack.Push(2147483646)
		minStack.Push(2147483646)
		minStack.Push(2147483647)
		So(minStack.Top(), ShouldEqual, 2147483647)
		minStack.Pop()
		So(minStack.GetMin(), ShouldEqual, 2147483646)
		minStack.Pop()
		So(minStack.GetMin(), ShouldEqual, 2147483646)
		minStack.Pop()
		minStack.Push(2147483647)
		So(minStack.Top(), ShouldEqual, 2147483647)
		So(minStack.GetMin(), ShouldEqual, 2147483647)
		minStack.Push(-2147483648)
		So(minStack.Top(), ShouldEqual, -2147483648)
		So(minStack.GetMin(), ShouldEqual, -2147483648)
		minStack.Pop()
		So(minStack.GetMin(), ShouldEqual, 2147483647)
	})
}
