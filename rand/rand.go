package randutil

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// 随机base到base+n内的一个整数
func RandBaseInt(base int, n int) int {
	if n < 0 {
		return base
	}

	return base + rand.Intn(n)
}

// 随机0到n-1内的一个整数
func RandInt(n int) int {
	if n <= 0 {
		return 0
	}

	return rand.Intn(n)
}

// 随机1到1内的一个浮点数
func RandFloat32() float32 {
	return rand.Float32()
}
