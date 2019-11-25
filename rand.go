package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums[rand.Intn(len(nums))])
}
