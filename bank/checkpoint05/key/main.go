// >示例如下：
// ============================================
// 输入：
// 3
// 1 2 3
// 输出：
// [[1 2 3][1 3 2][2 1 3][2 3 1][3 1 2][3 2 1]]
// ============================================

// >代码模板:

// func permute(nums []int) [][]int {
//     // insert your code

// }
// func main() {
//     var n int
//         fmt.Scanf("%d", &n)

//         testSlice := make([]int, n)
//     // 标准输入n个不重复的数字

//	    res := permute(testSlice)
//	    fmt.Println(res)
//	}
//
// 嘻嘻嘻我不会嘻嘻嘻
package main

import (
	"fmt"
)

func permute(n int, nums []int) [][]int {
	// insert your code
	var permute [][]int
	permute = make([][]int)
	for i := 0; i < n*n-1; i++ {
		for j := 0; j < n; j++ {

		}
	}

}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	testSlice := make([]int, n)
	//标准输入n个不重复的数字
	for i := 0; i < n; i++ {
		var num int
		fmt.Println("依次输入不重复的n个数字")
		fmt.Scanln(&num)
		testSlice = append(testSlice, num)
	}
	res := permute(testSlice)
	fmt.Println(res)
}
