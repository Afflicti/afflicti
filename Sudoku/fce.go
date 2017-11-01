package main

import (
	"fmt"
	"math/rand"
)

func CheckR(sudoku *[9][9]int, i int, j *int)  {
	for k:=0 ; k<9 ; k++  {
		if(k == *j){
			continue
		}
		if(sudoku[i][*j] == sudoku [i][k]){
			*j--
			break
		}
	}
}



func Napln(sudoku *[9][9]int)  {
	for i:=0 ; i<9 ; i++  {
		for j:=0 ; j<9 ; j++  {
			sudoku[i][j] = rand.Intn(9)+1
			CheckR(sudoku, i, &j)
		}
	}
}

func Vypis(sudoku *[9][9]int)  {
	for i:=0 ; i<9 ; i++  {
		for j:=0 ; j<9 ; j++  {
			if(j==3 || j==6){
				fmt.Print(" | ", sudoku[i][j])
			} else {
				fmt.Print("\t", sudoku[i][j])
			}
		}
		if(i==2 || i==5){
			fmt.Print("\n\t-\t-\t-\t-\t-\t-\t-\t-\t-\n")
		} else {
			fmt.Println()
		}
	}
}

