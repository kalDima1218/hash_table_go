package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Item interface {
	hash()int
}

type HashTable struct{
	load, size int
	val []Item
	t []int
}

func newHashTable() HashTable {
	var tmp HashTable
	tmp.size, tmp.load = 3, 0
	tmp.val = make([]Item, 3)
	tmp.t = make([]int, 3)
	return tmp
}

func _newHashTableSized(size int) HashTable {
	var tmp HashTable
	tmp.size, tmp.load = size, 0
	tmp.val = make([]Item, size)
	tmp.t = make([]int, size)
	return tmp
}

func h1(x, size int) int{
	return x % size
}

func h2(x, size int) int{
	return 1 + x % (size - 2)
}

func (mp *HashTable) insert(x Item){
	if float64(mp.load) / float64(mp.size) >= 0.5{
		mp._rehash()
	}
	var i = h1(x.hash(), mp.size)
	var d = h2(x.hash(), mp.size)
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

func (mp *HashTable) count(x Item) int{
	var i = h1(x.hash(), mp.size)
	var d = h2(x.hash(), mp.size)
	for mp.t[i] > 0{
		if mp.val[i] == x{
			return mp.t[i]
		}
		i = (i + d) % mp.size
	}
	return 0
}

func (mp *HashTable) clear(){
	*mp = newHashTable()
}

func (mp *HashTable) _rehash(){
	var _mp HashTable
	_mp.size, _mp.load = mp.size * 2, 0
	_mp.val = make([]Item, mp.size * 2)
	_mp.t = make([]int, _mp.size * 2)
	for i := 0; i < mp.size; i++{
		for j := 0; j < mp.t[i]; j++{
			_mp.insert(mp.val[i])
		}
	}
	*mp = _mp
}



type ItemInt struct{
	val int
}

func (i ItemInt)hash() int{
	return int(math.Abs(float64(i.val)))
}

func newItemInt(x int) ItemInt{
	var tmp ItemInt
	tmp.val = x
	return tmp
}

type ItemString struct{
	val string
}

func (i ItemString)hash() int{
	const b, m = 131, 1000000007
	var h = 0
	for _, c := range i.val{
		h*=b
		h%=m
		h+=int(c+1)
		h%=m
	}
	return h
}

func newItemString(x string) ItemString{
	var tmp ItemString
	tmp.val = x
	return tmp
}

func main() {
	var mp HashTable
	mp = newHashTable()

	mp.insert(newItemString("aba"))
	mp.insert(newItemString("caba"))
	fmt.Println("aba", mp.count(newItemString("aba")))
	fmt.Println("abac", mp.count(newItemString("abac")))
	fmt.Println("caba", mp.count(newItemString("caba")))

	mp.clear()

	for i := 0; i < 10; i++{
		mp.insert(newItemInt(rand.Int()%10))
	}
	for i := 0; i < 10; i++{
		if mp.count(newItemInt(i)) > 0{
			fmt.Println(i, mp.count(newItemInt(i)))
		}
	}
}
