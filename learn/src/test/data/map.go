/*
 * @Author: matiastang
 * @Date: 2022-04-13 18:43:03
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-13 19:19:53
 * @FilePath: /matias-Golang/learn/src/test/data/map.go
 * @Description: map
 */
package data

import "fmt"

func init() {
	fmt.Println("map init")
}

func MapTest() {
	// 定义 变量strMap
	var strMap map[int]string
	// 进行初始化
	strMap = make(map[int]string)

	// 给map 赋值
	for i := 0; i < 5; i++ {
		strMap[i] = "迈莫coding"
	}

	// 打印出map值
	for k, v := range strMap {
		fmt.Println(k, ":", v)
	}

	// 打印出map 长度
	fmt.Println(len(strMap))
}
