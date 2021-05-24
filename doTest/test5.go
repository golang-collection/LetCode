package main

import (
	"fmt"
)

/**
* @Author: super
* @Date: 2021-04-17 16:04
* @Description:
**/

func change(s string) (string, bool) {
	news := ""
	flag := false
	for i, j := 0, 1; j<len(s); j++{
		if s[i] == 'a' && s[j] == 'b'{
			news = s[:i] + "bba" + s[j+1:]
		}
		i++
	}
	if news == ""{
		flag = true
	}
	return news, flag
}

func change1(s string) int{
	i := 0
	j := 1
	//tail := j
	count := 0
	for ; j<len(s); {
		if s[i] == 'a' && s[j] == 'b'{
			s = s[:i] + "bba" + s[j+1:]
			count++
			if i != 0 && s[i - 1] == 'a'{
				i--
				j--
				continue
			}
		}
		i++
		j++
	}
	return count
}

func main() {
	s := "aaababababababababbbaabbababbabababaababab"
	count := change1(s)
	fmt.Println(count % 1000000007)
}