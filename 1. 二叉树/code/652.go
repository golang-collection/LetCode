package code

import (
	"strconv"
	"strings"
)

/**
* @Author: super
* @Date: 2021-01-29 21:32
* @Description:
**/

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	myMap := make(map[string]int)
	ret := make([]*TreeNode, 0)

	dnf(root, myMap, &ret)
	return ret
}

func dnf(root *TreeNode, myMap map[string]int, ret *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	lstring := dnf(root.Left, myMap, ret)
	rstring := dnf(root.Right, myMap, ret)

	var builder strings.Builder
	builder.WriteString(lstring)
	builder.WriteString(",")
	builder.WriteString(rstring)
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(root.Val))
	retstring := builder.String()

	if _, ok := myMap[retstring]; !ok {
		myMap[retstring] = 1
	} else {
		if myMap[retstring] == 1 {
			*ret = append(*ret, root)
		}
		myMap[retstring]++
	}

	return retstring
}