package datastructures

import (
	"testing"
)

func TestHash(t *testing.T) {
	ht := &HashTable{}

	// Test hashing consistency
	key := "hello"
	expectedIndex := ht.hash(key)
	for i := 0; i < 10; i++ {
		index := ht.hash(key)
		if index != expectedIndex {
			t.Errorf("Hash inconsistency, expected %d, got %d", expectedIndex, index)
		}
	}
}

func TestInsertAndFind(t *testing.T) {
	ht := &HashTable{}

	// Test insert and find without collision
	ht.insert("key1", "value1")
	node := ht.find("key1")
	if node == nil || node.value != "value1" {
		t.Errorf("Insert or Find failed, expected value 'value1', got %v", node)
	}

	// Test insert and find with collision
	ht.insert("key2", "value2") // Assume "key2" hashes to the same index as "key1"
	node = ht.find("key2")
	if node == nil || node.value != "value2" {
		t.Errorf("Insert or Find failed with collision, expected value 'value2', got %v", node)
	}

	// Test find on non-existent key
	node = ht.find("nonexistent")
	if node != nil {
		t.Errorf("Find failed, expected nil for nonexistent key, got %v", node)
	}
}

func TestInsertCollision(t *testing.T) {
	ht := &HashTable{}

	// This test depends on "key1" and "keyA" having the same hash value
	// This is usually controlled by the hash function and might need adjusting based on actual hash function behavior
	key1 := "key1"
	keyA := "keyA"
	if ht.hash(key1) != ht.hash(keyA) {
		t.Skip("Skipping collision test due to non-colliding hash")
	}

	ht.insert(key1, "value1")
	ht.insert(keyA, "valueA")

	// Check both entries are retrievable and correctly linked
	node1 := ht.find(key1)
	nodeA := ht.find(keyA)

	if node1 == nil || node1.value != "value1" {
		t.Errorf("Collision handling failed for %s, expected value 'value1', got %v", key1, node1)
	}
	if nodeA == nil || nodeA.value != "valueA" {
		t.Errorf("Collision handling failed for %s, expected value 'valueA', got %v", keyA, nodeA)
	}
}
