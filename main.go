package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

)

func main (){
	reader := bufio.NewReader(os.Stdin)
	m, _ := reader.ReadString('\n')
	fmt.Println(isPalindrom(string(m)))

}

func isPalindrom (some string) string {
	l :=strings.ToLower(some)
	for i:=0; i < len(l)/2; i++ {
		if (l[i] != l[len(l) -1 - i] ) {
			return "This isn't a palindrome"
		}
	}	
	return "This is a palindrome"
}