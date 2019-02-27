package interfaces

import "fmt"

// VowelsFinder interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

// MyString custom string type
type MyString string

// FindVowels MyString implements
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func init() {
	fmt.Println()
	fmt.Println("===> enter interfaces package")

	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	fmt.Println("<=== exit interfaces package")
	fmt.Println()
}
