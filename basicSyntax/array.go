package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[1] = "Yo"
	fmt.Println(a[0], a[1])

	// []T 是一个元素类型为 T 的 slice。
	// len(s) 返回 slice s 的长度。
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	// startGame()

	// slices()

	// makeSlice()

	// slice 的零值是 nil
	var z []int
	if z == nil {
		fmt.Println("nil!")
	}
}

func startGame() {
	// Create a tic-tac-toe board
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func slices() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func makeSlice() {
	// slice 由函数 make 创建。这会分配一个全是零值的数组并且
	// 返回一个 slice 指向这个数组
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
}
