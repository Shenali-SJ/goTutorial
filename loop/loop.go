package loop

import "fmt"

func PrintNumbers() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		if i > 7 {
			break
		}

		fmt.Printf("%d ", i)
	}
}

func ShowNumbers() {
	outerLoop:
		for i := 0; i < 5; i++{
			for j := 1; j < 5; j++ {
				fmt.Printf("i = %d, j = %d \n", i, j)
				if i == j {
					break outerLoop
				}
			}

		}
}

func PrintEvenNumbers() {
	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
}

func PrintSequence() {
	for i, j := 1, 10; i <= 10 && j <= 20; i, j = i+1, j+1 {
		fmt.Printf("%d X %d = %d \n", i, j, i * j)
	}
}
