package main

import (
	"fmt"
	"regexp"
)

func alphanumeric(str string) bool {
	regex,err1 := regexp.Compile("/[0-9]|[a-z]|[A-Z]|//S*$/g")
	matchResult := regex.MatchString(str)
	// matchResult, err2 := regexp.MatchString("/[0-9]|[a-z]|[A-Z]+\\S[^\\_]{0}/g", str)
	
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	return matchResult
}

func alphanumeric1(s string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return r.MatchString(s)
}

func alphanumeric2(str string) bool {
  
  if len(str) == 0 {
    return false
  }
  
  z := true
  
  for i := 0 ;  i < len(str); i++ {
    if str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' || str[i] >= '0' && str[i] <= '9'{
    } else {
      z = false
			break
    }
  }
  
  return z 
}

func main() {
 text := "PassW0rd"
//  text := ".*? "
 fmt.Println(alphanumeric2(text))
}