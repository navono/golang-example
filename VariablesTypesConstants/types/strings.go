package types

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Println("print string with bytes(UTF-8):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Println("print string with chars:")

	// for i := 0; i < len(s); i++ {
	// 	// 假设都是 1 字节，在打印 2 字节以上的字符串时，会出错
	// 	fmt.Printf("%c ", s[i])
	// }

	// 正确的方案
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func init() {
	// strings are converted to a slice of runes

	fmt.Println()
	fmt.Println("===> enter types package (string)")

	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")

	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	length(name)

	printCharsAndBytes(name)

	// UTF-8 编码
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	// byteSlice := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	fmt.Println("<=== exit types package (string)")
	fmt.Println()
}
