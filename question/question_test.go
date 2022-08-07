package question

import (
	"container/heap"
	"fmt"
	"math"
	"sync"
	"testing"
)

//// 155. 最小栈  s
//type MinStack struct {
//	stack    []int
//	minStack []int
//}
//
//func Constructor() MinStack {
//	return MinStack{
//		stack:    []int{},
//		minStack: []int{intsets.MaxInt},
//	}
//}
//
//func (this *MinStack) Push(val int) {
//	this.stack = append(this.stack, val)
//	//if this.minStack[len(this.minStack)-1] < val {
//	//	this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
//	//} else {
//	//	this.minStack = append(this.minStack, val)
//	//}
//	this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//func (this *MinStack) Pop() {
//	if len(this.stack) <= 0 {
//		return
//	}
//	this.stack = this.stack[:len(this.stack)-1]
//	this.minStack = this.minStack[:len(this.minStack)-1]
//}
//
//func (this *MinStack) Top() int {
//	return this.stack[len(this.stack)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	return this.minStack[len(this.minStack)-1]
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

////20. 有效的括号 s
//func isValid(s string) bool {
//	//这里用反向的，每次看前一个是不是当前的左括号。如果不是直接返回就行。
//	pairMap := map[byte]byte{
//		'{': '}',
//		'[': ']',
//		'(': ')',
//	}
//	byteStack := make([]byte, 0, len(s))
//	for _, char := range s {
//
//		if len(byteStack) == 0 {
//			byteStack = append(byteStack, byte(char))
//			continue
//		} else if v, exist := pairMap[byteStack[len(byteStack)-1]]; exist && v == byte(char) {
//			byteStack = byteStack[:len(byteStack)-1]
//		} else {
//			byteStack = append(byteStack, byte(char))
//		}
//	}
//	if len(byteStack) != 0 {
//		return false
//	}
//	return true
//}
//
//func TestValue(t *testing.T) {
//	rs := isValid("([])")
//	fmt.Printf("rs:%v", rs)
//}

//739. 每日温度  m
//除了变量命名，实现上和答案差异不大
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	for index, temperature := range temperatures {
		stackIndex := len(stack) - 1
		for ; stackIndex >= 0 && temperatures[stack[stackIndex]] < temperature; stackIndex-- {
			ans[stack[stackIndex]] = index - stack[stackIndex]
			stack = stack[:stackIndex]
		}
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

//8 整数转罗马数字
//func myAtoi(s string) int {
//	lenS := len(s)
//	num := 0
//	isNegative := false
//
//	i := 0
//	for ; i < lenS; i++ {
//		if s[i] == ' ' {
//			continue
//		}
//
//		if s[i] == '-' {
//			isNegative = true
//			i++
//		} else if s[i] == '+' {
//			i++
//		}
//		break
//	}
//
//	for ; i < lenS; i++ {
//
//		posNum := int(s[i] - '0')
//		if posNum < 0 || posNum > 9 {
//			break
//		}
//		newNum := num*10 + posNum
//		if isNegative && newNum > 1<<31 {
//			return -(1 << 31)
//		} else if !isNegative && newNum > 1<<31-1 {
//			return 1<<31 - 1
//		}
//		num = newNum
//	}
//
//	if isNegative {
//		num = -num
//	}
//	return num
//}
//
//func TestValue1(t *testing.T) {
//	rs := myAtoi("   -12345fajfk")
//	fmt.Printf("rs:%+v", rs)
//}
//
////  12  整数转罗马
//func intToRoman(num int) string {
//	//这里要用数组
//	numUnitMapRomanStr := map[int]string{
//		1000: "M",
//		900:  "CM",
//		500:  "D",
//		400:  "CD",
//		100:  "C",
//		90:   "XC",
//		50:   "L",
//		40:   "XL",
//		10:   "X",
//		9:    "IX",
//		5:    "V",
//		4:    "IV",
//		1:    "I",
//	}
//	keys := make([]int, 0, len(numUnitMapRomanStr))
//	for numUnit := range numUnitMapRomanStr {
//		keys = append(keys, numUnit)
//	}
//	fmt.Printf("%+v", keys)
//	sort.Slice(keys, func(i, j int) bool {
//		return keys[i] > keys[j]
//	})
//	rs := ""
//	for _, value := range keys {
//		numUnit, romanStr := value, numUnitMapRomanStr[value]
//		for num-numUnit >= 0 {
//			num = num - numUnit
//			rs += romanStr
//		}
//	}
//	return rs
//}
//
//func TestValue1(t *testing.T) {
//	rs := intToRoman(58)
//	fmt.Printf("rs:%+v", rs)
//}

// 15 三数之和
//func threeSum(nums []int) [][]int {
//	if len(nums) < 3 {
//		return nil
//	}
//	sort.Ints(nums)
//	fmt.Printf("list:%+v\n", nums)
//
//	rs := make([][]int, 0)
//	numsLen := len(nums)
//	for first := 0; first < numsLen; first++ {
//		if first > 0 && nums[first] == nums[first-1] {
//			continue
//		}
//		third := numsLen - 1
//		diff := 0 - nums[first]
//		for second := first + 1; second < third; {
//			if second > first+1 && nums[second] == nums[second-1] {
//				second++
//				continue
//			}
//			sum := nums[third] + nums[second]
//			if sum < diff {
//				second++
//			} else if sum > diff {
//				third--
//			} else if diff == sum {
//				rs = append(rs, []int{nums[first], nums[second], nums[third]})
//				second++
//			}
//		}
//	}
//	return rs
//}
//
//func TestValue1(t *testing.T) {
//	rs := threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4})
//	fmt.Printf("rs:%+v", rs)
//}
//
//func TestCircle(t *testing.T) {
//
//	chanA, chanB, chanC := make(chan struct{}), make(chan struct{}), make(chan struct{})
//	wg := sync.WaitGroup{}
//	wg.Add(3)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 1; i++ {
//			<-chanA
//			fmt.Println("a")
//			chanB <- struct{}{}
//		}
//		<-chanA
//	}()
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 1; i++ {
//			<-chanB
//			fmt.Println("b")
//			chanC <- struct{}{}
//		}
//	}()
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 1; i++ {
//			<-chanC
//			fmt.Println("c")
//			chanA <- struct{}{}
//		}
//	}()
//	chanA <- struct{}{}
//
//	wg.Wait()
//}
//
//func TestSequenceCircle(t *testing.T) {
//	letterNum := 26
//	chanList := make([]chan struct{}, letterNum)
//	for i := 0; i < letterNum; i++ {
//		chanList[i] = make(chan struct{})
//	}
//
//	for i := 0; i < letterNum; i++ {
//		go func(seqNo int) {
//			for j := 0; j < 2; j++ {
//				<-chanList[seqNo]
//				fmt.Printf("%c\n", 'A'+seqNo)
//				chanList[(seqNo+1)%letterNum] <- struct{}{}
//			}
//		}(i)
//	}
//	chanList[0] <- struct{}{}
//	time.Sleep(1 * time.Second)
//}

//// 17 电话号码的字母组合
////递归可以不阶段字符串，而是传入一个 index
//var numStrMap = map[int]string{
//	2: "abc",
//	3: "def",
//	4: "ghi",
//	5: "jkl",
//	6: "mno",
//	7: "pqrs",
//	8: "tuv",
//	9: "wxyz",
//}
//
//func letterCombinations(digits string) []string {
//	if len(digits) < 1 {
//		return []string{}
//	}
//	return recurLetterCombinations("", digits)
//}
//
//func recurLetterCombinations(pref string, digits string) []string {
//	if len(digits) < 1 {
//		return []string{pref}
//	}
//
//	digit := string(digits[0])
//	num, _ := strconv.ParseInt(digit, 10, 32)
//	numStr := numStrMap[int(num)]
//	//fmt.Printf("num:%v, numStr:%v\n", num, numStr)
//
//	rs := []string{}
//	for i := 0; i < len(numStr); i++ {
//		tmpRs := recurLetterCombinations(pref+string(numStr[i]), digits[1:])
//		//fmt.Printf("pref:%v, tmpRs:%v\n", pref, tmpRs)
//		rs = append(rs, tmpRs...)
//	}
//	return rs
//}
//
//func TestValue17(t *testing.T) {
//	rs := letterCombinations("2")
//	fmt.Printf("rs :%+v", rs)
//}

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

////做的有点复杂，还是应该想三数之和那样，写两个循环，然后用双指针法
////目前的做法，复杂度不满足要求。
////不要因为某个小困难就放弃整个方案。
//var elemList = make([]int, 4)
//var rsList = make([][]int, 0)
//
//func fourSum(nums []int, target int) [][]int {
//	sort.Ints(nums)
//	//fmt.Printf("nums:%v", nums)
//
//	elemList = make([]int, 4)
//	rsList = make([][]int, 0)
//
//	NumSumRecur(nums, 0, target, 0)
//	return rsList
//}
//
//func NumSumRecur(nums []int, start int, target int, count int) {
//	if count > 3 {
//		//fmt.Printf("return elemList:%v,start:%v,target:%v,count:%v\n", elemList, start, target, count)
//		return
//	}
//	//fmt.Printf("elemList:%v,start:%v,target:%v,count:%v\n", elemList, start, target, count)
//
//	numLen := len(nums)
//	for index := start; index < numLen; index++ {
//		//fmt.Printf("for elemList:%v,index:%v,start:%v,count:%v\n", elemList, index, start, count)
//
//		if index > start && nums[index] == nums[index-1] {
//			//fmt.Println("1")
//			continue
//		}
//		elemList[count] = nums[index]
//		//if nums[index] > target {
//		//	fmt.Println("2")
//		//	break
//		//} else
//		if count == 3 && nums[index] == target {
//			//fmt.Println("3")
//
//			rsList = append(rsList, []int{elemList[0], elemList[1], elemList[2], elemList[3]})
//			//fmt.Printf("rs:%v\n", rsList)
//
//			return
//		} else if count != 3 {
//			//fmt.Println("4")
//			NumSumRecur(nums, index+1, target-nums[index], count+1)
//		} else {
//			//fmt.Println("5")
//
//		}
//	}
//	return
//}
//
//func TestValueFourSUm(t *testing.T) {
//	rs := fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11)
//	fmt.Printf("rs:%v\n", rs)
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	end := head
//	count := 0
//	for ; end != nil && count < n; count++ {
//		end = end.Next
//	}
//	start := head
//	for end != nil && end.Next != nil {
//		end = end.Next
//		start = start.Next
//	}
//	if end != nil {
//		start.Next = start.Next.Next
//	} else if n == count {
//		return head.Next
//	}
//	return head
//}
//
//func TestRemoveNth(t *testing.T) {
//	var next *ListNode
//	for i := 0; i < 3; i++ {
//		tmp := &ListNode{
//			Val:  i,
//			Next: next,
//		}
//		next = tmp
//	}
//	head := next
//	index := head
//
//	rs := removeNthFromEnd(head, 4)
//
//	index = rs
//	for index != nil {
//		fmt.Printf("var:%v\n", index.Val)
//		index = index.Next
//	}
//
//}

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
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
 */

// @lc code=start
//基本想出来思路了。倒序遍历，找到第一个降序，和后面比他大的第一个数交换（不用再向前判断了）。然后交换她左边的数组
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
		if index == 0 {
			ans = append(ans, []int{num})
			continue
		}
		ansLen := len(ans)
		for j := 0; j < ansLen; j++ {
			list := ans[0]
			listLen := len(list)
			for i := 0; i < listLen+1; i++ {
				newList := append([]int{num}, list...)
				newList[i], newList[0] = newList[0], newList[i]
				ans = append(ans, newList)
			}
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
