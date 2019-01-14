package DynamicProgramming

import "fmt"
import "sort"
import "strings"
import "math"

type needlemanWunsh struct {
	DNASequence1    string
	DNASequence2    string
	ScoreMatrix     [][]int
	DirectionMatrix [][]string
}

//takes in two strings of DNA sequences
func NeedlemanWunsch(FirstSequence string, SecondSequence string) needlemanWunsh {
	nedWun := needlemanWunsh{}
	nedWun.DNASequence1 = FirstSequence
	nedWun.DNASequence2 = SecondSequence

	//dynamically creating the score matrix and initilazing its values
	nedWun.ScoreMatrix = make([][]int, len(FirstSequence)+1)
	nedWun.DirectionMatrix = make([][]string, len(FirstSequence)+1)

	for row := range nedWun.ScoreMatrix {
		nedWun.ScoreMatrix[row] = make([]int, len(SecondSequence)+1)
		nedWun.DirectionMatrix[row] = make([]string, len(SecondSequence)+1)
	}
	for i := 1; i < len(nedWun.ScoreMatrix); i++ {
		nedWun.ScoreMatrix[i][0] = nedWun.ScoreMatrix[i-1][0] - 2
		nedWun.DirectionMatrix[i][0] = "U"
	}
	for i := 1; i < len(nedWun.ScoreMatrix[0]); i++ {
		nedWun.ScoreMatrix[0][i] = nedWun.ScoreMatrix[0][i-1] - 2
		nedWun.DirectionMatrix[0][i] = "L"
	}
	nedWun.DirectionMatrix[0][0] = "End"
	return nedWun
}

func (NW needlemanWunsh) GlobalAlignment() {
	for i := 1; i < len(NW.ScoreMatrix); i++ {
		for j := 1; j < len(NW.ScoreMatrix[i]); j++ {
			value, direction := NW.maxDirection(i, j)
			NW.ScoreMatrix[i][j] = value
			NW.DirectionMatrix[i][j] = direction
		}
	}
}

func (NW needlemanWunsh) maxDirection(i int, j int) (int, string) {
	var MaxArray = make([]int, 3)
	var Up int = NW.ScoreMatrix[i-1][j] + scoring(NW.DNASequence1[i-1], '_')
	var Left int = NW.ScoreMatrix[i][j-1] + scoring('_', NW.DNASequence2[j-1])
	var Diag int = NW.ScoreMatrix[i-1][j-1] + scoring(NW.DNASequence1[i-1], NW.DNASequence2[j-1])
	MaxArray[0] = Up
	MaxArray[1] = Left
	MaxArray[2] = Diag
	sort.Ints(MaxArray)
	var DirBuilder strings.Builder
	if MaxArray[2] == Up {
		DirBuilder.WriteString("U")
	}
	if MaxArray[2] == Left {
		DirBuilder.WriteString("L")
	}
	if MaxArray[2] == Diag {
		DirBuilder.WriteString("D")
	}
	var Directions string = DirBuilder.String()
	return MaxArray[2], Directions
}

func scoring(FirstLetter byte, SecondLetter byte) int {
	if FirstLetter == '_' || SecondLetter == '_' {
		return -2
	}
	if FirstLetter == SecondLetter {
		return 1
	}
	return -1
}

//Two functions to display both matrices that are used in the algorithm
func (NW needlemanWunsh) DisplayScoreMatrix() {
	for i := range NW.ScoreMatrix {
		fmt.Println(NW.ScoreMatrix[i])
	}
}

func (NW needlemanWunsh) DisplayDirectionMatrix() {
	for i := range NW.DirectionMatrix {
		fmt.Println(NW.DirectionMatrix[i])
	}
}

//Reverses a returns a given string
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

//Traverses through the matrix to find the aligment for the two sequences
//It does this but checking the Direction matrix for the directions at a
//given a location and finded the lowest value for it
func (NW needlemanWunsh) FindAllOptimalPath() (string, string) {
	var row int = len(NW.DNASequence1)
	var column int = len(NW.DNASequence2)
	var Sequence1, Sequence2 strings.Builder

	NW.DisplayDirectionMatrix()

	for NW.DirectionMatrix[row][column] != "End" {
		var HighestScore int = math.MinInt64
		var HighestDirection string

		for _, char := range NW.DirectionMatrix[row][column] {
			if char == 'D' && NW.ScoreMatrix[row-1][column-1] > HighestScore {
				HighestDirection = "D"
				HighestScore = NW.ScoreMatrix[row-1][column-1]
			}
			if char == 'L' && NW.ScoreMatrix[row][column-1] > HighestScore {
				HighestDirection = "L"
				HighestScore = NW.ScoreMatrix[row][column-1]
			}
			if char == 'U' && NW.ScoreMatrix[row-1][column] > HighestScore {
				HighestDirection = "U"
				HighestScore = NW.ScoreMatrix[row-1][column]
			}
		}

		if HighestDirection == "D" {
			Sequence1.WriteByte(NW.DNASequence1[row-1])
			Sequence2.WriteByte(NW.DNASequence2[column-1])
			row--
			column--
		} else if HighestDirection == "L" {
			Sequence1.WriteByte('_')
			Sequence2.WriteByte(NW.DNASequence2[row-1])
			column--
		} else if HighestDirection == "U" {
			Sequence1.WriteByte(NW.DNASequence1[row-1])
			Sequence2.WriteByte('_')
			row--
		}
	}
	return reverse(Sequence1.String()), reverse(Sequence2.String())
}
