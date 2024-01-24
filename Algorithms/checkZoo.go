package main

import "fmt"

func main() {
	a := "zzzoooooo"

	wordMap := [26]int{}

	for i := 0; i < len(a); i++ {
		wordMap[a[i]-'a']++
	}

	if (wordMap['z'-'a'] * 2) == wordMap['o'-'a'] {
		fmt.Println("zoo")
	}
}
