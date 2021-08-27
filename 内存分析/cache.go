package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type Track struct {
	trackId string
}

type entry struct {
	bucket int64
	tracks []Track
	camTypes map[string]bool
	Next *entry
	Prev *entry
}

func NewEntry(bucket int64) *entry {
	return &entry{bucket: bucket, camTypes: make(map[string]bool), tracks: make([]Track,0)}
}
type YYCache struct {
	head  *entry
	tail  *entry
	len   int
	table map[int64]*entry
	start int64 // end
	end int64 // start
	mu sync.RWMutex

}

func NewYYCache() *YYCache {
	head := &entry{}
	tail := &entry{}
	head.Next = tail
	tail.Prev = head
	return &YYCache{head: head, tail: tail, table: make(map[int64]*entry)}
}

func(cache *YYCache) shouldIgnore(bucket int64) bool {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	return bucket <= cache.start
}
func(cache *YYCache) AddEntry(entry *entry) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if cache.len == 0 {
		// 空的加到尾结点
		temp := cache.tail.Prev
		temp.Next = entry
		entry.Prev = temp
		entry.Next = cache.tail
		cache.tail.Prev = entry

		// update cache's len and cache's currentBucket
		cache.len++
		cache.end = entry.bucket
		cache.table[entry.bucket] = entry
	}else {
		cur := cache.head.Next
		for cur != cache.tail {
			if cur.bucket > entry.bucket {
				break
			}
			cur = cur.Next
		}
		temp := cur.Prev
		entry.Next = cur
		cur.Prev = entry
		entry.Prev = temp
		temp.Next = entry
		cache.len++
		cache.end = max(cache.end, entry.bucket)
		cache.table[entry.bucket] = entry
	}
}


func(cache *YYCache) PopEntry() *entry {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	ent := cache.head.Next
	cache.head.Next = ent.Next
	ent.Next.Prev = cache.head
	cache.len--
	cache.start = ent.bucket
	delete(cache.table, ent.bucket)
	ent.Prev = nil
	ent.Next = nil
	return ent
}

func(cache *YYCache) print() {
	fmt.Printf("len: %d, start: %d, end: %d, list: ", cache.len, cache.start, cache.end)
	cur := cache.head.Next
	for cur != cache.tail {
		fmt.Printf("%d-> ", cur.bucket)
		cur = cur.Next
	}
}

func readMemStats() {

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}


func main() {
	cache := NewYYCache()
	for i:=1;i<1000;i++ {
		e := NewEntry(int64(i))
		cache.AddEntry(e)
	}
	readMemStats()
	for i:=100;i<500;i++ {
		id := fmt.Sprintf("xx-%d", i)
		t := Track{trackId: id}
		cache.table[int64(i)].tracks = append(cache.table[int64(i)].tracks, t)
	}
	time.Sleep(1*time.Second)
	for i:=1;i<1000;i++ {
		cache.PopEntry()

	}
	time.Sleep(5*time.Second)
	runtime.GC()
	readMemStats()


	cache.print()
}