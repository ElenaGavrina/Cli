package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main (){
	reader := bufio.NewReader(os.Stdin)
	m, err := reader.ReadString('\n')
	if err != nil {
    	    fmt.Println(err)
  	}
	fmt.Print(isPalindrom(m))
}	

func isPalindrom (some string) bool {
    l :=strings.TrimSpace(strings.ToLower(some))
    for i:=0; i < len(l)/2; i++ {
    	if (l[i] != l[len(l) -1 - i])  {
    	    return false
    	}
    }
    return true
}
