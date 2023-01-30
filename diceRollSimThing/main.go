package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().Unix())
	dice := flag.Int("d", 6, "The number of sides you wish for the di(c)e to have (default is 6).")
	rolls := flag.Int("r", 1, "The number of times you wish the di(c)e to be rolled (default is 1).")
	sum := flag.Bool("s", false, "Get the sum of all the dice roll results, instead of printing out each result individually (default is false).")

	flag.Parse()

	if *dice < 3 {
		fmt.Printf("There are no %v-sided dice, silly!\n", *dice)
		os.Exit(1)
	}

	rollResults := rollDice(*dice, rolls)
	if *sum {
		resultSum := 0
		for _, dice := range rollResults {
			resultSum += dice
		}
		fmt.Printf("The sum of all the dice rolls is %v!\n", resultSum)
	} else {
		for ind, roll := range rollResults {
			fmt.Printf("Roll %v got you a %v!\n", ind+1, roll)
		}
	}
}

func rollDice(num int, rolls *int) []int {
	var rollResultList []int
	for i := 0; i < *rolls; i++ {
		rollResultList = append(rollResultList, rand.Intn(num)+1)
	}
	return rollResultList
}
