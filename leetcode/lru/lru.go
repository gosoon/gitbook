package main

import (
	"container/list"
	"fmt"
)

type Cache interface {
	Add(key, value interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
	//GetAllocatable() (float64, float64, error)
}

type entry struct {
	key   interface{}
	value interface{}
}

type cache struct {
	length int

	eleList   map[interface{}]*list.Element
	ll        *list.List
	OnEvicted func(key interface{}, value interface{})
}

func NewCache(length int, onEvictedFunc func(key interface{}, value interface{})) Cache {
	return &cache{
		length:    length,
		ll:        list.New(),
		eleList:   make(map[interface{}]*list.Element),
		OnEvicted: onEvictedFunc,
	}
}

func (c *cache) Add(key interface{}, value interface{}) {
	if c.eleList == nil {
		c.eleList = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	// update existed value
	if v, found := c.eleList[key]; found {
		c.ll.MoveToFront(v)
		v.Value.(*entry).value = value
	}

	// add value
	ele := c.ll.PushFront(&entry{key: key, value: value})
	c.eleList[key] = ele

	fmt.Println("len:", len(c.eleList))
	// remove oldest ele
	if c.length != 0 && len(c.eleList) > c.length {
		c.RemoveOldest()
	}
	//c.ll.Len()
}

func (c *cache) RemoveOldest() {
	if c.eleList == nil {
		return
	}

	ele := c.ll.Back()
	if ele != nil {
		fmt.Println(ele.Value.(*entry).key)
		c.ll.Remove(ele)
		delete(c.eleList, ele.Value.(*entry).key)
	}
}

func (c *cache) Delete(key interface{}) {
	if c.eleList == nil {
		return
	}
	if ele, found := c.eleList[key]; found {
		c.ll.Remove(ele)
		delete(c.eleList, key)
	}
}

func (c *cache) Get(key interface{}) interface{} {
	if c.eleList == nil {
		return nil
	}
	if ele, found := c.eleList[key]; found {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return nil
}
