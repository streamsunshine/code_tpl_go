package question

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"github.com/gookit/properties"
	"golang.org/x/tools/container/intsets"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 * 链表, 退出循环
 *
 * 当两个条件满足一个就退出的条件如何写；
 * 如何避免处理链表头指针为空。
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	node := head //tail  更合适
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{
			Val: carry % 10,
		}
		carry = carry / 10
		node = node.Next
	}
	if carry != 0 {
		node.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 * 连续截取的才算。
 * 字符串遍历，剪枝
 *
 * 我这里的方案的区别，利用了 map 记录位置，实现了左侧可以跳跃，不必一个一个挪动
 *
 * 学到了如何去遍历连续字串，并基于此对问题进行分析
 */
// 连续的字串，遍历子串可以把头或者尾依次固定，然后想优化方案。
// 这里先固定头部，移动尾部，直到遇到重复字串。此时考虑是否可以直接将头移动？假设有更好的，那么一定会有重复
// 因此可以移动头。而重复性判断采用  map 来做

// 官方的答案依次将 head 向后移动一下，然后每次继续移动右侧，直到遇到重复（开始移动左）
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLen := 0
	posMap := make(map[byte]int, 0)
	for index, v := range []byte(s) {
		if pos, exist := posMap[v]; exist {
			if pos+1 > start {
				start = pos + 1
			}
			posMap[v] = index
		} else {
			posMap[v] = index
		}
		//fmt.Printf("in:%v, pos:%v, v:%v\n", index, start, v)
		curLen := index - start + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

func TestLengthOfLongestSubstrin(t *testing.T) {
	rs := lengthOfLongestSubstring("abba")
	//rs := lengthOfLongestSubstring("abcabcbb")
	fmt.Printf("rs:%v\n", rs)
}

// @lc code=end

// 155. 最小栈  s
// 能在常数时间内检索到最小元素的栈
// 通过多维护一个栈，实现常数时间的返回。而最小堆要 o(logn) 的复杂度。

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{intsets.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	//if this.minStack[len(this.minStack)-1] < val {
	//	this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	//} else {
	//	this.minStack = append(this.minStack, val)
	//}
	this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func (this *MinStack) Pop() {
	if len(this.stack) <= 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// @lc code=end 。
//20. 有效的括号 s
// 栈
// 维护数组（栈), 左括号入栈；右括号出栈（判空），最后判断栈是否为空。   能早点发现不匹配。
// 或者反之，从栈的角度出发，空就填，非空就判断新的符号是否匹配，匹配就缩，否则就长。最后判空
func isValid(s string) bool {
	//这里用反向的，每次看前一个是不是当前的左括号。如果不是直接返回就行。
	pairMap := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	byteStack := make([]byte, 0, len(s))
	for _, char := range s {

		if len(byteStack) == 0 {
			byteStack = append(byteStack, byte(char))
			continue
		} else if v, exist := pairMap[byteStack[len(byteStack)-1]]; exist && v == byte(char) {
			byteStack = byteStack[:len(byteStack)-1]
		} else {
			byteStack = append(byteStack, byte(char))
		}
	}
	if len(byteStack) != 0 {
		return false
	}
	return true
}

func TestIsValid(t *testing.T) {
	rs := isValid("([])")
	fmt.Printf("rs:%v", rs)
}

//739. 每日温度  m
//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现
//在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
// 单调栈（构造一个单点递减的栈，当不单调递减时属于特殊处理的逻辑）
//
//除了变量命名，实现上和答案差异不大
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures)) //单调栈 noHigherIndexArr，这里的长度可以设置为 101
	for index, temperature := range temperatures {
		//从当前栈的栈顶开始
		stackIndex := len(stack) - 1
		//如果栈顶元素小于当前元素，则向前遍历，刷新天数，因为已经找到了更大的就可以移除了。最后停止在没找大更大的 index 处
		for ; stackIndex >= 0 && temperatures[stack[stackIndex]] < temperature; stackIndex-- {
			ans[stack[stackIndex]] = index - stack[stackIndex]
			stack = stack[:stackIndex]
		}
		//将当前元素下标保存起来
		stack = append(stack, index)
	}
	return ans
}

func TestValue(t *testing.T) {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	rs := dailyTemperatures(arr)
	fmt.Printf("rs:%+v", rs)
}

// copy 官方的方法
//func dailyTemperatures(temperatures []int) []int {
//	length := len(temperatures)
//	ans := make([]int, length)
//	next := make([]int, 101)
//	for i := 0; i < 101; i++ {
//		next[i] = math.MaxInt32
//	}
//	for i := length - 1; i >= 0; i-- {
//		warmerIndex := math.MaxInt32
//		for t := temperatures[i] + 1; t <= 100; t++ {
//			if next[t] < warmerIndex {
//				warmerIndex = next[t]
//			}
//		}
//		if warmerIndex < math.MaxInt32 {
//			ans[i] = warmerIndex - i
//		}
//		next[temperatures[i]] = i
//	}
//	return ans
//}
//

// 8 字符串转整数
// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
// 状态机
//
// 参考官方解法：https://leetcode.cn/problems/string-to-integer-atoi/    参考中并没有太关注非法数据的问题。
func myAtoi(s string) int {
	lenS := len(s)
	num := 0
	isNegative := false

	i := 0
	for ; i < lenS; i++ {
		if s[i] == ' ' {
			continue
		}

		if s[i] == '-' {
			isNegative = true
			i++
		} else if s[i] == '+' {
			i++
		}
		break
	}

	for ; i < lenS; i++ {

		posNum := int(s[i] - '0')
		if posNum < 0 || posNum > 9 {
			break
		}
		newNum := num*10 + posNum
		if isNegative && newNum > 1<<31 {
			return -(1 << 31)
		} else if !isNegative && newNum > 1<<31-1 {
			return 1<<31 - 1
		}
		num = newNum
	}

	if isNegative {
		num = -num
	}
	return num
}

func TestMyAtoi(t *testing.T) {
	rs := myAtoi("   -12345fajfk")
	fmt.Printf("rs:%+v", rs)
}

//  12  整数转罗马
// 比较基础，考察点不多
//
// 与答案中的一个解法类似
func intToRoman(num int) string {
	//这里要用数组
	numUnitMapRomanStr := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := make([]int, 0, len(numUnitMapRomanStr))
	for numUnit := range numUnitMapRomanStr {
		keys = append(keys, numUnit)
	}
	fmt.Printf("%+v", keys)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	rs := ""
	for _, value := range keys {
		numUnit, romanStr := value, numUnitMapRomanStr[value]
		for num-numUnit >= 0 {
			num = num - numUnit
			rs += romanStr
		}
	}
	return rs
}

func TestIntToRoman(t *testing.T) {
	rs := intToRoman(58)
	fmt.Printf("rs:%+v", rs)
}

//15 三数之和
//排序跳过防重，双指针法
//
//遍历数组，取差，然后双指针凑这个差值
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	fmt.Printf("list:%+v\n", nums)

	rs := make([][]int, 0)
	numsLen := len(nums)
	for first := 0; first < numsLen; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := numsLen - 1
		diff := 0 - nums[first]
		for second := first + 1; second < third; {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			sum := nums[third] + nums[second]
			if sum < diff {
				second++
			} else if sum > diff {
				third--
			} else if diff == sum {
				rs = append(rs, []int{nums[first], nums[second], nums[third]})
				second++
			}
		}
	}
	return rs
}

func TestThreeSum(t *testing.T) {
	rs := threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4})
	fmt.Printf("rs:%+v", rs)
}

// 面试真题。 利用 waitgroup，顺序打印字符
func TestCircle(t *testing.T) {
	chanA, chanB, chanC := make(chan struct{}), make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			<-chanA
			fmt.Println("a")
			chanB <- struct{}{}
		}
		<-chanA
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			<-chanB
			fmt.Println("b")
			chanC <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			<-chanC
			fmt.Println("c")
			chanA <- struct{}{}
		}
	}()
	chanA <- struct{}{}

	wg.Wait()
}

func TestSequenceCircle(t *testing.T) {
	letterNum := 26
	chanList := make([]chan struct{}, letterNum)
	for i := 0; i < letterNum; i++ {
		chanList[i] = make(chan struct{})
	}

	for i := 0; i < letterNum; i++ {
		go func(seqNo int) {
			for j := 0; j < 2; j++ {
				<-chanList[seqNo]
				fmt.Printf("%c\n", 'A'+seqNo)
				chanList[(seqNo+1)%letterNum] <- struct{}{}
			}
		}(i)
	}
	chanList[0] <- struct{}{}
	time.Sleep(1 * time.Second)
}

// 17 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 回溯，递归
//与答案的区别：递归可以不截断字符串，而是传入一个 index
var numStrMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	return recurLetterCombinations("", digits)
}

func recurLetterCombinations(pref string, digits string) []string {
	if len(digits) < 1 {
		return []string{pref}
	}

	digit := string(digits[0])
	num, _ := strconv.ParseInt(digit, 10, 32)
	numStr := numStrMap[int(num)]
	//fmt.Printf("num:%v, numStr:%v\n", num, numStr)

	rs := []string{}
	for i := 0; i < len(numStr); i++ {
		tmpRs := recurLetterCombinations(pref+string(numStr[i]), digits[1:])
		//fmt.Printf("pref:%v, tmpRs:%v\n", pref, tmpRs)
		rs = append(rs, tmpRs...)
	}
	return rs
}

func TestLetterCombinations(t *testing.T) {
	rs := letterCombinations("2")
	fmt.Printf("rs :%+v", rs)
}

// 测试结构体定义，指针结构体和直接定义的区别
//type Person struct {
//	age int
//}
//
//func (p *Person) howOld() int {
//	return p.age
//}
//
//func (p *Person) growUp() {
//	p.age += 1
//}
//
//func TestValuexxx1(t *testing.T) {
//	// qcrao 是值类型
//	qcrao := Person{age: 18}
//
//	// 值类型 调用接收者也是值类型的方法
//	fmt.Println(qcrao.howOld())
//
//	// 值类型 调用接收者是指针类型的方法
//	qcrao.growUp()
//	fmt.Println(qcrao.howOld())
//
//	// ----------------------
//
//	// stefno 是指针类型
//	stefno := &Person{age: 100}
//
//	// 指针类型 调用接收者是值类型的方法
//	fmt.Println(stefno.howOld())
//
//	// 指针类型 调用接收者也是指针类型的方法
//	stefno.growUp()
//	fmt.Println(stefno.howOld())
//}

// 18 四数之和
// 排序跳过去重，双指针法
//
//做的有点复杂，还是应该想三数之和那样，写两个循环，然后用双指针法。降低复杂度的方法就是双指针，其余就看几层循环了。
//目前的做法，复杂度不满足要求。
//不要因为某个小困难就放弃整个方案。
var elemList = make([]int, 4)
var rsList = make([][]int, 0)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	//fmt.Printf("nums:%v", nums)

	elemList = make([]int, 4)
	rsList = make([][]int, 0)

	NumSumRecur(nums, 0, target, 0)
	return rsList
}

func NumSumRecur(nums []int, start int, target int, count int) {
	if count > 3 {
		//fmt.Printf("return elemList:%v,start:%v,target:%v,count:%v\n", elemList, start, target, count)
		return
	}
	//fmt.Printf("elemList:%v,start:%v,target:%v,count:%v\n", elemList, start, target, count)

	numLen := len(nums)
	for index := start; index < numLen; index++ {
		//fmt.Printf("for elemList:%v,index:%v,start:%v,count:%v\n", elemList, index, start, count)

		if index > start && nums[index] == nums[index-1] {
			//fmt.Println("1")
			continue
		}
		elemList[count] = nums[index]
		//if nums[index] > target {
		//	fmt.Println("2")
		//	break
		//} else
		if count == 3 && nums[index] == target {
			//fmt.Println("3")

			rsList = append(rsList, []int{elemList[0], elemList[1], elemList[2], elemList[3]})
			//fmt.Printf("rs:%v\n", rsList)

			return
		} else if count != 3 {
			//fmt.Println("4")
			NumSumRecur(nums, index+1, target-nums[index], count+1)
		} else {
			//fmt.Println("5")

		}
	}
	return
}

func TestValueFourSUm(t *testing.T) {
	rs := fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11)
	fmt.Printf("rs:%v\n", rs)
}

/**
 * 19 删除链表的第 N 个节点
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 * 哑节点（不用特殊处理头），双指针
 *
 *
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	end := head
	count := 0
	for ; end != nil && count < n; count++ {
		end = end.Next
	}
	start := head
	for end != nil && end.Next != nil {
		end = end.Next
		start = start.Next
	}
	if end != nil {
		start.Next = start.Next.Next
	} else if n == count {
		return head.Next
	}
	return head
}

func TestRemoveNth(t *testing.T) {
	var next *ListNode
	for i := 0; i < 3; i++ {
		tmp := &ListNode{
			Val:  i,
			Next: next,
		}
		next = tmp
	}
	head := next
	index := head

	rs := removeNthFromEnd(head, 4)

	index = rs
	for index != nil {
		fmt.Printf("var:%v\n", index.Val)
		index = index.Next
	}

}

/*
 * @lc app=leetcode.cn id=22 lang=golang
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合
 * [22] 括号生成
 * 遍历，递归，剪枝
 *
 * 对于前后添加括号的方式，不满足  (())(())  这样就不行。
 * 要先转换成一个 2^2n 的问题， 然后考虑裁剪，比如 ( 和 ) 都是 n 个
 */

// @lc code=start

var list []string

func generateParenthesis(n int) []string {
	list = []string{}
	src := make([]byte, 2*n+1)
	recurGenerateParenthesis(&src, 0, 0, n)
	return list
}

func recurGenerateParenthesis(currStr *[]byte, leftNum, rightNum, n int) {
	if rightNum > leftNum {
		return
	}
	diff := leftNum + rightNum - 2*n
	if diff == 0 && leftNum == rightNum {
		list = append(list, string((*currStr)[:leftNum+rightNum]))
		return
	} else if diff > 0 {
		return
	}
	//fmt.Printf("currStr:%v,leftNum:%v,rightNum:%v, n:%v\n", currStr, leftNum, rightNum, n)

	(*currStr)[leftNum+rightNum] = ')'
	recurGenerateParenthesis(currStr, leftNum, rightNum+1, n)
	(*currStr)[leftNum+rightNum] = '('
	recurGenerateParenthesis(currStr, leftNum+1, rightNum, n)
}

func TestGenerateParenthesis(t *testing.T) {
	rs := generateParenthesis(3)
	for index, str := range rs {
		fmt.Printf("index:%d, str:%s\n", index, str)
	}
}

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 * 链表
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rs := head.Next
	lastTail := &ListNode{}
	for head != nil && head.Next != nil {
		next := head.Next

		head.Next = next.Next
		next.Next = head
		//拼接
		lastTail.Next = next
		//更新
		lastTail = head
		head = head.Next
	}
	return rs
}

func TestSwapPairs(t *testing.T) {
	nodeList := InitNodeList(3)
	PrintNodeList(nodeList)

	fmt.Println("-----------")

	nodeList = swapPairs(nodeList)
	PrintNodeList(nodeList)

}

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */
// @lc code=start
// 参考解法，积累 srcDivisor 的 2,4,8 倍数到一个 list，通过减法判断出 list 的 index 作为结果中对应 bit 位是否应该填充 1。
// 都要关注越界的最终值。
// 存在的问题，对于 go 的异或和取反都是 ^ 不了解。对于计算机

func divide(srcDividend int, srcDivisor int) int {
	dividend := uint32(srcDividend)
	divisor := uint32(srcDivisor)

	isNeg := false
	if srcDividend < 0 {
		isNeg = !isNeg
		dividend = uint32(-srcDividend)
	}
	if srcDivisor < 0 {
		isNeg = !isNeg
		divisor = uint32(-srcDivisor)
	}

	if srcDividend == math.MinInt32 {
		dividend = uint32(srcDividend)

		if srcDivisor == -1 {
			return math.MaxInt32
		} else if srcDivisor == 1 {
			return math.MinInt32
		}
	}
	if srcDivisor == 0 {
		return 0
	}

	topBit := uint32(0x80000000) //获取最高位
	value := uint32(0)
	rs := uint32(0)
	indexBit := topBit
	for ; indexBit > 0; indexBit = indexBit >> 1 {
		rs = rs << 1
		value = value << 1
		if dividend&indexBit > 0 {
			value = value + 1
		}

		if value < divisor {
			continue
		}

		fmt.Printf("rs:%b\n", rs)
		fmt.Printf("value:%b\n", value)
		rs = rs + 1
		value = value - divisor
	}
	fmt.Printf("rs:%x\n", rs)

	if isNeg {
		return -int(rs)
	}
	return int(rs)
}

func TestDivide(t *testing.T) {
	value := -2147483648
	//fmt.Printf("dev:%d\n", value)
	//fmt.Printf("dev:%d\n", uint32(value))

	rs := divide(value, 2)
	fmt.Printf("rs:%v\n", rs)
}

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 * 例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
 * 规律
 */

// @lc code=start
//基本想出来思路了。倒序遍历，找到第一个*降序*，和后面比他大的第一个数交换（不用再向前判断了）。由于后面都是升序，所以反转后面的数组
func nextPermutation(nums []int) {
	numsLen := len(nums)
	if numsLen <= 0 {
		return
	}
	lastNum := nums[numsLen-1]
	convertStartIndex := 0
	for i := numsLen - 2; i >= 0; i-- {
		if nums[i] < lastNum {
			for j := numsLen - 1; j >= 0; j-- {
				if nums[j] > nums[i] {
					tmpValue := nums[i]
					nums[i] = nums[j]
					nums[j] = tmpValue
					convertStartIndex = i + 1
					break
				}
			}
			break
		} else {
			lastNum = nums[i]
		}
	}
	for i := 0; convertStartIndex+i < numsLen-1-i; i++ {
		tmpValue := nums[numsLen-1-i]
		nums[numsLen-1-i] = nums[convertStartIndex+i]
		nums[convertStartIndex+i] = tmpValue
	}
}

//
//func nextPermutation(nums []int) {
//	n := len(nums)
//	i := n - 2
//	for i >= 0 && nums[i] >= nums[i+1] {
//		i--
//	}
//	if i >= 0 {
//		j := n - 1
//		for j >= 0 && nums[i] >= nums[j] {
//			j--
//		}
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//	reverse(nums[i+1:])
//}
//func reverse(a []int) {
//	for i, n := 0, len(a); i < n/2; i++ {
//		a[i], a[n-1-i] = a[n-1-i], a[i]
//	}
//}

func TestNextPermutation(t *testing.T) {
	//nums := []int{1, 2, 4, 5, 6}
	//nums := []int{6, 5, 4, 3, 2, 1}
	nums := []int{2, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 * [ 4 5 1 2 3 ]
 * 扩展，考虑一个上线左右的数组，如何查询。 n+m 的复杂度。从左下角开始移动，大了向上，小了向右
 */

// @lc code=start
// 没搞清楚到底如果安排 if else。 关键点在于中间的 index 能将数组划分为有序和无序的两个部分。
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		//fmt.Printf("start:%v, end:%v\n", start, end)
		tmp := (start + end) / 2
		value := nums[tmp]
		if value == target {
			return tmp
		}
		if value <= nums[end] {
			if target > value && target <= nums[end] {
				start = tmp + 1
				continue
			}
		} else {
			if target < nums[start] || target > value {
				start = tmp + 1
				continue
			}
		}
		end = tmp - 1
	}
	return -1
}

func TestSearch(t *testing.T) {
	nums := []int{3, 5, 1}

	for _, target := range nums {
		rs := search(nums, target)
		fmt.Printf("target: %v,rs:%v\n", target, rs)

	}
	target := 10
	rs := search(nums, target)
	fmt.Printf("target: %v,rs:%v\n", target, rs)
}

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
// 思考各个分支条件。如果一个思路不行，要灵活转换
// 要有使用动态规划的意识
// 即使看会了状态转移方程，实际代码写出来还是很困难，第一次花了一天，后面再理解一下。
func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	matchRecord := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		matchRecord[i] = make([]bool, pLen+1)
	}
	matchRecord[0][0] = true
	for j := 2; j <= pLen; j++ {
		if p[j-1] == '*' {
			matchRecord[0][j] = matchRecord[0][j-2]
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] == '*' {
				isMatch := false
				if j-1 >= 0 {
					isMatch = matchRecord[i][j-2]
				}

				if isMatchByte(s[i-1], p[j-2]) {
					isMatch = isMatch || matchRecord[i-1][j]
				}
				matchRecord[i][j] = isMatch
			} else if isMatchByte(s[i-1], p[j-1]) {
				matchRecord[i][j] = matchRecord[i-1][j-1]
			}
		}
	}
	for i := 0; i <= sLen; i++ {
		for j := 0; j <= pLen; j++ {
			fmt.Printf("%v\t", matchRecord[i][j])
		}
		fmt.Printf("\n")
	}
	return matchRecord[sLen][pLen]
}

func isMatchByte(src byte, dest byte) bool {
	isMatch := false
	if dest == '.' {
		isMatch = true
	} else if dest == src {
		isMatch = true
	}
	return isMatch
}

func TestIsMatch(t *testing.T) {
	//src := "a"
	//dest := ".*.a*"
	src := "a"
	dest := ".*c"
	//src := "abcd"
	//dest := "d*"
	rs := isMatch(src, dest)
	fmt.Printf("rs:%v\n", rs)
}

// @lc code=end
/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

//@lc code=start
//答案好简单，差距不是一点半点啊
//思路是获取第一个等于 target 以及第一个大于 target 的 index -1 。
//关于 二分法的条件判断，只要处理好
/*
n	n+1
target target
target  大
小      target
三种情况，而不至于死循环即可
*/

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 1 {
		return []int{-1, -1}
	}

	start := 0
	end := numsLen - 1
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}

	forward := start
	end = numsLen - 1
	for start < end {
		mid := (start + end + 1) >> 1
		if nums[mid] <= target {
			start = mid
		} else {
			end = mid - 1
		}
	}
	if nums[start] != target {
		return []int{-1, -1}
	}

	return []int{forward, start}
}

//func searchRange(nums []int, target int) []int {
//	leftmost := sort.SearchInts(nums, target)
//	if leftmost == len(nums) || nums[leftmost] != target {
//		return []int{-1, -1}
//	}
//	rightmost := MySearchInts(nums, target+1) - 1
//	return []int{leftmost, rightmost}
//}
//
//func MySearchInts(nums []int, target int) int {
//	numsLen := len(nums)
//	start := 0
//	end := numsLen //如果这里用 numsLen - 1 就会有导致判断失败。
//
//	for start < end {
//		mid := (start + end) >> 1
//		if nums[mid] >= target {
//			end = mid
//		} else {
//			start = mid + 1
//		}
//	}
//	return start
//}

func TestSearchRange(t *testing.T) {
	//nums := []int{1, 4, 4, 4, 4, 5}
	//nums := []int{5, 7, 7, 8, 8, 10}

	nums := []int{2, 2}
	target := 2
	rs := searchRange(nums, target)
	fmt.Printf("searchRange:%v", rs)
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
//还是想复杂了。可以通过一次遍历实现。九宫格的按照 i/3 j/3 来统计就可以了。
func isValidSudoku(board [][]byte) bool {
	hori := isValidSudokuInner(board, transIJHori) &&
		isValidSudokuInner(board, transIJVeri) &&
		isValidSudokuInner(board, transIJBlock)

	return hori
}
func isValidSudokuInner(board [][]byte, tranIJFunc func(int, int) (int, int)) bool {
	isValid := true
	for i := 0; i < 9; i++ {
		flagMap := make(map[byte]struct{}, 9)
		for j := 0; j < 9; j++ {
			k, z := tranIJFunc(i, j)
			char := board[k][z]
			if char == '.' {
				continue
			}
			if _, exist := flagMap[char]; exist {
				isValid = false
				break
			}
			flagMap[char] = struct{}{}
		}
		if isValid == false {
			break
		}
	}
	return isValid
}

func transIJHori(i, j int) (int, int) {
	return i, j
}

func transIJVeri(i, j int) (int, int) {
	return j, i
}

func transIJBlock(i, j int) (int, int) {
	tmpJ := j%3 + i/3*3
	i = i%3*3 + j/3
	return i, tmpJ
}

func TestIsValidSudoku(t *testing.T) {
	//testSudoku := [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	testSudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '8', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	rs := isValidSudoku(testSudoku)
	fmt.Printf("judgeRs:%v\n", rs)
	//for i := 0; i < 9; i++ {
	//	for j := 0; j < 9; j++ {
	//		tmpI, tmpJ := transIJBlock(i, j)
	//		fmt.Printf("i:%v,j:%v,tmpI:%v,tmpJ:%v\n", i, j, tmpI, tmpJ)
	//	}
	//}
}

// @lc code=end

type BloomFilter struct {
	bitMap       []byte
	hashFuncList []func(value string) int
	rwLock       *sync.RWMutex
	len          int
}

func (flt *BloomFilter) Init(mapLen int, funcList ...func(value string) int) {
	flt.bitMap = make([]byte, mapLen/8+1)
	flt.hashFuncList = funcList
	flt.rwLock = &sync.RWMutex{}
	flt.len = mapLen
}

// Contains
// 纠结了一下 defer 写在 lock 前还是后
func (flt *BloomFilter) CheckExist(value string) bool {
	for _, hashFunc := range flt.hashFuncList {
		hashValue := hashFunc(value)

		isExist := true
		func() {
			flt.rwLock.RLock()
			defer flt.rwLock.RUnlock()
			hashValue %= flt.len                                      //有遗漏
			exist := flt.bitMap[hashValue/8] & (1 << (hashValue % 8)) //定义单独变量要考虑空间占用
			if exist == 0 {
				isExist = false
			}
		}()
		if !isExist {
			return false
		}
	}
	return true
}

// Add
func (flt *BloomFilter) SetBloomValue(value string) {
	for _, hashFunc := range flt.hashFuncList {
		hashValue := hashFunc(value)
		hashValue %= flt.len
		func() {
			flt.rwLock.Lock()
			defer flt.rwLock.Unlock()
			flt.bitMap[hashValue/8] |= 1 << hashValue % 8
		}()
	}
}

func TestBloomFilter(t *testing.T) {
	filter := &BloomFilter{}
	filter.Init(10, func(value string) int { return len(value) })
	str := ""
	wg := sync.WaitGroup{}
	for i := 0; i < 11; i++ {
		str += "1"
		wg.Add(1) //这个不能放在 go 里面，因为不能保证他执行到了
		go func(str string) {
			defer wg.Done()
			isExist := filter.CheckExist(str)
			fmt.Printf("value:%s, rs:%v\n", str, isExist)
			if !isExist {
				filter.SetBloomValue(str)
			}
		}(str)
	}
	wg.Wait()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 哨兵节点
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	index := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			index.Next = list1
			list1 = list1.Next
		} else {
			index.Next = list2
			list2 = list2.Next
		}
		index = index.Next
	}
	index.Next = list1
	if list2 != nil {
		index.Next = list2
	}
	return head.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//耗时确实比较严重，而且逻辑也比较复杂
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	prev := head

	index := 0
	list := &ListNode{}
	for len(lists) > 0 {
		nextIndex := 0
		for index, list = range lists {
			if list == nil {
				break
			}
			if lists[nextIndex].Val > list.Val {
				nextIndex = index
			}
		}
		if list == nil {
			copy(lists[index:], lists[index+1:]) //这里的 copy 耗时
		}
		prev.Next = lists[nextIndex]
		lists[nextIndex] = lists[nextIndex].Next
		prev = prev.Next
	}
	return head.Next
}

//代码简单
func mergeKLists2(lists []*ListNode) *ListNode {
	var ans *ListNode
	for _, list := range lists {
		ans = mergeTwoLists(ans, list)
	}
	return ans
}

//归并， logk 的提升
func mergeKLists3(lists []*ListNode) *ListNode {
	return mergeKListsRecur(lists, 0, len(lists)-1)
}

//采用递归的方式，两两合并
func mergeKListsRecur(lists []*ListNode, l, r int) *ListNode {
	//这里的两个条件关注一下
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(mergeKListsRecur(lists, l, mid), mergeKListsRecur(lists, mid+1, r))
}

type ListNodeMinQueue []*ListNode

func (l *ListNodeMinQueue) Less(i, j int) bool { return (*l)[i].Val < (*l)[j].Val }
func (l *ListNodeMinQueue) Len() int           { return len(*l) }
func (l *ListNodeMinQueue) Swap(i, j int)      { (*l)[i], (*l)[j] = (*l)[j], (*l)[i] }
func (l *ListNodeMinQueue) Push(i interface{}) { *l = append(*l, i.(*ListNode)) }
func (l *ListNodeMinQueue) Pop() interface{} {
	len := l.Len()
	tmp := (*l)[len-1]
	*l = (*l)[:len-1]
	return tmp
}

//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/solution/hua-dong-chuang-kou-de-zui-da-zhi-by-lee-ymyo/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

//采用优先级队列的方式
func mergeLists4(lists []*ListNode) *ListNode {
	minHeap := &ListNodeMinQueue{}
	for _, list := range lists {
		if list == nil {
			continue
		}
		heap.Push(minHeap, list)
	}

	//fmt.Printf("len:%d\n", minHeap.Len())
	head := &ListNode{}
	prev := head
	for minHeap.Len() > 0 {
		prev.Next = heap.Pop(minHeap).(*ListNode)
		if prev.Next.Next != nil {
			heap.Push(minHeap, prev.Next.Next)
		}
		prev = prev.Next
	}
	return head.Next
}

func TestMergeLists(t *testing.T) {
	lists := make([]*ListNode, 0, 3)
	for i := 0; i < 3; i++ {
		listHead := &ListNode{}
		list := listHead
		for j := 0; j < 5; j++ {
			tmp := &ListNode{
				Val: j + int(math.Pow(10, float64(i))),
			}
			list.Next = tmp
			list = list.Next
		}
		lists = append(lists, listHead.Next)
	}

	rs := mergeLists4(lists)

	index := rs
	for index != nil {
		fmt.Printf("var:%v\n", index.Val)
		index = index.Next
	}

}

var cmbList [][]int
var tmpList []int

func combinationSum(candidates []int, target int) [][]int {
	cmbList = make([][]int, 0)
	tmpList = make([]int, 0)
	combinationSumRecur(candidates, 0, target)
	return cmbList
}

func combinationSumRecur(candidates []int, index int, target int) {
	//fmt.Printf("index:%v, target:%v\n", index, target)
	if target == 0 {
		list := append([]int(nil), tmpList...)
		cmbList = append(cmbList, list)
		return
	} else if target < 0 || index >= len(candidates) {
		return
	}

	combinationSumRecur(candidates, index+1, target)

	tmpList = append(tmpList, candidates[index])
	combinationSumRecur(candidates, index, target-candidates[index])
	tmpList = tmpList[:len(tmpList)-1]
}

func TestCombinationSum(t *testing.T) {
	rs := combinationSum([]int{2, 3, 6, 7}, 8)
	fmt.Printf("rs:%v\n", rs)
}

func findInt(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (start + end + 1) / 2
		if nums[mid] == target {
			start = mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if nums[start] == target {
		return start
	}
	return -1
}

func TestFindInt(t *testing.T) {
	testArr := []int{-1, 2, 6, 6, 6, 7, 7}
	rs := findInt(testArr, 6)
	fmt.Printf("rs:%v\n", rs)
}

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	for index, num := range nums {
		//第一个特殊处理
		if index == 0 {
			ans = append(ans, []int{num})
			continue
		}
		//通过对少一个元素的进行插入，获取当前的全排列
		ansLen := len(ans)
		for j := 0; j < ansLen; j++ {
			list := ans[0]
			listLen := len(list)
			for i := 0; i < listLen+1; i++ {
				newList := append([]int{num}, list...)
				newList[i], newList[0] = newList[0], newList[i]
				ans = append(ans, newList)
			}
			//已经获取了少一个元素的数组，能产生的全部全排列，可以将其踢出
			ans = ans[1:]
		}
	}
	return ans
}

func TestPermute(t *testing.T) {
	rs := permute([]int{1, 2, 3})
	fmt.Printf("rs:%v\n", rs)
}

// @lc code=end

type Average struct {
	Sum  int //平均值
	Cnt  int //计数
	Lock *sync.RWMutex
}

func GetObject() *Average {
	return &Average{
		Lock: &sync.RWMutex{},
	}
}

func (a *Average) AddValue(value int) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	a.Sum += value
	a.Cnt += 1
}

func (a *Average) GetValue() int {
	a.Lock.RLock()
	defer a.Lock.RUnlock()
	return a.Sum / a.Cnt
}

type StaticsDataInterface interface {
	InputValue(value int)
	GetResult() int
}

type StaticsData struct {
	objMap map[string]StaticsDataInterface
}

func (s *StaticsData) RegisterStaticsObj(objName string, obj StaticsDataInterface) {
	s.objMap[objName] = obj
}

func (s *StaticsData) InputValue(value int) {
	for _, obj := range s.objMap {
		obj.InputValue(value)
	}
}

func (s *StaticsData) GetResult(objName string) {

}

func Triple(nums []int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)
	rs := [][]int{}
	for index, num := range nums {
		start := index + 1
		end := numsLen - 1

		if index > 0 && num == nums[index-1] {
			continue
		}

		cmpValue := 0 - num
		for start < end {
			if nums[start] == nums[start-1] {
				start++
				continue
			}
			curValue := nums[start] + nums[end]
			if curValue == cmpValue {
				rs = append(rs, []int{nums[index], nums[start], nums[end]})
				start++
				end--
			} else if curValue > cmpValue {
				end--
			} else {
				start++
			}
		}
	}
	return rs
}

func TestTriple(t *testing.T) {
	nums := []int{-2, 0, 0, 2, 2}
	rs := Triple(nums)
	fmt.Printf("rs:%v\n", rs)
}

type MyHeap struct {
	list []int
}

func (h *MyHeap) Push(elem int) {
	h.list = append(h.list, elem)
	pos := len(h.list) - 1
	for pos > 0 {
		parentPos := (pos - 1) / 2
		if h.list[parentPos] > h.list[pos] {
			h.list[parentPos], h.list[pos] = h.list[pos], h.list[parentPos]
		}
		//fmt.Printf("pos:%v\n", pos)
		pos = parentPos
	}
	//for _, value := range h.list {
	//	fmt.Printf("%v\t", value)
	//}
}

func (h *MyHeap) Pop() int {
	listLen := len(h.list)
	if listLen <= 0 {
		return math.MaxInt
	}

	elem := h.list[0]
	h.list[0] = h.list[listLen-1]
	listLen = listLen - 1
	h.list = h.list[:listLen]

	pos := 0
	for pos < listLen {
		childPos := 2*pos + 1
		if childPos < listLen && h.list[childPos] < h.list[pos] {
			h.list[childPos], h.list[pos] = h.list[pos], h.list[childPos]
		}
		if childPos+1 < listLen && h.list[childPos+1] < h.list[pos] {
			childPos = childPos + 1
			h.list[childPos], h.list[pos] = h.list[pos], h.list[childPos]
		}
		pos = childPos
	}

	return elem
}

func TestHead(t *testing.T) {
	myHeap := MyHeap{
		list: make([]int, 0),
	}
	for i := 0; i < 10; i++ {
		myHeap.Push(10 - i)
	}
	for i := 0; i < 1; i++ {
		rs := myHeap.Pop()
		fmt.Printf("%v\n", rs)
	}
}

func GetTopK(nums []int, kNum int) []int {
	countMap := make(map[int]int, 0)
	for _, num := range nums {
		countMap[num]++
	}

	type KV struct {
		K int
		V int
	}
	kvList := make([]KV, 0, len(countMap))
	for k, v := range countMap {
		kvList = append(kvList, KV{K: k, V: v})
	}
	sort.Slice(kvList, func(i, j int) bool {
		return kvList[i].V > kvList[j].V
	})

	kvLen := len(kvList)
	if kNum > kvLen {
		kNum = kvLen
	}
	rs := make([]int, 0, kNum)
	i := 0
	for ; i < kNum; i++ {
		rs = append(rs, kvList[i].K)
	}
	for j := i; j < kvLen && kvList[i-1].V == kvList[j].V; j++ {
		rs = append(rs, kvList[j].K)
	}

	return rs
}

func TestGetTopK(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 6, 7}
	rs := GetTopK(nums, 3)
	fmt.Printf("rs:%v\n", rs)
}

//
//func TestGetAb(t *testing.T) {
//	n, m := 0, 0
//	_, _ = fmt.Scanf("%d, %d", &n, &m)
//
//	userList := make([]Pq, n)
//	for id, v := range userList {
//		v.Id = id
//	}
//
//	totalLine := n * (n - 1) / 2
//	for i := 0; i <= totalLine; i++ {
//		x, y, a, b := 0, 0, 0, 0
//		_, _ = fmt.Scanf("%d, %d, %d, %d", &x, &y, &a, &b)
//
//		userList[x-1].Q += a
//		userList[y-1].Q += n
//		if a > b {
//			userList[x-1].P += 1
//		} else {
//			userList[x-1].P += 1
//		}
//	}
//
//	sort.Slice(userList, func(i, j int) bool {
//		if userList[i].P > userList[j].P {
//			return true
//		} else if userList[i].P == userList[j].P && userList[i].Q > userList[i].Q {
//			return true
//		}
//		return false
//	})
//
//	rs := make(map[int]struct{}, n)
//	for index, v := range userList {
//		if index < m {
//			rs[v.Id] = struct{}{}
//		} else {
//			break
//		}
//	}
//
//	for i := 0; i < n; i++ {
//		num := 0
//		if _, exist := rs[i]; exist {
//			num = 1
//		}
//		fmt.Printf("%d", num)
//	}
//}

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	lastStr := countAndSay(n - 1)
	rs := make([]byte, 0)

	lastBytes := []byte(lastStr)
	count := 0
	for index, char := range lastBytes {
		if index != 0 && char != lastBytes[index-1] {
			rs = append(rs, byte('0'+count), lastBytes[index-1])
			count = 1
		} else {
			count += 1
		}
	}
	rs = append(rs, byte('0'+count), lastBytes[len(lastBytes)-1])
	//fmt.Printf("n:%d, rs:%v\n", n, string(rs))
	return string(rs)
}

// @lc code=end

func TestCountAndSay(t *testing.T) {
	rs := countAndSay(5)
	fmt.Printf("rs: %v\n", rs)
}

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 * 1、对重复的 i 要特别处理，增加一个计数，避免 1 2 3 .. 个的重复运算。  -- 可以在通用的基础上优化
 * 2、排序：由于数值是递增的，可以用于剪枝；同时也方面上面的统计;
 * 3、顺序处理的加上上一步的处理，可以不用 map 去重。
 */

// @lc code=start

var rsMap map[string][]int

func combinationSum2(candidates []int, target int) [][]int {
	rsMap = make(map[string][]int)
	sort.Ints(candidates)
	initList := make([]int, 0, 10)
	combinationSum2Recur(initList, candidates, target)

	rs := make([][]int, 0, len(rsMap))
	for _, elem := range rsMap {
		rs = append(rs, elem)
	}
	return rs
}

func combinationSum2Recur(list []int, candidates []int, target int) {
	if len(candidates) == 0 {
		return
	}

	val := candidates[0]

	count := 0
	for _, cand := range candidates {
		if cand != val {
			break
		}
		count++
	}
	candidates = candidates[count:]

	combinationSum2Recur(list, candidates, target)

	for i := 0; i < count; i++ {
		list = append(list, val)
		target -= val
		if target == 0 {
			//fmt.Printf("val:%v, list:%v\n", val, list)
			var str strings.Builder
			for _, v := range list {
				str.WriteString(fmt.Sprintf("%d_", v))
			}
			rsMap[str.String()] = append([]int{}, list...)
			return
		} else if target < 0 {
			return
		}
		combinationSum2Recur(list, candidates, target)
	}
}

// @lc code=end

func TestCombinationSum2(t *testing.T) {
	//list := []int{10, 1, 2, 7, 6, 1, 5}
	list := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target := 27

	rs := combinationSum2(list, target)
	for _, list := range rs {
		fmt.Println(list)
	}
}

func TestCombinationSum211(t *testing.T) {
	type A struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	var testList []*A
	testStr := `[{"a": 1, "b": "www"}, {"a": 2, "b": "wwwa"}]`
	err := json.Unmarshal([]byte(testStr), &testList)
	fmt.Printf("list:%v, err:%v", testList, err)
}

func TestAbc1(t *testing.T) {
	a = []int{1, 2, 3}
	b := a
	b[1] = 100
	b = append(b, 4)
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func TestAtomic(t *testing.T) {
	n := int32(0)
	atomic.AddInt32(&n, 1)
	fmt.Printf("rs:%v\n", n)
}

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 * 在此次跳跃评估下次跳跃最远能到哪里。我的解法空间和时间占用偏大。
 * 用贪心法从头依次遍历。 记录当前点 到 其跳到最远处 之间的点，跳到最远是多少。作为下次跳跃的重点。
 * 不必担心最远的点有 0 ，因为遍历了中间所有的点，要跳不过去，谁也过不去。
 */

// @lc code=start
func jump(nums []int) int {
	numsLen := len(nums)

	minStepMap := make(map[int]int, numsLen)

	for i := numsLen - 2; i >= 0; i-- {
		jumpStep := nums[i]

		if i+jumpStep >= numsLen-1 {
			minStepMap[i] = 1
			continue
		}

		stepToEnd := numsLen
		for j := 1; j <= jumpStep; j++ {
			nextStepIndex := j + i
			if stepToEnd > 1+minStepMap[nextStepIndex] {
				stepToEnd = 1 + minStepMap[nextStepIndex]
			}
		}
		minStepMap[i] = stepToEnd
	}
	//for k, v := range minStepMap {
	//	fmt.Printf("k:%v, v:%v\n", k, v)
	//}
	return minStepMap[0]
}

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	rs := jump(nums)
	fmt.Printf("rs :%v\n", rs)
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 * 有两个关键；1、排序，不让相邻的两个相同元素出现在同一个位置即可。2、用给数组的去填空，通过 map 记录已经使用的元素
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	rs := make([][]int, 0)
	numsLen := len(nums)
	numsUsed := make(map[int]struct{}, 0)

	var recurFunc func(oneList []int, fillIdx int)

	recurFunc = func(oneList []int, fillIdx int) {
		if fillIdx == numsLen {
			tmp := append([]int{}, oneList...)
			rs = append(rs, tmp)
			return
		}

		lastIndex := -1
		for index := 0; index < numsLen; index++ {
			if _, exist := numsUsed[index]; exist {
				continue
			}

			if lastIndex >= 0 && nums[lastIndex] == nums[index] {
				continue
			}
			lastIndex = index
			numsUsed[index] = struct{}{}

			oneList[fillIdx] = nums[index]
			//fmt.Printf("list:%v, num:%v, fillId:%v, numsUsed:%v\n", oneList, nums[index], fillIdx, numsUsed)

			recurFunc(oneList, fillIdx+1)
			delete(numsUsed, index)
		}
	}

	oneList := make([]int, numsLen)
	recurFunc(oneList, 0)
	return rs
}

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	rs := permuteUnique(nums)
	fmt.Printf("rs:%v\n", rs)
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 * 规律推出来了，当时相比于答案的反转还是复杂了一些。  这里只有一条规律， 反转后，列变成行. 行于新列之和为 n-1.
 * 如果没有记住这个规律，先从最外层，然后向里层，观察 4 * 4 数组的规律即可。
 */

// @lc code=start
func rotate(matrix [][]int) {
	mLen := len(matrix)
	iterNum := mLen / 2

	for i := 0; i < iterNum; i++ {
		//第一次循环， 0，0
		for j := 0; j < mLen-1-2*i; j++ {
			tmp := matrix[i][i+j]
			matrix[i][i+j] = matrix[mLen-1-i-j][i]
			matrix[mLen-1-i-j][i] = matrix[mLen-1-i][mLen-1-i-j]
			matrix[mLen-1-i][mLen-1-i-j] = matrix[i+j][mLen-1-i]
			matrix[i+j][mLen-1-i] = tmp
			//tmp := matrix[i][i]
			//matrix[i][i] = matrix[mLen-1-i][i]
			//matrix[mLen-1-i][i] = matrix[mLen-1-i][mLen-1-i]
			//matrix[mLen-1-i][mLen-1-i] = matrix[i][mLen-1-i]
			//matrix[i][mLen-1-i] = tmp
		}
	}
}

func TestRotate(t *testing.T) {
	testMatrix := [][]int{
		{5, 1, 9, 11}, {2, 4, 8, 10},
		{13, 3, 6, 7}, {15, 14, 12, 16},
	}
	rotate(testMatrix)

	for _, v := range testMatrix {
		for _, val := range v {
			fmt.Printf("%v\t", val)
		}
		fmt.Printf("\n")
	}
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 * 也可以把字符出现的次数作为 key，因为都是小写字母，所以 26 个就够了。
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	strListMap := make(map[string][]string, 0)
	for _, v := range strs {
		byteList := []byte(v)
		sort.Slice(byteList, func(i, j int) bool {
			return byteList[i] < byteList[j]
		})
		str := string(byteList)
		if _, exist := strListMap[str]; exist {
			strListMap[str] = append(strListMap[str], v)
		} else {
			strListMap[str] = []string{v}
		}
	}
	rs := make([][]string, 0)
	for _, strList := range strListMap {
		rs = append(rs, strList)
	}
	return rs
}

func TestGroupAnagrams(t *testing.T) {
	testMatrix := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	rs := groupAnagrams(testMatrix)

	for _, v := range rs {
		for _, val := range v {
			fmt.Printf("%v\t", val)
		}
		fmt.Printf("\n")
	}
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 * 可以采用迭代的方式，移位是 1 就乘到结果中，知道 n 为 0
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	var myPowRecur func(float64, int) float64

	myPowRecur = func(f float64, n int) float64 {
		if n == 0 {
			return 1.0
		} else if n == 1 {
			return f
		}

		v := myPowRecur(f, n/2)
		v = v * v * myPowRecur(f, n%2)
		return v
	}

	if n < 0 {
		return 1 / myPowRecur(x, -n)
	} else {
		return myPowRecur(x, n)
	}
}

func TestMyPow(t *testing.T) {
	x := 2.0
	n := 5

	rs := myPow(x, n)
	fmt.Printf("%v\n", rs)

}

// @lc code=end

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 * 这里没想到用动态规划，考虑的是使用双指针法。
 * 这里还可以进一步优化，在 nums 上面改，不用额外空间。
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt
	}
	dp := make([]int, len(nums))
	maxV := nums[0]

	for idx, v := range nums {
		if idx == 0 {
			dp[0] = nums[0]
		} else if dp[idx-1] > 0 {
			dp[idx] = dp[idx-1] + v
		} else {
			dp[idx] = v
		}
		if dp[idx] > maxV {
			maxV = dp[idx]
		}
	}
	return maxV
}

// @lc code=end

func TestYaml(t *testing.T) {
	buff, err := ioutil.ReadFile("custom_config.yaml")
	if err != nil {
		fmt.Printf("read err:%v\n", err)
	}
	fmt.Printf("buff:%v\n", string(buff))
	data := make(map[string]interface{}, 0)
	err = yaml.Unmarshal(buff, &data)
	if err != nil {
		fmt.Printf("unmarshal err:%v\n", err)
	}

	propRs, err := properties.Marshal(data)
	fmt.Printf("data:%v\n", data)
	fmt.Printf("propRs:%v, err:%v\n", string(propRs), err)
}

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 * 这个虽然做出来了，但是逻辑过于复杂。推算的过程很繁琐。
 * 答案的第一个方法设置了一个 direction = [[0,1], [-1,0], [0,-1], [1,0] 的数组，并记录访问过的元素，如果遇到了已访问的就改变方向
 * 第二个方法，设置 top， left， right， bottom 的点位进行遍历，可读性要明显比我的好。尽管思路相同。
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	rowLen := len(matrix)
	if rowLen <= 0 {
		return []int{}
	}
	columnLen := len(matrix[0])
	minLen := columnLen
	if rowLen < columnLen {
		minLen = rowLen
	}

	rs := make([]int, 0, rowLen*columnLen)
	for i := 0; i < minLen/2+minLen%2; i++ {
		for j := i; j < columnLen-i; j++ {
			rs = append(rs, matrix[i][j])
		}
		for j := i + 1; j < rowLen-i; j++ {
			rs = append(rs, matrix[j][columnLen-i-1])
		}
		if minLen-i-1 > i {
			for j := columnLen - i - 2; j >= i; j-- {
				rs = append(rs, matrix[rowLen-i-1][j])
			}
			for j := rowLen - i - 2; j > i; j-- {
				rs = append(rs, matrix[j][i])
			}
		}

	}
	return rs
}

// @lc code=end
func TestSpiralOrder(t *testing.T) {
	testMatrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	//testMatrix := [][]int{
	//	{5, 1, 9},
	//	{2, 4, 8},
	//	{13, 3, 6},
	//}
	rs := spiralOrder(testMatrix)

	for _, v := range rs {
		fmt.Printf("%v\t", v)
	}
}
