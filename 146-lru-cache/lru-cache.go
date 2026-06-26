type node struct {
    key, value int
    prev, next *node
}

type LRUCache struct {
    capacity int
    cache    map[int]*node
    head     *node // dummy; head.next is most-recently-used
    tail     *node // dummy; tail.prev is least-recently-used
}


func Constructor(capacity int) LRUCache {
    head, tail := &node{}, &node{}

    head.next = tail
    tail.prev = head

    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*node),
        head:     head,
        tail:     tail,

    }
}


func (this *LRUCache) remove(n *node) {
  n.prev.next = n.next
  n.next .prev = n.prev
}

func (this *LRUCache) addToFront(n *node) {
   n.next = this.head.next
   n.prev= this.head
   this.head.next.prev = n
    this.head.next = n

}

func (this *LRUCache) Get(key int) int {
    n, ok := this.cache[key]
    if !ok {
        return -1
    }
    this.remove(n)
    this.addToFront(n)
    return n.value
}


func (this *LRUCache) Put(key int, value int)  {
     n, ok := this.cache[key]
     if ok{
       n.value = value   
       this.remove(n)
       this.addToFront(n)
       return 
     }
     if len(this.cache) == this.capacity {
        lru := this.tail.prev
        this.remove(lru)
        delete(this.cache, lru.key)
    }
    n = &node{key: key, value: value}
    this.cache[key] = n
     this.addToFront(n)
 
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */