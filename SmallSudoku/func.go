package main

import (
	"math/rand"
	"fmt"
)

func Nul(sudoku *[3][3]int)  {
	for i:=0 ; i<3 ; i++ {
		for j:=0 ; j<3 ; j++  {
			sudoku[i][j]=0
		}
	}
}

func CheckR(sudoku *[3][3]int, i , j int) bool {
	for k:=0 ; k<3 ; k++  {
		if k == j {
			continue
		}
		if sudoku[i][j] == sudoku [i][k] {
			return false
		}
	}
	return true
}

func CheckS(sudoku *[3][3]int, i , j int) bool {
	for k:=0 ; k<3 ; k++  {
		if k == i {
			continue
		}
		if sudoku[i][j] == sudoku [k][j] {
			return false
		}
	}
	return true
}

func CheckC(sudoku *[3][3]int, i , j int) bool {

	for  k:=0 ; k<3 ; k++ {
		for  l:=0 ; l<3 ; l++  {
			if i == k && j == l {
				continue
			}
			if sudoku[i][j] == sudoku [k][l] {
				return false
			}
		}
	}
	return true
}


func Napln(sudoku *[3][3]int)  {
	for i:=0 ; i<3 ; i++  {
		for j:=0 ; j<3 ; j++  {
			sudoku[i][j] = rand.Intn(9) + 1
			if CheckC(sudoku, i, j) == false || CheckR(sudoku, i, j) == false || CheckS(sudoku, i, j) == false{
				j--
			}
		}
	}
}

func Vypis(sudoku *[3][3]int)  {
	for i:=0 ; i<3 ; i++ {
		for j:=0 ; j<3 ; j++  {
			fmt.Print("\t", sudoku[i][j])
		}
		fmt.Println()
	}
}