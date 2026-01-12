package yacht

import "log"

func Score(dice []int, category string) int {

	score := 1
	switch category {
	case "ones":
		score = count(dice, 1)
	case "twos":
		score = 2 * count(dice, 2)
	case "threes":
		score = 3 * count(dice, 3)
	case "fours":
		score = 4 * count(dice, 4)
	case "fives":
		score = 5 * count(dice, 5)
	case "sixes":
		score = 6 * count(dice, 6)
	case "full house":
		score = FullHouse(dice)
	case "four of a kind":
		score = FourOfAKind(dice)
	case "little straight":
		score = LittleStraight(dice)
	case "big straight":
		score = BigStraight(dice)
	case "choice":
		score = Coice(dice)
	case "yacht":
		score = Yacht(dice)
	}
	return score
}

func Yacht(dice []int) int {
	first := dice[0]
	if count(dice, first) == 5 {
		return 50
	}
	return 0
}

func Coice(dice []int) int {
	return sum(dice)
}

func BigStraight(dice []int) int {
	return scoreStraight(dice, 2, 6)
}

func LittleStraight(dice []int) int {
	return scoreStraight(dice, 1, 5)
}

func scoreStraight(dice []int, from int, to int) int {
	table := map[int]bool{}
	for i := from; i <= to; i++ {
		table[i] = false
	}
	for i, nr := range dice {
		log.Printf("%v -> %v\n", i, nr)
		table[nr] = true
	}
	for _, v := range table {
		if !v {
			return 0
		}
	}
	return 30
}

func FourOfAKind(dice []int) int {
	first := dice[0]
	second := 0
	for _, diceNr := range dice {
		if diceNr != first {
			second = diceNr
		}
	}
	if count(dice, first) >= 4 {
		return 4 * first
	}
	if count(dice, second) == 4 {
		return 4 * second
	}
	return 0
}

func FullHouse(dice []int) int {
	first := dice[0]
	second := 0
	for _, diceNr := range dice {
		if diceNr != first {
			second = diceNr
		}
	}
	count1 := count(dice, first)
	count2 := count(dice, second)
	if (count1 == 3 && count2 == 2) || (count1 == 2 && count2 == 3) {
		return sum(dice)
	}
	return 0
}

func count(dice []int, nr int) int {
	count := 0
	for _, diceNr := range dice {
		if diceNr == nr {
			count++
		}
	}
	return count
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Later I can refactor it all to use a more functional programming style
//func Sum(numbers []int) int {
//	add := func(acc, x int) int { return acc + x }
//	return Reduce(numbers, add, 0)
//}
//
//func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
//	var result = initialValue
//	for _, x := range collection {
//		result = f(result, x)
//	}
//	return result
//}
