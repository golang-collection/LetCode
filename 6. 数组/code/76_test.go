package code

import (
	"testing"
)

/**
* @Author: super
* @Date: 2021-01-07 07:52
* @Description:
**/

func TestMap(t *testing.T){
	need := make(map[byte]int)
	need['a']++
	t.Log(need['a'])
}