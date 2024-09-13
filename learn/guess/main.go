/*
 * @Author: matiastang
 * @Date: 2024-09-13 10:39:34
 * @LastEditors: matiastang
 * @LastEditTime: 2024-09-13 11:02:02
 * @FilePath: /matias-Golang/learn/guess/main.go
 * @Description: 猜数字游戏
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"strconv"
	"strings"
	"os"
	"log"
)

func main() {
	// 获取当前时间的秒级时间戳
	seccends := time.Now().Unix()
	// 使用当前时间的秒级时间戳作为随机数种子
	rand.Seed(seccends)
	// 生成1到100之间的随机数
	target := rand.Intn(100) + 1
	// 打印目标数字
	fmt.Printf("target:%d \n", target)
	// 提示用户输入范围
	fmt.Println("I have a number between 1 and 100.")
	// 提示用户输入
	fmt.Println("Can you guess it?")
	// 创建一个缓冲区读取器
	reader := bufio.NewReader(os.Stdin)
	success := false
	// 循环10次
	for i := 0; i < 5; i++ {
		fmt.Println("You have %d guesses left. Please type your guess and then press enter.", 10 - i)
		// 读取用户输入的内容
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// 去除输入字符串中的换行符
		input = strings.TrimSpace(input)
		// 将输入字符串转换为整数
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Too low.")
		} else if guess > target {
			fmt.Println("Too high.")
		} else {
			fmt.Println("You got it!")
			success = true
			break
		}	
	}
	if !success {
		fmt.Println("You've run out of guesses. The number was", target)
	}
}
