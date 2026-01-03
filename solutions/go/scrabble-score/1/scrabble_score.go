package scrabble

import "strings"

var scoreTable = buildScoreTable()

func Score(word string) int {
	score := 0
	for _, letter := range strings.ToUpper(word) {
		score += scoreTable[letter]
	}
	return score
}

func buildScoreTable() map[rune]int {
	input :=
		map[int]string{
			1:  "AEIOULNRST",
			2:  "DG",
			3:  "BCMP",
			4:  "FHVWY",
			5:  "K",
			8:  "JX",
			10: "QZ",
		}

	scoreTable := map[rune]int{}
	for score, letters := range input {
		for _, letter := range letters {
			scoreTable[letter] = score
		}
	}
	return scoreTable
}
