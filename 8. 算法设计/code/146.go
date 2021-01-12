package code

/**
* @Author: super
* @Date: 2021-01-12 07:36
* @Description:
**/

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func (this *LRUCache) AddNode(key, val int) {
	//双向链表添加
	temp := &DLinkedNode{
		key:   key,
		value: val,
		next:  this.tail.next,
		prev:  this.tail,
	}
	this.tail.next = temp
	this.tail = temp
	//map添加
	this.cache[key] = temp
}

func (this *LRUCache) Remove() {
	//双向链表删除
	key := this.head.next.key
	temp := this.head.next
	temp.prev = nil
	this.head.next = temp.next
	if temp.next != nil {
		temp.next.prev = this.head
		temp.next = nil
	}
	//map删除
	delete(this.cache, key)
}

func (this *LRUCache) Adjust(key int){
	temp := this.cache[key]
	if temp != this.tail{
		temp.prev.next = temp.next
		temp.next.prev = temp.prev
		temp.prev = this.tail
		temp.next = this.tail.next
		this.tail.next = temp
		temp.prev = this.tail
		this.tail = temp
	}
}

func Constructor(capacity int) LRUCache {
	head := &DLinkedNode{}
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     head,
		tail:     head,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok{
		this.Adjust(key)
		return v.value
	}else{
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok{
		this.cache[key].value = value
		this.Adjust(key)
	}else{
		if this.size < this.capacity{
			this.size++
		}else{
			this.Remove()
		}
		this.AddNode(key, value)
	}
}
