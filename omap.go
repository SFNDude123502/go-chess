package main

import "fmt"

type Key interface {
	~int |
		~uint |
		~float32 |
		~float64 |
		~string |
		~rune |
		~byte
}

type oNode[K Key, V any] struct {
	left   *oNode[K, V]
	right  *oNode[K, V]
	key    K
	value  V
	ignore bool
}

type omap[K Key, V any] struct {
	head     *oNode[K, V]
	len      int
	iterable []*oNode[K, V]
}

func mkOmap[K Key, V any](key K, value V) *omap[K, V] {
	var root *oNode[K, V] = &oNode[K, V]{key: key, value: value, left: nil, right: nil}
	return &omap[K, V]{head: root, iterable: []*oNode[K, V]{root}, len: 1}
}

func (o *omap[K, V]) set(key K, value V) {
	pos := o.head
	for {
		if key == pos.key {
			pos.value = value
			pos.ignore = false
			return
		} else if key < pos.key {
			if pos.left != nil {
				pos = pos.left
				continue
			}
			pos.left = &oNode[K, V]{key: key, value: value}
			o.iterable = append(o.iterable, pos.left)
			break
		} else {
			if pos.right != nil {
				pos = pos.right
			}
			pos.right = &oNode[K, V]{key: key, value: value}
			o.iterable = append(o.iterable, pos.right)
			break
		}
	}
	o.len++
}

func (o *omap[K, V]) get(key K) (out V) {
	pos := o.head
	for pos != nil {
		if key == pos.key {
			if !pos.ignore {
				out = pos.value
			}
			return
		}
		if key < pos.key {
			pos = pos.left
		} else {
			pos = pos.right
		}
	}
	return
}

func (o *omap[K, V]) delete(key K) {
	pos := o.head
	for pos != nil {
		if key == pos.key {
			pos.ignore = true
			return
		} else if key < pos.key {
			pos = pos.left
		} else {
			pos = pos.right
		}
	}
}

func (o *omap[K, V]) print() {
	for _, ival := range o.iterable {
		fmt.Print("[", ival.key, ":"+fmt.Sprint(ival.value), "]")
	}
	fmt.Print("\n")
}

func (o *omap[K, V]) realPrint() {
	var out [][]*oNode[K, V]
	o.reRealPrint(o.head, 0, &out)
	for _, ival := range out {
		for _, jval := range ival {
			fmt.Print("[", jval.key, ":", fmt.Sprint(jval.value), "]")
		}
		fmt.Print("\n")
	}
}

func (o *omap[K, V]) reRealPrint(pos *oNode[K, V], lvl int, out *[][]*oNode[K, V]) {
	if pos == nil {
		return
	}
	ouTmp := *out
	if len(ouTmp) == lvl {
		ouTmp = append(ouTmp, []*oNode[K, V]{})
	}
	ouTmp[lvl] = append(ouTmp[lvl], pos)
	*out = ouTmp
	o.reRealPrint(pos.left, lvl+1, out)
	o.reRealPrint(pos.right, lvl+1, out)
}
