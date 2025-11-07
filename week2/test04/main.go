//  使用for循环生成20个goroutine,
// 每个goroutine随机休眠0~1000ms，
// 并向一个channel传入随机数和goroutine编号(从1-20)，
// 等待这些goroutine都生成完后，
// 想办法给这些随机数按照编号进行排序
// (输出排序前和排序后的结果,要求不使用额外的空间存储这20个数据)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Number struct {
	num int
	id  int
}

func main() {

	ch := make(chan struct {
		num int
		id  int
	}, 20)
	var Num Number
	// var wg sync.WaitGroup
	for i := 1; i <= 20; i++ {
		go func(i int) {
			//随机休眠0~1000ms
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
			Num.num = rand.Intn(100)
			Num.id = i
			ch <- Num
		}(i)
	}
	// wg.Wait()
	// close(ch)

	result := make([]struct {
		num int
		id  int
	}, 0, 20)
	fmt.Println("排序前：")
	for i := 0; i < 20; i++ {
		v := <-ch
		result = append(result, v)
		fmt.Printf("id:%v,num:%v\n", v.id, v.num)
	}

	//排序
	for i := 0; i < 20; i++ {
		for j := 0; j < 19-i; j++ {
			if result[j].id > result[j+1].id {
				temp := result[j]
				result[j] = result[j+1]
				result[j+1] = temp
			}

		}
	}

	fmt.Println("排序后：")
	for _, v := range result {
		fmt.Printf("id:%v,num:%v\n", v.id, v.num)
	}

}
