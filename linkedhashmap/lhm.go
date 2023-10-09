package linkedhashmap

import (
	dll "github.com/neilwhitlow/collections/doublylinkedlist"
)

type KVP[K comparable, V any] struct {
	Key   K
	Value V
	node  *dll.Node[*KVP[K, V]]
}

type LinkedHashMap[K comparable, V any] struct {
	kvpairs map[K]*KVP[K, V]
	dll     *dll.DoublyLinkedList[*KVP[K, V]]
}

func New[K comparable, V any](options ...any) *LinkedHashMap[K, V] {
	// lack of overloading bites again. feels wrong using ...any
	// just to allow specifying capacity or not, but I suppose
	// better than always forcing capacity parameter to be supplied.
	// maybe add a struct of options later
	linkedHashMap := &LinkedHashMap[K, V]{}
	var initialCapacity int
	for _, o := range options {
		switch option := o.(type) {
		case int:
			initialCapacity = option
		}
	}
	linkedHashMap.initialize(initialCapacity)
	return linkedHashMap
}

func (lhm *LinkedHashMap[K, V]) initialize(capacity int) {
	lhm.kvpairs = make(map[K]*KVP[K, V], capacity)
	lhm.dll = dll.New[*KVP[K, V]]()
}

func (lhm *LinkedHashMap[K, V]) Put(key K, value V) (priorValue V, exists bool) {
	if existingKVP, exists := lhm.kvpairs[key]; exists {
		priorValue := existingKVP.Value
		existingKVP.Value = value
		return priorValue, true
	}

	newKVP := &KVP[K, V]{
		Key:   key,
		Value: value,
	}
	newKVP.node = lhm.dll.AddLast(newKVP)
	lhm.kvpairs[key] = newKVP

	return
}

func (lhm *LinkedHashMap[K, V]) Get(key K) (val V) {
	if kvp, exists := lhm.kvpairs[key]; exists {
		return kvp.Value
	}
	return
}

func (lhm *LinkedHashMap[K, V]) First() *KVP[K, V] {
	if lhm == nil || lhm.dll == nil || lhm.dll.First() == nil {
		return nil
	}
	return lhm.dll.First().Value
}

func (lhm *LinkedHashMap[K, V]) Last() *KVP[K, V] {
	if lhm == nil || lhm.dll == nil || lhm.dll.Last() == nil {
		return nil
	}
	return lhm.dll.Last().Value
}

func (kvp *KVP[K, V]) Next() *KVP[K, V] {
	if kvp.node.Next() == nil {
		return nil
	}
	return kvp.node.Next().Value
}

func (kvp *KVP[K, V]) Prev() *KVP[K, V] {
	if kvp.node.Prev() == nil {
		return nil
	}
	return kvp.node.Prev().Value
}

func (lhm *LinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, lhm.dll.Len())
	for currentNode := lhm.dll.First(); currentNode != nil; currentNode = currentNode.Next() {
		keys = append(keys, currentNode.Value.Key)
	}
	return keys
}

// Len returns the size of the hashmap.
func (lhm *LinkedHashMap[K, V]) Len() int {
	return lhm.dll.Len()
}
