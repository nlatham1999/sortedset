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

func TestSortedSet_After(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	next, _ := ss.After(1)
	if next != 2 {
		panic("SortedSet_Next failed")
	}
}

func TestSortedSet_Before(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	previous, _ := ss.Before(2)
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

func TestSortedSet_Remove_First(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	ss.Remove(1)
	v := ss.List()
	if len(v) != 2 {
		panic("SortedSet_Remove_First failed")
	}
	if ss.head != 2 {
		panic("SortedSet_Remove_First failed")
	}
}

func TestSortedSet_Remove_Last(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	ss.Remove(3)
	v := ss.List()
	if len(v) != 2 {
		panic("SortedSet_Remove_Last failed")
	}
	if ss.tail != 2 {
		panic("SortedSet_Remove_Last failed")
	}
}

func TestSortedSet_Next(t *testing.T) {
	ss := NewSortedSet(1, 2, 3, 4, 5)

	// make sure that creating a new SortedSet sets the pointer to the head
	v, _ := ss.Next()
	if v != 2 {
		panic("SortedSet_Next failed")
	}

	v, _ = ss.Next()
	if v != 3 {
		panic("SortedSet_Next failed")
	}

	v, _ = ss.Next()
	if v != 4 {
		panic("SortedSet_Next failed")
	}

	// make sure that calling First sets the pointer to the head
	ss.First()
	v, _ = ss.Next()
	if v != 2 {
		panic("SortedSet_Next failed")
	}

	// make sure that adding a new value to an empty set sets the pointer to the head
	ss = NewSortedSet()
	ss.Add(1)
	ss.Add(2)
	ss.Add(3)
	ss.Add(4)
	ss.Add(5)
	v, _ = ss.Next()
	if v != 2 {
		panic("SortedSet_Next failed")
	}

	// remove 2 and make sure that the pointer is now at 3
	ss.Remove(2)
	if ss.Current() != 2 {
		panic("SortedSet_Next failed")
	}
	v, _ = ss.Next()
	if v != 3 {
		panic("SortedSet_Next failed")
	}

	ss.Remove(4)
	if ss.Current() != 3 {
		panic("SortedSet_Next failed")
	}
	v, _ = ss.Next()
	if v != 5 {
		panic("SortedSet_Next failed")
	}

	ss = NewSortedSet(1, 2, 3)
	ss.Next()
	if ss.Current() != 2 {
		panic("SortedSet_Next failed")
	}

	ss.InsertAfter(4, 2)
	if ss.Current() != 2 {
		panic("SortedSet_Next failed")
	}
	v, _ = ss.Next()
	if v != 4 {
		panic("SortedSet_Next failed")
	}
}

func TestSortedSet_Previous(t *testing.T) {

	ss := NewSortedSet(1, 2, 3, 4, 5)

	// make sure that creating a new SortedSet sets the pointer to the head
	v, _ := ss.Previous()
	if v != nil {
		panic("SortedSet_Previous failed")
	}

	ss.Last()
	if ss.Current() != 5 {
		panic("SortedSet_Previous failed")
	}
	v, _ = ss.Previous()
	if v != 4 {
		panic("SortedSet_Previous failed")
	}

	v, _ = ss.Previous()
	if v != 3 {
		panic("SortedSet_Previous failed")
	}

	// make sure that adding a new value to an empty set sets the pointer to the head
	ss = NewSortedSet()
	ss.Add(1)
	ss.Add(2)
	ss.Add(3)
	ss.Add(4)
	ss.Add(5)

	ss.Last()

	// remove 2 and make sure that the pointer is now at 3
	ss.Remove(5)
	if ss.Current() != 5 {
		panic("SortedSet_Previous failed")
	}
	v, _ = ss.Previous()
	if v != 4 {
		panic("SortedSet_Previous failed")
	}

	ss.Remove(4)
	if ss.Current() != 4 {
		panic("SortedSet_Previous failed")
	}
	v, _ = ss.Previous()
	if v != 3 {
		panic("SortedSet_Previous failed")
	}

	ss = NewSortedSet(1, 2, 3)
	ss.Last()

	ss.InsertBefore(9, 3)
	if ss.Current() != 3 {
		panic("SortedSet_Previous failed")
	}
	v, _ = ss.Previous()
	if v != 9 {
		panic("SortedSet_Previous failed")
	}
}
