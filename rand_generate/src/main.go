/* 主程序主要完成以下工作：
 * 	a) 获取并解析命令行输入；
 *  b）调用对应随机数生成器，生成对应数量的随机数；
 *  c) 将生成的随机数写到对应的文件中
 *  d) 打印生成对应数量的随机数所花费的时间
 */
package main

import (
	//	"bufio"
	"flag"
	"fmt"
	//	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var numcount *int = flag.Int("n", 1000, "The number of Needing generate.")
var outfile *string = flag.String("o", "outfile", "File to receive generated values.")
var method *string = flag.String("m", "method", "random methods.")

func main() {
	/* 解析参数 */
	flag.Parse()
	if numcount != nil {
		fmt.Println("numcount = ", *numcount, "outfile = ", *outfile, "method = ", *method)
	}

	t1 := time.Now()
	switch *method {

	default:
		fmt.Println("Random method ", *method, "is either unknown or unsupported.")
	}

	if *numcount > 0 {
		values := make([]int, *numcount)
		for i := 0; i < *numcount; i++ {
			values[i] = rand_generator_1()
		}
		writeValues(values, *outfile)

	} else {
		fmt.Println("需要产生的数据数量小于等于零")
	}

	t2 := time.Now()
	fmt.Println("The generating process costs", t2.Sub(t1), "to complete.")

}

/*
 * 将一个数组切片的内容写到文件中
 */
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\r\n")
	}
	return nil
}

/*
 *非并发： 函数 rand_generator_1 ，返回 int
 */
func rand_generator_1() int {
	return rand.Int()
}

/*
 * 单路并发：函数 rand_generator_2，返回 通道(Channel)
 */
func rand_generator_2() chan int {
	// 创建通道
	out := make(chan int)
	// 创建协程
	go func() {
		for {
			//向通道内写入数据，如果无人读取会等待
			out <- rand.Int()
		}
	}()
	return out
}
