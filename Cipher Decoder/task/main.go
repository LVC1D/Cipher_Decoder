package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// write your code here
	var g, p, B, A, S int
	fmt.Scanf("g is %d and p is %d\n", &g, &p)
	fmt.Println("OK")

	b := 7

	B = modExp(g, b, p)

	fmt.Scanf("A is %d\n", &A)
	S = modExp(A, b, p)

	fmt.Printf("B is %d\n", B)
	shift := S % 26

	fmt.Println(CaesarEncrypt("Will you marry me?", shift))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := CaesarDecrypt(scanner.Text(), shift)
		if line == "Yeah, okay!" {
			fmt.Println(CaesarEncrypt("Great!", shift))
			break
		} else if line == "Let's be friends." {
			fmt.Println(CaesarEncrypt("What a pity!", shift))
			break
		} else {
			break
		}
	}

}

func modExp(base, b, p int) int {
	var e int
	c := 1
	for e < b {
		e++
		c = (c * base) % p
	}
	return c
}

func CaesarEncrypt(text string, shift int) string {
	var result strings.Builder
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			newChar := 'A' + (char-'A'+int32(shift)+26)%26
			result.WriteRune(newChar)
		} else if char >= 'a' && char <= 'z' {
			newChar := 'a' + (char-'a'+int32(shift)+26)%26
			result.WriteRune(newChar)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func CaesarDecrypt(text string, shift int) string {
	var result strings.Builder
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			newChar := 'A' + (char-'A'-int32(shift)+26)%26
			result.WriteRune(newChar)
		} else if char >= 'a' && char <= 'z' {
			newChar := 'a' + (char-'a'-int32(shift)+26)%26
			result.WriteRune(newChar)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
