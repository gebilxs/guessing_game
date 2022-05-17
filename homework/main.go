package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	//reader := bufio.NewReader(os.Stdin)
	var guess int
	for {
		//input, err := reader.ReadString('\n')
		//if err != nil {
		//	fmt.Println("An error occured while reading input. Please try again", err)
		//	continue
		//}
		//input = strings.TrimSuffix(input, "\r\n") //以上都要添加这种情况，程序可以继续运行
		//// 由于windows平台的\r\n 和linux平台的\n不同
		//guess, err := strconv.Atoi(input)

		_, err := fmt.Scanf("%d\n", &guess)
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
