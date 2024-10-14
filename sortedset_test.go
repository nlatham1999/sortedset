package sortedset

import (
	"testing"
)

func TestNewSortedSet(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	if !ss.Contains(1) {
		panic("NewSortedSet failed")
	}
	if !ss.Contains(2) {
		panic("NewSortedSet failed")
	}
	if !ss.Contains(3) {
		panic("NewSortedSet failed")
	}
}

func TestSortedSet_Add(t *testing.T) {
	ss := NewSortedSet()
	ss.Add(1)
	if !ss.Contains(1) {
		panic("SortedSet_Add failed")
	}
}

func TestSortedSet_Add_ErrItemExists(t *testing.T) {
	ss := NewSortedSet(1)
	err := ss.Add(1)
	if err != ErrItemExists {
		panic("SortedSet_Add_ErrItemExists failed")
	}
}

func TestSortedSet_Contains(t *testing.T) {
	ss := NewSortedSet(1)
	if !ss.Contains(1) {
		panic("SortedSet_Contains failed")
	}
}

func TestSortedSet_Contains_false(t *testing.T) {
	ss := NewSortedSet()
	if ss.Contains(1) {
		panic("SortedSet_Contains_false failed")
	}
}

func TestSortedSet_Remove(t *testing.T) {
	ss := NewSortedSet(1)
	ss.Remove(1)
	if ss.Contains(1) {
		panic("SortedSet_Remove failed")
	}
}

func TestSortedSet_List(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	if len(ss.List()) != 3 {
		panic("SortedSet_List failed")
	}

	if ss.List()[0] != 1 {
		panic("SortedSet_List failed")
	}

	if ss.List()[1] != 2 {
		panic("SortedSet_List failed")
	}

	if ss.List()[2] != 3 {
		panic("SortedSet_List failed")
	}
}

func TestSortedSet_Next(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	next, _ := ss.Next(1)
	if next != 2 {
		panic("SortedSet_Next failed")
	}
}

func TestSortedSet_Previous(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	previous, _ := ss.Previous(2)
	if previous != 1 {
		panic("SortedSet_Previous failed")
	}
}

func TestSortedSet_InsertBefore(t *testing.T) {
	ss := NewSortedSet(1, 3)
	ss.InsertBefore(2, 3)
	v := ss.List()
	if v[1] != 2 {
		panic("SortedSet_InsertBefore failed")
	}
}

func TestSortedSet_InsertBeforeAtHead(t *testing.T) {
	ss := NewSortedSet(1, 3)
	ss.InsertBefore(2, 1)
	v := ss.List()
	if v[0] != 2 {
		panic("SortedSet_InsertBeforeAtHead failed")
	}
}

func TestSortedSet_InsertBeforeAtTail(t *testing.T) {
	ss := NewSortedSet(1, 2)
	ss.InsertBefore(3, 2)
	v := ss.List()
	if v[2] != 2 {
		panic("SortedSet_InsertBeforeAtTail failed")
	}
}

func TestSortedSet_InsertBeforeLen1(t *testing.T) {
	ss := NewSortedSet(1)
	ss.InsertBefore(2, 1)
	if ss.head != 2 {
		panic("SortedSet_InsertBeforeLen1 failed")
	}
	if ss.tail != 1 {
		panic("SortedSet_InsertBeforeLen1 failed")
	}
}

func TestSortedSet_InsertAfter(t *testing.T) {
	ss := NewSortedSet(1, 3)
	ss.InsertAfter(2, 1)
	v := ss.List()
	if v[1] != 2 {
		panic("SortedSet_InsertAfter failed")
	}
}

func TestSortedSet_InsertAfterAtHead(t *testing.T) {
	ss := NewSortedSet(1, 3)
	ss.InsertAfter(2, 3)
	v := ss.List()
	if v[0] != 1 {
		panic("SortedSet_InsertAfterAtHead failed")
	}
}

func TestSortedSet_InsertAfterAtTail(t *testing.T) {
	ss := NewSortedSet(1, 2)
	ss.InsertAfter(3, 2)
	v := ss.List()
	if v[2] != 3 {
		panic("SortedSet_InsertAfterAtTail failed")
	}
}

func TestSortedSet_InsertAfterLen1(t *testing.T) {
	ss := NewSortedSet(1)
	ss.InsertAfter(2, 1)
	if ss.head != 1 {
		panic("SortedSet_InsertAfterLen1 failed")
	}
	if ss.tail != 2 {
		panic("SortedSet_InsertAfterLen1 failed")
	}
}

func TestSortedSet_SortAsc(t *testing.T) {
	ss := NewSortedSet(3, 2, 1)
	ss.SortAsc(func(e interface{}) interface{} {
		return e
	})
	v := ss.List()
	if v[0] != 1 {
		panic("SortedSet_SortAsc failed")
	}
	if v[1] != 2 {
		panic("SortedSet_SortAsc failed")
	}
	if v[2] != 3 {
		panic("SortedSet_SortAsc failed")
	}
}

func TestSortedSet_SortDesc(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	ss.SortDesc(func(e interface{}) interface{} {
		return e
	})
	v := ss.List()
	if v[0] != 3 {
		panic("SortedSet_SortDesc failed")
	}
	if v[1] != 2 {
		panic("SortedSet_SortDesc failed")
	}
	if v[2] != 1 {
		panic("SortedSet_SortDesc failed")
	}
}

func TestSortedSet_AscLen1(t *testing.T) {
	ss := NewSortedSet(1)
	ss.SortAsc(func(e interface{}) interface{} {
		return e
	})
	v := ss.List()
	if v[0] != 1 {
		panic("SortedAscLen1 failed")
	}
}

func TestSortedSet_DescLen0(t *testing.T) {
	ss := NewSortedSet()
	ss.SortDesc(func(e interface{}) interface{} {
		return e
	})
	v := ss.List()
	if len(v) != 0 {
		panic("SortedDescLen0 failed")
	}
}

func TestSortedSet_ComplexStruct(t *testing.T) {
	type coord struct {
		x, y int
	}
	ss := NewSortedSet(coord{1, 2}, coord{2, 3}, coord{3, 4})
	ss.SortDesc(func(e interface{}) interface{} {
		c := e.(coord)
		return c.x
	})

	v := ss.List()
	if v[0].(coord).x != 3 {
		panic("SortedSet_ComplexStruct failed")
	}
	if v[1].(coord).x != 2 {
		panic("SortedSet_ComplexStruct failed")
	}
	if v[2].(coord).x != 1 {
		panic("SortedSet_ComplexStruct failed")
	}
}
