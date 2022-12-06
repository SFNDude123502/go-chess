package main

import "fmt"

type Uk interface {
	~string |
		~byte |
		~rune |
		~int |
		~uint |
		~float32 |
		~float64
}

func hash[K Uk](key K, size int) uint64 {
	var val string = fmt.Sprint(key)
	var out uint64 = 0
	for j := range val {
		out += uint64(val[j])
	}
	return out % uint64(size)
}

type uPair[K Uk, V any] struct {
	key   K
	value V
	next  *uPair[K, V]
}

type umap[K Uk, V any] struct {
	arr  []*uPair[K, V]
	size int
	len  int
}

func (u *umap[K, V]) mkPair(key K, value V) *uPair[K, V] {
	return &uPair[K, V]{key: key, value: value}
}

func mkUmap[K Uk, V any](size int) *umap[K, V] {
	var u *umap[K, V] = &umap[K, V]{size: size, len: 0, arr: make([]*uPair[K, V], size)}
	for i := range u.arr {
		u.arr[i] = nil
	}
	return u
}

func (u *umap[K, V]) set(key K, value V) {
	var item *uPair[K, V] = u.mkPair(key, value)
	var index = hash(item.key, u.size)
	var pos = u.arr[index]

	if pos == nil {
		if u.len == u.size {
			fmt.Println("Insert Error: Hash Table is full")
			return
		}

		u.arr[index] = item
		u.len++
	} else {
		if pos.key == key {
			u.arr[index].value = value
		} else {
			u.arr[index].next = u.mkPair(key, value)
		}
	}
}

func (u *umap[K, V]) get(key K) (out V) {
	var index uint64 = hash(key, u.size)
	var item *uPair[K, V] = u.arr[index]
	for item != nil {
		if item.key == key {
			out = item.value
			break
		}
		item = item.next
	}
	return
}

func (u *umap[K, V]) delete(key K) {
	var index uint64 = hash(key, u.size)
	var item *uPair[K, V] = u.arr[index]
	if item == nil {
		return
	}
	if item.key == key {
		u.arr[index] = item.next
		u.len--
		return
	}
}

func (u *umap[K, V]) iterator() (out []*uPair[K, V]) {
	var pos *uPair[K, V]
	for _, ival := range u.arr {
		pos = ival
		for pos != nil {
			out = append(out, pos)
			pos = pos.next
		}

	}
	return
}

func (u *umap[K, V]) print() {
	for _, ival := range u.iterator() {
		fmt.Println("[", ival.key, ":", ival.value, "]")
	}
}

func (u *umap[K, V]) realPrint() {
	var pos *uPair[K, V]
	for _, ival := range u.arr {
		pos = ival
		if pos != nil {
			for pos != nil {
				fmt.Print("[", ival.key, ":", ival.value, "]")
				pos = pos.next
			}
			fmt.Print("\n")
		}
	}
}
