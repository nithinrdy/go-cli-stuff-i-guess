package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().Unix())
	dice := flag.String("d", "d6", "Input of the form dX, where X is the number of sides you wish the die to have (default is d6).")
	rolls := flag.Int("r", 1, "The number of times you wish the di(c)e to be rolled (default is 1).")
	sum := flag.Bool("s", false, "Get the sum of all the dice roll results, instead of printing out each result (default is false).")

	flag.Parse()

	if *rolls < 1 {
		fmt.Printf("You wish to \"un-roll\" the di(c)e???\n")
		os.Exit(1)
	}

	if string((*dice)[0]) != "d" {
		fmt.Println("The argument for -d needs to be of the form dX, where X is an integer")
		os.Exit(1)
	} else {
		var numStr = string((*dice)[1:])
		if strings.IndexFunc(numStr, func(c rune) bool {
			return c < '0' || c > '9'
		}) == -1 {
			numInt, _ := strconv.Atoi(numStr)
			rollResults := rollDice(numInt, rolls)
			if *sum {
				resultSum := 0
				for _, dice := range rollResults {
					resultSum += dice
				}
				fmt.Printf("The sum of all the dice rolls is %v\n", resultSum)
			} else {
				for ind, roll := range rollResults {
					fmt.Printf("Roll %v got you a %v!\n", ind+1, roll)
				}
			}
		} else {
			fmt.Println("The argument for -d needs to be of the form dX, where X is an integer")
			os.Exit(1)
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
