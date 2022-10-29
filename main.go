package main

import "fmt"

type Pq struct {
	P  int
	Q  int
	Id int
}

func main() {

	n := 0
	rs := make([]int, 0)
	_, _ = fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		n = 0
		_, _ = fmt.Scanf("%d", &n)

		numList := make([]int, 0, n)
		sum := 0
		for i := 0; i < n; i++ {
			var num int
			_, _ = fmt.Scan(&num)
			//fmt.Printf("test:%d", num)
			sum += num
			numList = append(numList, num)
		}

		_, _ = fmt.Scanf("%d", &n)
		//fmt.Printf("test:%d", n)

		index := 0
		finishTime := 0
		for i := 0; i < n; i++ {
			x, y := 0, 0
			_, _ = fmt.Scanf("%d %d", &x, &y)
			//fmt.Printf("test1:%d,%d\n", x, y)

			if x == 1 {
				numList = append(numList, y)
				sum += y
			} else if x == 2 {
				y %= sum
				for y > 0 {
					leftTime := numList[index] - finishTime
					if leftTime > y {
						finishTime = y
						y = 0
					} else if leftTime == y {
						index++
						finishTime = 0
						y = 0
					} else {
						index++
						y -= leftTime
						finishTime = 0
					}
					//fmt.Printf("www %d. %d\n", y, leftTime)
					index %= len(numList)
				}
				rs = append(rs, index+1)
			}
			//fmt.Printf("%d\n", index+1)
		}
	}
	for i := 0; i < len(rs); i++ {
		fmt.Printf("%d\n", rs[i])
		if i != len(rs)-1 {
			fmt.Printf("\n")
		}
	}
}

//
//func main() {
//	n, m := 0, 0
//	_, _ = fmt.Scanf("%d %d", &n, &m)
//
//	userList := make([]Pq, n)
//	for id, _ := range userList {
//		userList[id].Id = id
//	}
//
//	totalLine := n * (n - 1) / 2
//	for i := 0; i < totalLine; i++ {
//		x, y, a, b := 0, 0, 0, 0
//		_, _ = fmt.Scanf("%d %d %d %d", &x, &y, &a, &b)
//
//		userList[x-1].Q += a
//		userList[y-1].Q += b
//		if a > b {
//			userList[x-1].P += 1
//		} else {
//			userList[y-1].P += 1
//		}
//	}
//	//for _, v := range userList {
//	//	fmt.Printf("blist: %d, %d, %d\n", v.Id, v.P, v.Q)
//	//}
//
//	sort.Slice(userList, func(i, j int) bool {
//		if userList[i].P > userList[j].P {
//			return true
//		} else if userList[i].P == userList[j].P && userList[i].Q > userList[j].Q {
//			return true
//		}
//		return false
//	})
//
//	rs := make(map[int]struct{}, n)
//	for index, v := range userList {
//		//fmt.Printf("list: %d, %d, %d\n", v.Id, v.P, v.Q)
//		if index < m {
//			rs[v.Id] = struct{}{}
//		} else if index > 0 && v.P == userList[index-1].P && v.Q == userList[index-1].Q {
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
//		if i != n-1 {
//			fmt.Printf("\n")
//		}
//	}
//}
//
//func main() {
//	n := 0
//	_, _ = fmt.Scanf("%d", &n)
//
//	rList := make([]int, n)
//	for i := 0; i < n; i++ {
//		var num int
//		_, _ = fmt.Scan(num)
//		rList = append(rList, num)
//	}
//
//	sList := make([]int, n)
//	for i := 0; i < n; i++ {
//		var num int
//		_, _ = fmt.Scan(num)
//		sList = append(sList, num)
//	}
//
//}
