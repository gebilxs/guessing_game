package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// version 1.0
//func main() {
//	maxNum := 100
//	secret := rand.Intn(maxNum)
//	fmt.Println("the secret number is ", secret)
//}

//version 1.5
//func main() {
//	maxNum := 100
//	rand.Seed(time.Now().UnixNano()) //利用时间戳进行更改
//	secret := rand.Intn(maxNum)
//	fmt.Println("the secret number is ", secret)
//}

//version 2.0
//func main() {
//	maxNum := 100
//	rand.Seed(time.Now().UnixNano())
//	secret := rand.Intn(maxNum)
//	fmt.Println("the secret number is ", secret)
//
//	fmt.Println("please input your guess")
//	reader := bufio.NewReader(os.Stdin) //利用这个方法把他修改成一个只读的流
//	input, err := reader.ReadString('\n')
//	if err != nil {
//		fmt.Println("error occurred while reading", err)
//		return
//	}
//	input = strings.TrimSuffix(input, "\n") //去掉换行符
//
//	guess, err := strconv.Atoi(input) //转换成数字  strconv 还有其他的转化过程
//	if err != nil {
//		fmt.Println("invalid input", err)
//		return
//	}
//	fmt.Println("your guess is", guess)
//}

//version2.5
//func main() {
//	maxNum := 100
//	rand.Seed(time.Now().UnixNano())
//	secretNumber := rand.Intn(maxNum)
//	fmt.Println("The secret number is ", secretNumber)
//
//	fmt.Println("Please input your guess")
//	reader := bufio.NewReader(os.Stdin)
//	input, err := reader.ReadString('\n')
//	if err != nil {
//		fmt.Println("An error occured while reading input. Please try again", err)
//		return
//	}
//	input = strings.TrimSuffix(input, "\n")
//
//	guess, err := strconv.Atoi(input)
//	if err != nil {
//		fmt.Println("Invalid input. Please enter an integer value")
//		return
//	}
//	fmt.Println("You guess is", guess)
//	if guess > secretNumber {
//		fmt.Println("Your guess is bigger than the secret number. Please try again")
//	} else if guess < secretNumber {
//		fmt.Println("Your guess is smaller than the secret number. Please try again")
//	} else {
//		fmt.Println("Correct, you Legend!")
//	}
//}

//version 3

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\r\n") //以上都要添加这种情况，程序可以继续运行
		// 由于windows平台的\r\n 和linux平台的\n不同
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value", err)
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
