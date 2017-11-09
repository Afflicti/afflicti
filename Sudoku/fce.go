package main

import (
	"fmt"
	"math/rand"
)

func CheckR(sudoku *[9][9]int, i , j int) bool {
	for k:=0 ; k<9 ; k++  {
		if k == j {
			continue
		}
		if sudoku[i][j] == sudoku [i][k] {
			return false
		}
	}
	return true
}

func CheckS(sudoku *[9][9]int, i , j int) bool {
	for k:=0 ; k<9 ; k++  {
		if k == i {
			continue
		}
		if sudoku[i][j] == sudoku [k][j] {
			return false
		}
	}
	return true
}

func CheckC(sudoku *[9][9]int, i , j , startR, startS int) bool {
	
	for  k:=0 ; k<3 ; k++ {
		for  l:=0 ; l<3 ; l++  {
			if i == k+startR && j == l+startS {
				continue
			}
			if sudoku[i][j] == sudoku [k+startR][l+startS] {
				return false
			}

		}
	}
	return true
}

func Napln(sudoku *[9][9]int)  {
	for i:=0 ; i<9 ; i++  {
		for j:=0 ; j<9 ; j++  {
			sudoku[i][j] = rand.Intn(9) + 1
			if CheckC(sudoku, i, j, i-(i%3), j-(j%3)) == false || CheckR(sudoku, i, j) == false || CheckS(sudoku, i, j) == false {
				j--
			}
		}
	}
}

func Vypis(sudoku *[9][9]int)  {
	for i:=0 ; i<9 ; i++  {
		for j:=0 ; j<9 ; j++  {
			if j==3 || j==6 {
				fmt.Print(" | ", sudoku[i][j])
			} else {
				fmt.Print("\t", sudoku[i][j])
			}
		}
		if i==2 || i==5 {
			fmt.Print("\n\t-\t-\t-\t-\t-\t-\t-\t-\t-\n")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("\n")
}


func Zadej(sudoku *[9][9]int)  {

	for i:=0 ; i<9 ; i++ {
		for j:=0 ; j<9 ; j++  {
			fmt.Scanf("%d", sudoku[i][j])
			fmt.Print(1)
		}
		fmt.Println()
	}
}

func Nul(sudoku *[9][9]int)  {
	for i:=0 ; i<9 ; i++ {
		for j:=0 ; j<9 ; j++  {
			sudoku[i][j]=0
		}
	}
}

