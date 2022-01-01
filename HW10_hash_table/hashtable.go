package HW10_hash_table

import (
	"fmt"
	"github.com/avbru/algo/containers"
)

type Hasher interface {
	GetHashCode() uint
}

type Hint uint

func (v Hint) GetHashCode() uint {
	return uint(v)
}

type HashTable struct {
	len        uint
	loadFactor uint //Average N in buckets
	growFactor uint
	buckets    []containers.List
}

type Elem struct {
	Key Hasher
	Val interface{}
}

func New() *HashTable {
	ht := HashTable{
		loadFactor: 2,
		growFactor: 4,
		buckets:    make([]containers.List, 2)}

	for i := uint(0); i < 2; i++ {
		ht.buckets[i] = containers.NewList()
	}
	return &ht
}

func (ht *HashTable) Put(key Hasher, val interface{}) {
	//Replace
	bucketIdx := key.GetHashCode() % uint(len(ht.buckets))
	elem := ht.get(key, bucketIdx)
	if elem != nil {
		elem.Value = Elem{key, val}
		return
	}

	//Resize
	if ht.len/uint(len(ht.buckets)) >= ht.loadFactor {
		ht.resize()
		bucketIdx = key.GetHashCode() % uint(len(ht.buckets))
	}

	//Insert
	ht.buckets[bucketIdx].PushFront(Elem{
		Key: key,
		Val: val,
	})
	ht.len++
}

func (ht *HashTable) Get(key Hasher) (interface{}, bool) {
	bucketIdx := key.GetHashCode() % uint(len(ht.buckets))
	if elem := ht.get(key, bucketIdx); elem != nil {
		return elem.Value.(Elem).Val, true
	}

	return nil, false
}

func (ht *HashTable) get(key Hasher, bucketIdx uint) *containers.Item {
	if ht.buckets[bucketIdx].Len() == 0 {
		return nil
	}

	for elem := ht.buckets[bucketIdx].Front(); ; {
		if elem.Value.(Elem).Key == key {
			return elem
		}
		if elem.Prev == nil {
			break
		}

		elem = elem.Prev
	}
	return nil
}

func (ht *HashTable) Delete(key Hasher) {
	bucketIdx := key.GetHashCode() % uint(len(ht.buckets))
	elem := ht.get(key, bucketIdx)

	if elem != nil {
		ht.buckets[bucketIdx].Remove(elem)
		ht.len--
	}
}

func (ht *HashTable) resize() {
	newSize := uint(len(ht.buckets)) * ht.growFactor
	newBuckets := make([]containers.List, newSize)
	for i := uint(0); i < uint(len(newBuckets)); i++ {
		newBuckets[i] = containers.NewList()
	}

	for _, bucket := range ht.buckets {
		for elem := bucket.Front(); ; {
			if elem == nil {
				break
			}

			e := elem.Value.(Elem)
			bucketIdx := e.Key.GetHashCode() % newSize
			newBuckets[bucketIdx].PushFront(e)
			if elem.Prev == nil {
				break
			}
			elem = elem.Prev
		}
	}
	ht.buckets = newBuckets
}

func (ht *HashTable) draw() {
	println("HASHTABLE len:", ht.len, "loadFactor: ", ht.len/uint(len(ht.buckets)))
	for k, bucket := range ht.buckets {
		println("Bucket idx: ", k)
		for elem := bucket.Front(); ; {
			if elem == nil {
				break
			}

			fmt.Printf("%v %s | ", elem.Value.(Elem).Key, elem.Value.(Elem).Val)
			if elem.Prev == nil {
				println("")
				break
			}
			elem = elem.Prev

		}
	}

	return
}
