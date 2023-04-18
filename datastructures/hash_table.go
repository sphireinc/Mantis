package datastructures

const tableSize = 10

type HashTableNode struct {
	key   string
	value string
	next  *HashTableNode
}

type HashTable struct {
	table [tableSize]*HashTableNode
}

func (ht *HashTable) hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % tableSize
}

func (ht *HashTable) insert(key, value string) {
	index := ht.hash(key)
	if ht.table[index] == nil {
		ht.table[index] = &HashTableNode{key, value, nil}
	} else {
		newNode := &HashTableNode{key, value, ht.table[index]}
		ht.table[index] = newNode
	}
}

func (ht *HashTable) find(key string) *HashTableNode {
	index := ht.hash(key)
	node := ht.table[index]
	for node != nil && node.key != key {
		node = node.next
	}
	return node
}
