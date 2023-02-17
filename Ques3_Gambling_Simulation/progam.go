package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Ques3
	var n, stake, trails, lostCount, wonCount, goal, numOfTimesWon, totalBets int
	fmt.Print("Enter Number of times to run the Simulation: ")
	fmt.Scan(&n)
	fmt.Print("Enter Stake amount: ")
	fmt.Scan(&stake)
	temp := stake
	fmt.Print("Enter Goal: ")
	fmt.Scan(&goal)
	for i := 0; i < n; i++ {
		fmt.Println("\nSimulation:", i+1, "\n")
		stake = temp
		for true {
			fmt.Println("------bet placed------")
			rn := rand.Intn(2-0) + 0
			switch rn {
			case 0:
				fmt.Println("---LOST---")
				stake = stake - 1
				trails = trails + 1
				lostCount = lostCount + 1
				fmt.Println("Stake left:", stake)
			case 1:
				fmt.Println("---WON---")
				stake = stake + 1
				trails = trails + 1
				wonCount = wonCount + 1
				fmt.Println("Stake left:", stake)
			}
			if stake == 0 {
				break
			}
			if stake == goal {
				break
			}
		}
		if stake == 0 {
			fmt.Println("\nYou,re Broke!!")
			fmt.Println("Totals Trails", trails)
			fmt.Println("No. of bets won:", wonCount)
			fmt.Println("Bet win percantage: ")
			fmt.Printf("%.2f", float64(wonCount)*100/float64(trails))
			fmt.Println("\n")
			totalBets = totalBets + trails
		}
		if stake == goal {
			fmt.Println("\nYou Reached the Goal!!")
			fmt.Println("Totals Trails", trails)
			fmt.Println("No. of bets won:", wonCount)
			fmt.Println("Bet win percantage: ")
			fmt.Printf("%.2f", float64(wonCount)*100/float64(trails))
			fmt.Println("\n")
			numOfTimesWon = numOfTimesWon + 1
			totalBets = totalBets + trails
		}
		trails = 0
		lostCount = 0
		wonCount = 0
	}
	fmt.Println("Number of times won in Simulation:", numOfTimesWon)
	fmt.Println("Percent win in Stimulation:", float64(numOfTimesWon)*100/float64(n))
	fmt.Println("Avg number of bets made per simulation:", totalBets/n)
}
