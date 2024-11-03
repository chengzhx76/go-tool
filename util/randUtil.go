package util

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 指定范围的随机数
func RangeInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

/*
count: 生成的个数
min: 范围值[最小]
max: 范围值[最大]
minGap: 间隔
*/
func RandomNumbers(count, min, max, minGap int) ([]int, error) {
	if (max - min) < (count-1)*(minGap-1) {
		return nil, fmt.Errorf("the range is insufficient to generate %d unique numbers with a minimum gap of %d", count, minGap)
	}
	selected := make(map[int]bool)
	var randomNumbers []int

	for len(randomNumbers) < count {
		num := rand.Intn(max-min+1) + min

		if !selected[num] && isValid(selected, num, minGap) {
			randomNumbers = append(randomNumbers, num)
			selected[num] = true
		}
	}

	return randomNumbers, nil
}

// isValid 检查新数字与现有的是否间隔至少为 minGap
func isValid(selected map[int]bool, num, minGap int) bool {
	for i := num - (minGap - 1); i <= num+(minGap-1); i++ {
		if selected[i] {
			return false
		}
	}
	return true
}
