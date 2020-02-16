package array_test

import "testing"

func TestArry(t *testing.T) {
	//var source = []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	//slice := source[1:2:4]
	//t.Log(slice)
	//t.Log(len(slice))
	//t.Log(cap(slice))


	var a = make([]int, 5, 6)
	t.Log("长度", len(a))
	t.Log("容量", cap(a))
	t.Log("值", a)
	a = []int{1, 2, 3, 4, 5}
	t.Log("赋值后")
	t.Log("长度", len(a))
	t.Log("容量", cap(a))
	t.Log("值", a)
	t.Log("append后")
	a = append(a, 6)
	t.Log("长度", len(a))
	t.Log("容量", cap(a))
	t.Log("值", a)
	b := []int{1,2,3,4,5,6,7,8}
	t.Log(len(b))
	t.Log(cap(b))
	b = append(b, 9)
	t.Log(len(b))
	t.Log(cap(b))
	var c [5]int
	t.Log(len(c))
	t.Log(cap(c))
	t.Log(c)
	var ttt []*int
	for k := range b{
		b[k] = 0
	}
	t.Log(ttt)
}
