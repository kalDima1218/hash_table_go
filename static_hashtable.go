package main

import (
	"fmt"
	"math"
	"math/rand"
)

type HashTable struct{
	load, size int
	val []int
	t []int
}

func newHashTable() *HashTable {
	var tmp HashTable
	tmp.size, tmp.load = 3, 0
	tmp.val, tmp.t = make([]int, 3), make([]int, 3)
	return &tmp
}

func hash(x int) int{
	return int(math.Abs(float64(x)))
}

func h1(x, size int) int{
	return x % size
}

func h2(x, size int) int{
	return 1 + x % (size - 2)
}

func (mp *HashTable) insert(x int){
	if float64(mp.load) / float64(mp.size) >= 0.5{
		mp.rehash()
	}
	var i = h1(hash(x), mp.size)
	var d = h2(hash(x), mp.size)
	for mp.t[i] > 0{
		if mp.val[i] == x{
			mp.t[i]+=1
			return
		}
		i = (i + d)%mp.size
	}
	mp.t[i] = 1
	mp.val[i] = x
	mp.load++
}

func (mp *HashTable) count(x int) int{
	var i = h1(hash(x), mp.size)
	var d = h2(hash(x), mp.size)
	for mp.t[i] > 0{
		if mp.val[i] == x{
			return mp.t[i]
		}
		i = (i + d) % mp.size
	}
	return 0
}

func (mp *HashTable) rehash(){
	var _mp HashTable
	_mp.size, _mp.load = mp.size * 2, 0
	_mp.val, _mp.t = make([]int, mp.size * 2), make([]int, _mp.size * 2)
	for i := 0; i < mp.size; i++{
		for j := 0; j < mp.t[i]; j++{
			_mp.insert(mp.val[i])
		}
	}
	*mp = _mp
}

func main() {
	var mp HashTable
	mp = *newHashTable()
	for i := 0; i < 10; i++{
		mp.insert(rand.Int()%10)
	}
	for i := 0; i < 10; i++{
		if mp.count(i) > 0{
			fmt.Println(i, mp.count(i))
		}
	}
}
