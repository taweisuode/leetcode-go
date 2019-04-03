package main

import "container/list"

type LRU struct {
	cap     int
	list    *list.List
	hashMap map[int]*list.Element
}

func set(lru *LRU, k int, v string) {
	if item, ok := lru.hashMap[k]; ok {
		lru.list.MoveToFront(item)
	}
	lru.list.PushFront(item)
	if lru.cap < lru.list.Len() {
		last := lru.list.Back()
		lru.list.Remove(last)
	}

}

func get(lru *LRU, k int) interface{} {
	if item, ok := lru.hashMap[k]; ok {
		lru.list.MoveToFront(item)
		return item.Value
	}
	return nil
}

func midSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
