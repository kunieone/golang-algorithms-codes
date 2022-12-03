package utils

import (
	"math/rand"
	"time"
)

type Ord interface {
	string | uint | int | float32 | float64
}

/*
实现三元表达式的功能
*/
func IF(condition bool, t, f interface{}) interface{} {
	if condition {
		return t
	} else {
		return f
	}
}

func sf(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

func SpawnArrN(length int) []int {
	var arr []int

	for i := 0; i < length; i++ {
		arr = append(arr, i+1)
	}
	return sf(arr)
}

// 取得最大值
func Max(a, b int) (maxValue int) {
	switch {
	case a > b:
		maxValue = a
	case a <= b:
		maxValue = b
	}
	return
}
