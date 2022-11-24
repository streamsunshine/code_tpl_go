package question

import (
	"errors"
	"fmt"
	"github.com/gookit/goutil/testutil/assert"
	"math/rand"
	"testing"
)

/*
 * 快排
 * golang 的快排。
 */

func QuickSortRecur(list []int, start int, end int) {
	if start >= end {
		return
	}
	fmt.Printf("%v\t%v\n", start, end)
	for _, v := range list[start : end+1] {
		fmt.Printf("%v\t", v)
	}
	fmt.Printf("\n")

	cmpValue := list[end]
	lessIndex := start
	moreIndex := end - 1
	for lessIndex < moreIndex {
		if list[lessIndex] < cmpValue {
			lessIndex += 1
		} else if list[moreIndex] >= cmpValue {
			moreIndex -= 1
		} else {
			list[lessIndex], list[moreIndex] = list[moreIndex], list[lessIndex]
		}
	}
	if list[lessIndex] > cmpValue {
		list[end], list[lessIndex] = list[lessIndex], cmpValue
	} else {
		list[end], list[lessIndex+1] = list[lessIndex+1], cmpValue
		lessIndex = lessIndex + 1
	}
	QuickSortRecur(list, start, lessIndex-1)
	QuickSortRecur(list, lessIndex+1, end)
}

//参考 go 源码的写法
func QuickSortRecur2(data []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := hi
	a := lo
	c := hi - 1
	for {
		//这里必须用 hi，这样能保证 a 一定是大于等于 pivot。否则他会一直运动到 hi
		for ; a < hi && data[a] <= data[pivot]; a++ {
		}
		for ; a < c && data[c] > data[pivot]; c-- {
		}
		if a >= c {
			break
		}
		data[a], data[c] = data[c], data[a]
		a++
		c--
	}
	data[a], data[pivot] = data[pivot], data[a]
	QuickSortRecur2(data, lo, a-1)
	QuickSortRecur2(data, a+1, hi)
}

func QuickSort(list []int) {
	QuickSortRecur2(list, 0, len(list)-1)
}

func TestQuickSort(t *testing.T) {
	//testArr := []int{2, 3, 4, 7, 6, 5}
	testArr := []int{7, 6, 5, 4, 3, 4, 3, 2, 1}
	for i := 0; i < 1000; i++ {
		rand.Shuffle(7, func(i, j int) {
			testArr[i], testArr[j] = testArr[j], testArr[i]
		})
		//for _, v := range testArr {
		//	fmt.Printf("%v\t", v)
		//}
		//fmt.Printf("\n")
		QuickSort(testArr)
		assert.Eq(t, []int{1, 2, 3, 3, 4, 4, 5, 6, 7}, testArr)
		//for _, v := range testArr {
		//	fmt.Printf("%v\t", v)
		//}
	}

	//sort.Ints()
}

/*
 * heap
 * golang 源码倾向于 for 循环 +  break 的方式。
 */
type Heap struct {
	data []int
}

func (h *Heap) Push(v int) {
	h.data = append(h.data, v)
	dataLen := len(h.data)
	for i := dataLen - 1; i > 0; {
		//关键点-确保父亲计算正确
		p := (i - 1) / 2
		if p >= 0 && h.data[p] > h.data[i] {
			h.data[p], h.data[i] = h.data[i], h.data[p]
		} else {
			break
		}
		i = p
	}
}

func (h *Heap) Pop() (int, error) {
	dataLen := len(h.data)
	if dataLen < 1 {
		return 0, errors.New("empty heap")
	}
	rs := h.data[0]
	h.data[0] = h.data[dataLen-1]
	h.data = h.data[:dataLen-1]
	dataLen -= 1
	for i := 0; i < dataLen; {
		//确保儿子判断正确
		son := 2*i + 1
		rSon := son + 1
		if rSon < dataLen && h.data[rSon] < h.data[son] {
			son = rSon
		}
		if son < dataLen && h.data[son] < h.data[i] {
			h.data[son], h.data[i] = h.data[i], h.data[son]
		} else {
			break
		}
	}

	return rs, nil
}

func TestHeap(t *testing.T) {
	myHeap := MyHeap{
		list: make([]int, 0),
	}
	for i := 0; i < 10; i++ {
		myHeap.Push(10 - i)
	}
	for i := 0; i < 10; i++ {
		rs := myHeap.Pop()
		fmt.Printf("%v\n", rs)
	}
	//heap.Push()
}

/*
 * 二分查找
 */
func findX(data []int, target int) int {
	start := 0
	end := len(data) - 1
	for start < end {
		mid := (start + end) / 2

		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func TestFindX(t *testing.T) {
	arr := []int{1, 4, 7, 8}
	rs := findX(arr, 17)
	fmt.Printf("rs:%v\n", rs)
}
