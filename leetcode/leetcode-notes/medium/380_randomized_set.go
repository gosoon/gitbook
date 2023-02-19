package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	valToIndex map[int]int
	nums       []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{valToIndex: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, found := this.valToIndex[val]; found {
		return false
	}
	this.valToIndex[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, found := this.valToIndex[val]; !found {
		return false
	}
	l := len(this.nums)
	fmt.Printf("index:%+v nums:%+v\n", this.valToIndex, this.nums)
	// swap
	index := this.valToIndex[val]
	this.valToIndex[this.nums[l-1]] = index
	this.valToIndex[val] = l - 1
	this.nums[index], this.nums[l-1] = this.nums[l-1], this.nums[index]

	fmt.Printf("index:%+v nums:%+v\n", this.valToIndex, this.nums)
	// delete val
	this.nums = this.nums[:l-1]
	delete(this.valToIndex, val)
	fmt.Printf("index:%+v nums:%+v\n", this.valToIndex, this.nums)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func main() {
	obj := Constructor()
	param_1 := obj.Insert(1)
	fmt.Printf("result: %v  val:%+v nums:%+v \n", param_1, obj.valToIndex, obj.nums)
	param_1 = obj.Remove(2)
	fmt.Printf("result: %v  val:%+v nums:%+v \n", param_1, obj.valToIndex, obj.nums)
	param_1 = obj.Insert(2)
	fmt.Printf("result: %v  val:%+v nums:%+v \n", param_1, obj.valToIndex, obj.nums)
	param_2 := obj.GetRandom()
	fmt.Printf("result: %v  val:%+v nums:%+v \n", param_2, obj.valToIndex, obj.nums)
	param_1 = obj.Remove(1)
	fmt.Printf("result: %v  val:%+v nums:%+v \n", param_1, obj.valToIndex, obj.nums)
	param_1 = obj.Insert(2)
	fmt.Printf("result: %v  val:%+v nums:%+v \n", param_1, obj.valToIndex, obj.nums)
	param_2 = obj.GetRandom()
	fmt.Printf("result: %v  val:%+v nums:%+v \n", param_2, obj.valToIndex, obj.nums)
}
