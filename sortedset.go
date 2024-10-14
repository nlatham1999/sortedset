package sortedset

import (
	"errors"
	"sort"
)

type setKey struct {
	previous interface{}
	next     interface{}
}

type SortedSet struct {
	head   interface{}
	tail   interface{}
	values map[interface{}]*setKey
}

type sortedSetSorter struct {
	elements []interface{}
	f        func(e interface{}) interface{}
	reverse  bool
}

func (s *sortedSetSorter) Len() int {
	return len(s.elements)
}

func (s *sortedSetSorter) Less(i, j int) bool {

	var element1 interface{}
	var element2 interface{}
	if s.reverse {
		element1 = s.f(s.elements[i])
		element2 = s.f(s.elements[j])
	} else {
		element1 = s.f(s.elements[j])
		element2 = s.f(s.elements[i])
	}

	switch element1.(type) {
	case int:
		return element1.(int) > element2.(int)
	case string:
		return element1.(string) > element2.(string)
	case float64:
		return element1.(float64) > element2.(float64)
	default:
		return false
	}
}

func (s *sortedSetSorter) Swap(i int, j int) {
	s.elements[i], s.elements[j] = s.elements[j], s.elements[i]
}

var (
	ErrItemExists      = errors.New("item exists")
	ErrItemDoesntExist = errors.New("item doesn't exist")
)

// create a new sorted set with the given values
func NewSortedSet(values ...interface{}) *SortedSet {
	ss := &SortedSet{
		head:   nil,
		tail:   nil,
		values: make(map[interface{}]*setKey),
	}
	for _, value := range values {
		ss.Add(value)
	}
	return ss
}

// adds a value to the end of the set
func (ss *SortedSet) Add(value interface{}) error {
	if _, ok := ss.values[value]; ok {
		return ErrItemExists
	}

	key := &setKey{}

	if ss.head == nil {
		ss.head = value
		ss.tail = value
		key.previous = nil
		key.next = nil
	} else {
		key.previous = ss.tail
		key.next = nil
		ss.values[ss.tail].next = value
		ss.tail = value
	}
	ss.values[value] = key

	return nil
}

// returns if the value exists in the set
func (ss *SortedSet) Contains(value interface{}) bool {
	_, ok := ss.values[value]
	return ok
}

// returns the first value in the set
func (ss *SortedSet) First() interface{} {
	return ss.head
}

// inserts a value after another value
func (ss *SortedSet) InsertAfter(value, after interface{}) error {
	if _, ok := ss.values[value]; ok {
		return ErrItemExists
	}

	if after == ss.tail {
		ss.tail = value
	}

	key := &setKey{}

	if afterKey, ok := ss.values[after]; ok {
		key.previous = after
		key.next = afterKey.next
		afterKey.next = value
		ss.values[value] = key
	} else {
		return ErrItemDoesntExist
	}

	return nil
}

// inserts a value before another value
func (ss *SortedSet) InsertBefore(value, before interface{}) error {
	if _, ok := ss.values[value]; ok {
		return ErrItemExists
	}

	key := &setKey{}

	if _, ok := ss.values[before]; !ok {
		return ErrItemDoesntExist
	}

	if before == ss.head {
		key.previous = nil
		key.next = before
		ss.values[before].previous = value
		ss.head = value
		ss.values[value] = key
		return nil
	}

	ss.values[ss.values[before].previous].next = value
	key.previous = ss.values[before].previous
	ss.values[before].previous = value
	key.next = before
	ss.values[value] = key

	return nil
}

// returns the last value in the set
func (ss *SortedSet) Last() interface{} {
	return ss.tail
}

// returns the length of the set
func (ss *SortedSet) Len() int {
	return len(ss.values)
}

// returns the set as a slice
func (ss *SortedSet) List() []interface{} {
	slice := make([]interface{}, 0, len(ss.values))
	for key := ss.head; key != nil; key = ss.values[key].next {
		slice = append(slice, key)
	}
	return slice
}

// returns the next value in the set
func (ss *SortedSet) Next(value interface{}) (interface{}, error) {
	if key, ok := ss.values[value]; ok {
		return key.next, nil
	}
	return nil, ErrItemDoesntExist
}

// returns the previous value in the set
func (ss *SortedSet) Previous(value interface{}) (interface{}, error) {
	if key, ok := ss.values[value]; ok {
		return key.previous, nil
	}
	return nil, ErrItemDoesntExist
}

// removes a value from the set
func (ss *SortedSet) Remove(value interface{}) error {
	if _, ok := ss.values[value]; ok {
		key := ss.values[value]
		if key.previous != nil {
			ss.values[key.previous].next = key.next
		}
		if key.next != nil {
			ss.values[key.next].previous = key.previous
		}
		delete(ss.values, value)
	} else {
		return ErrItemDoesntExist
	}

	return nil
}

func (ss *SortedSet) sort(reverse bool, f func(e interface{}) interface{}) {
	// convert to list
	list := ss.List()

	// sort the list using the sorting interface
	sorter := &sortedSetSorter{
		elements: list,
		f:        f,
		reverse:  reverse,
	}
	sort.Sort(sorter)

	//reset the sorted set as the list
	ss.head = nil
	ss.tail = nil
	ss.values = map[interface{}]*setKey{}
	for _, value := range list {
		ss.Add(value)
	}

}

// sorts the set in ascending order
func (ss *SortedSet) SortAsc(f func(e interface{}) interface{}) {
	ss.sort(false, f)
}

// sorts the set in descending order
func (ss *SortedSet) SortDesc(f func(e interface{}) interface{}) {
	ss.sort(true, f)
}
