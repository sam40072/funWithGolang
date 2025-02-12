package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type fibinachiNumber struct {
	maxNumber int
	numbers   []int
}

type numberProblem struct {
	startNumber  int
	pathToNumber []int
	stepAmount   int
}
type resultsColatzSim struct {
	largestNumOfSteps        int
	largetstNumOfstepsNumber int
	averageSteps             int
	lowerRange               int
	upperRange               int
}
type coinflips struct {
	biggestHeadsStreak int
	biggestTailStreak  int
	flipSequence       []string
}

func fibinachi(number int) fibinachiNumber {
	if number < 1 {
		log.Fatal("number must be more than 0")
	}
	numbersList := []int{0, 1}
	for i := 0; i < number; i++ {
		var newNumber int = numbersList[i] + numbersList[i+1]
		numbersList = append(numbersList, newNumber)
	}
	fibi := fibinachiNumber{maxNumber: numbersList[len(numbersList)-1], numbers: numbersList}
	return fibi
}

func colatzConjunction(startNumber int) numberProblem {
	steps := []int{}
	var numb int = startNumber
	for numb != 1 {
		if numb%2 == 0 {
			numb /= 2
		} else {
			numb = (numb * 3) + 1
		}
		steps = append(steps, numb)
	}
	colatzNumb := numberProblem{startNumber: startNumber, pathToNumber: steps, stepAmount: len(steps)}
	return colatzNumb
}

func addUpList(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func lenLoop10(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func conjunctionSimulator(numberRange [2]int) resultsColatzSim {
	fmt.Printf("testing numbers %d - %d\n", numberRange[0], numberRange[1])
	stepsList := []int{}
	var maxStepNumber int = 0
	var maxStepsNum int = 0
	var numbersToCycle int = (numberRange[1] - numberRange[0])
	proggress := 0
	divisor := max(1, numbersToCycle/10) // Prints every 10%
	for i := numberRange[0]; i <= numberRange[1]; i++ {
		if i%divisor == 0 || i == numberRange[1] {
			fmt.Printf("progress %d/%d\n", proggress, numbersToCycle)

		}
		proggress += 1
		coltz := colatzConjunction(i)
		stepsList = append(stepsList, coltz.stepAmount)
		if coltz.stepAmount > maxStepNumber {
			maxStepNumber = coltz.stepAmount
			maxStepsNum = coltz.startNumber
		}
	}
	resultant := resultsColatzSim{largestNumOfSteps: maxStepNumber, largetstNumOfstepsNumber: maxStepsNum,
		averageSteps: addUpList(stepsList) / len(stepsList), lowerRange: numberRange[0], upperRange: numberRange[1]}
	return resultant
}

func flipCoin(flips int) coinflips {
	flipSeq := []string{}
	var conseqHeads int = 0
	var tempConseqHeads int = 0
	var conseqTails int = 0
	var tempConseqTails int = 0
	for i := 0; i < flips; i++ {
		flip := rand.Int31n(2)
		var flipOutcome string
		if flip == 0 {
			flipOutcome = "heads"
		} else {
			flipOutcome = "tails"
		}
		flipSeq = append(flipSeq, flipOutcome)
	}
	for ch := 0; ch < len(flipSeq); ch++ {
		if flipSeq[ch] == "heads" {
			tempConseqHeads++
			if tempConseqHeads > conseqHeads {
				conseqHeads = tempConseqHeads
			}
		} else {
			tempConseqHeads = 0
		}
	}
	for ct := 0; ct < len(flipSeq); ct++ {
		if flipSeq[ct] == "tails" {
			tempConseqTails++
			if tempConseqTails > conseqTails {
				conseqTails = tempConseqTails
			}
		} else {
			tempConseqTails = 0
		}
	}
	coinr := coinflips{biggestHeadsStreak: conseqHeads, biggestTailStreak: conseqTails, flipSequence: flipSeq}
	return coinr
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		log.Fatal("run help for options")
	}
	if arguments[0] == "help" {
		fmt.Println("f [fibinachi] <num of steps>")
		fmt.Println("c [collatz conjunction] <test number>")
		fmt.Println("cs [collatz many runs] <lower bound number> <upper bound number>")
		fmt.Println("cf [coinflip] <num of coins> <bool print coinflip list (t/f)>")
		return
	}
	argumentsToInt, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("error")
	}
	switch argOne := arguments[0]; argOne {
	case "f":
		fmt.Println("fibinachi")
		fibiNum := fibinachi(argumentsToInt)
		fmt.Println(fibiNum.numbers)
		fmt.Println(fibiNum.maxNumber)

	case "c":
		fmt.Println("colatz")
		coltzcunj := colatzConjunction(argumentsToInt)
		fmt.Printf("%v\noriginal number: %d\nSteps: %d\n", coltzcunj.pathToNumber, coltzcunj.startNumber, coltzcunj.stepAmount)
	case "cs":
		if len(arguments) < 3 {
			log.Fatal("cs missing extra arguments")
		}
		fmt.Println("colatz simulation")
		lowerBound := argumentsToInt
		if err != nil {
			log.Fatal("error thrown lowerbound")
		}
		upperbound, err := strconv.Atoi(arguments[2])
		if err != nil {
			log.Fatal("error thrown upperbound")

		}
		bounds := [2]int{lowerBound, upperbound}
		resultz := conjunctionSimulator(bounds)
		fmt.Printf("largest num of steps: %d\nnumber that got largest steps: %d\naverage steps: %d\n", resultz.largestNumOfSteps, resultz.largetstNumOfstepsNumber, resultz.averageSteps)
	case "cf":
		fmt.Println("coinflip")
		coinFlipResults := flipCoin(argumentsToInt)
		if len(arguments) == 3 {
			if arguments[2] == "t" {
				fmt.Printf("%v\nbiggest head streak: %d\nbiggest tails streak: %d", coinFlipResults.flipSequence, coinFlipResults.biggestHeadsStreak, coinFlipResults.biggestTailStreak)
			} else if arguments[2] == "f" {
				fmt.Printf("biggest head streak: %d\nbiggest tails streak: %d", coinFlipResults.biggestHeadsStreak, coinFlipResults.biggestTailStreak)
			} else {
				log.Fatal("wrong arguments. do help command")
			}
		} else {
			fmt.Printf("biggest head streak: %d\nbiggest tails streak: %d", coinFlipResults.biggestHeadsStreak, coinFlipResults.biggestTailStreak)
		}
	}
}
