package datastructures

const tableSize = 10

// HashTableNode represents our node within our hash table
type HashTableNode struct {
	key   string
	value string
	next  *HashTableNode
}

// HashTable represents our table of a given size
type HashTable struct {
	table [tableSize]*HashTableNode
}

// hash creates a hash from our key
func (ht *HashTable) hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % tableSize
}

// insert a kv pair into our hash table
func (ht *HashTable) insert(key, value string) {
	index := ht.hash(key)
	if ht.table[index] == nil {
		ht.table[index] = &HashTableNode{key, value, nil}
	} else {
		newNode := &HashTableNode{key, value, ht.table[index]}
		ht.table[index] = newNode
	}
}

// find a key within our hash table
func (ht *HashTable) find(key string) *HashTableNode {
	index := ht.hash(key)
	node := ht.table[index]
	for node != nil && node.key != key {
		node = node.next
	}
	return node
}
