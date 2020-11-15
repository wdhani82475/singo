package test

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T)  {
	aa(3)
	ar := []int{1,2,4,4,5}
	ret := upper_bound_(5,4,ar)
	fmt.Println(ret)
}

func aa(b int) int {
	fmt.Println(b)
	return b*b
}


/**
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
 */
func upper_bound_( n int ,  v int ,  a []int ) int {
	// write code here
	left := 0
	right  := n-1

	for left < right {
		mid  := left + (right-left)/2
		if a[mid] == v  {
			return mid+1
		}else{
			if a[mid]>v {
				right--
			}else{
				left++
			}
		}
	}
	return n+1
}
