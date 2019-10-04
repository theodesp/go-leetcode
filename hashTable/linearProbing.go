package hashTable

import (
	"errors"
	"math"
)

/*
 Hash Table implementation using Linear Probing
 */
type HashTable struct {
	size int64
	keys []int64
	values []interface{}
	limit int64
}

func NewHashTable(size int) *HashTable  {
	keys := make([]int64, size, size)
	for i, _ := range keys {
		keys[i] = -1
	}
	return &HashTable{
		size: int64(size),
		keys: keys,
		values:make([]interface{}, size, size),
		limit: 0,
	}
}

/*
  Put the value into the hash table
 */
func (ht *HashTable)Put(key int64, value interface{}) error {
	if ht.limit > ht.size {
		return errors.New("hash table is full")
	}
	// Linear Probing
	hashedIndex := ht.hash(key)

	for ht.keys[hashedIndex] != -1 {
		hashedIndex += 1
		hashedIndex = hashedIndex % ht.size
	}

	ht.keys[hashedIndex] = key
	ht.values[hashedIndex] = value
	ht.limit += 1
	return nil
}

/*
  Put the value into the hash table using quadratic probing
*/
func (ht *HashTable)PutQuadratic(key int64, value interface{}) error {
	if ht.limit > ht.size {
		return errors.New("hash table is full")
	}
	// Linear Probing
	hashedIndex := ht.hash(key)
	squareIndex := float64(1)

	for ht.keys[hashedIndex] != -1 {
		hashedIndex += int64(math.Pow(squareIndex, 2))
		hashedIndex = hashedIndex % ht.size
		squareIndex += 1
	}

	ht.keys[hashedIndex] = key
	ht.values[hashedIndex] = value
	ht.limit += 1
	return nil
}

/*
  Put the value into the hash table using double hashing
*/
func (ht *HashTable)PutDoubleHash(key int64, value interface{}) error {
	if ht.limit > ht.size {
		return errors.New("hash table is full")
	}
	// Linear Probing
	hashedIndex := ht.doubleHash(key)

	for ht.keys[hashedIndex] != -1 {
		hashedIndex += 1
		hashedIndex = hashedIndex % ht.size
	}

	ht.keys[hashedIndex] = key
	ht.values[hashedIndex] = value
	ht.limit += 1
	return nil
}

/*
  Get the value from the hash table using double hashing
*/
func (ht *HashTable)GetDoubleHash(key int64, value interface{}) interface{} {
	hashedIndex := ht.doubleHash(key)

	for ht.keys[hashedIndex] != -1 {
		hashedIndex += 1
		hashedIndex = hashedIndex % ht.size
	}

	return ht.values[hashedIndex]
}

/*
  Get the value from the hash table
*/
func (ht *HashTable)Get(key int64, value interface{}) interface{} {
	hashedIndex := ht.hash(key)

	for ht.keys[hashedIndex] != -1 {
		hashedIndex += 1
		hashedIndex = hashedIndex % ht.size
	}

	return ht.values[hashedIndex]
}

/*
  Get the value from the hash table using the Quadratic way
*/
func (ht *HashTable)GetQuadratic(key int64, value interface{}) interface{} {
	hashedIndex := ht.hash(key)
	squareIndex := float64(1)

	for ht.keys[hashedIndex] != -1 {
		hashedIndex += int64(math.Pow(squareIndex, 2))
		hashedIndex = hashedIndex % ht.size
		squareIndex += 1
	}

	return ht.values[hashedIndex]
}

/*
  Get the value from the hash table
*/
func (ht *HashTable)hash(key int64) int64 {
	return key % int64(ht.size)
}

/*
  Get the value from the hash table
*/
func (ht *HashTable)doubleHash(key int64) int64 {
	R := ht.size - 2
	return R - (key % int64(ht.size)) % R
}





