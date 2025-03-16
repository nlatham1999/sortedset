package sortedset

import (
	"fmt"
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

func TestSortedSet_Union(t *testing.T) {

	ss1 := NewSortedSet(1, 2, 3)
	ss2 := NewSortedSet(4, 5, 6)

	ss3 := ss1.Union(ss2)

	if !ss3.Contains(1) {
		panic("SortedSet_Union failed")
	}
	if !ss3.Contains(2) {
		panic("SortedSet_Union failed")
	}
	if !ss3.Contains(3) {
		panic("SortedSet_Union failed")
	}
	if !ss3.Contains(4) {
		panic("SortedSet_Union failed")
	}
	if !ss3.Contains(5) {
		panic("SortedSet_Union failed")
	}
	if !ss3.Contains(6) {
		panic("SortedSet_Union failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet(2, 3, 4)
	list := ss1.Union(ss2).List()
	if list[0] != 1 {
		panic("SortedSet_Union failed")
	}
	if list[1] != 2 {
		panic("SortedSet_Union failed")
	}
	if list[2] != 3 {
		panic("SortedSet_Union failed")
	}
	if list[3] != 4 {
		panic("SortedSet_Union failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet()
	list = ss1.Union(ss2).List()
	if list[0] != 1 {
		panic("SortedSet_Union failed")
	}
}

func TestSortedSet_Intersection(t *testing.T) {

	ss1 := NewSortedSet(1, 2, 3)
	ss2 := NewSortedSet(2, 3, 4)

	ss3 := ss1.Intersection(ss2)

	if !ss3.Contains(2) {
		panic("SortedSet_Intersection failed")
	}
	if !ss3.Contains(3) {
		panic("SortedSet_Intersection failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet(4, 5, 6)
	list := ss1.Intersection(ss2).List()
	if len(list) != 0 {
		panic("SortedSet_Intersection failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet()
	list = ss1.Intersection(ss2).List()
	if len(list) != 0 {
		panic("SortedSet_Intersection failed")
	}
}

func TestSortedSet_Difference(t *testing.T) {

	ss1 := NewSortedSet(1, 2, 3)
	ss2 := NewSortedSet(2, 3, 4)

	ss3 := ss1.Difference(ss2)

	if !ss3.Contains(1) {
		panic("SortedSet_Difference failed")
	}

	if ss3.Len() != 1 {
		panic("SortedSet_Difference failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet(4, 5, 6)
	list := ss1.Difference(ss2).List()
	if list[0] != 1 {
		panic("SortedSet_Difference failed")
	}
	if list[1] != 2 {
		panic("SortedSet_Difference failed")
	}
	if list[2] != 3 {
		panic("SortedSet_Difference failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet()
	list = ss1.Difference(ss2).List()
	if list[0] != 1 {
		panic("SortedSet_Difference failed")
	}
	if list[1] != 2 {
		panic("SortedSet_Difference failed")
	}
	if list[2] != 3 {
		panic("SortedSet_Difference failed")
	}
}

func TestSortedSet_SymmetricDifference(t *testing.T) {

	ss1 := NewSortedSet(1, 2, 3)
	ss2 := NewSortedSet(2, 3, 4)

	ss3 := ss1.SymmetricDifference(ss2)

	if !ss3.Contains(1) {
		panic("SortedSet_SymmetricDifference failed")
	}
	if !ss3.Contains(4) {
		panic("SortedSet_SymmetricDifference failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet(4, 5, 6)
	list := ss1.SymmetricDifference(ss2).List()
	if list[0] != 1 {
		panic("SortedSet_SymmetricDifference failed")
	}
	if list[1] != 2 {
		panic("SortedSet_SymmetricDifference failed")
	}
	if list[2] != 3 {
		panic("SortedSet_SymmetricDifference failed")
	}
	if list[3] != 4 {
		panic("SortedSet_SymmetricDifference failed")
	}
	if list[4] != 5 {
		panic("SortedSet_SymmetricDifference failed")
	}
	if list[5] != 6 {
		panic("SortedSet_SymmetricDifference failed")
	}

	ss1 = NewSortedSet(1, 2, 3)
	ss2 = NewSortedSet()
	list = ss1.SymmetricDifference(ss2).List()
	if list[0] != 1 {
		panic("SortedSet_SymmetricDifference failed")
	}
	if list[1] != 2 {
		panic("SortedSet_SymmetricDifference failed")
	}
	if list[2] != 3 {
		panic("SortedSet_SymmetricDifference failed")
	}
}

func TestSortedSet_Empty(t *testing.T) {
	ss := NewSortedSet()
	if !ss.Empty() {
		panic("SortedSet_Empty failed")
	}

	ss = NewSortedSet(1)
	if ss.Empty() {
		panic("SortedSet_Empty failed")
	}

	ss.Remove(ss.First())
	if !ss.Empty() {
		panic("SortedSet_Empty failed")
	}
}

func TestSortedSet_Ask(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	counter := 0
	result := 0
	ss.Ask(func(value1 interface{}) {
		ss.Ask(func(value2 interface{}) {
			ss.Ask(func(value3 interface{}) {
				counter++
				result += (value1.(int) * value2.(int) * value3.(int))
				if counter > 100 {
					panic("SortedSet_Ask failed")
				}
			})
		})
	})

	if result != 216 {
		panic("SortedSet_Ask failed")
	}

	if counter != 27 {
		panic("SortedSet_Ask failed")
	}
}

func TestSortedSet_AskWithARemove(t *testing.T) {
	// test removing the next value in the loop
	ss := NewSortedSet(1, 2, 3, 4, 5)
	loop1 := []interface{}{}
	ss.Ask(func(value1 interface{}) {
		ss.Ask(func(value2 interface{}) {
			if value1 == 2 && value2 == 1 {
				ss.Remove(3)
			}
		})
		loop1 = append(loop1, value1)
	})

	if len(loop1) != 4 {
		panic("SortedSet_AskWithARemove failed")
	}

	if loop1[0] != 1 {
		panic("SortedSet_AskWithARemove failed")
	}

	if loop1[1] != 2 {
		panic("SortedSet_AskWithARemove failed")
	}

	if loop1[2] != 4 {
		panic("SortedSet_AskWithARemove failed")
	}

	// removing current value
	ss = NewSortedSet(1, 2, 3, 4, 5)
	loop1 = []interface{}{}
	ss.Ask(func(value1 interface{}) {
		ss.Ask(func(value2 interface{}) {
			if value1 == 2 {
				ss.Remove(2)
			}
		})
		loop1 = append(loop1, value1)
	})

	if len(loop1) != 5 {
		panic("SortedSet_AskWithARemove failed")
	}

	// removing value two steps ahead
	ss = NewSortedSet(1, 2, 3, 4, 5)
	loop1 = []interface{}{}
	ss.Ask(func(value1 interface{}) {
		ss.Ask(func(value2 interface{}) {
			if value1 == 2 {
				ss.Remove(4)
			}
		})
		loop1 = append(loop1, value1)
	})

	if len(loop1) != 4 {
		panic("SortedSet_AskWithARemove failed")
	}
}

func TestSortedSet_AskWithAnAdd(t *testing.T) {

	ss := NewSortedSet(1, 2)
	loop1 := []interface{}{}
	ss.Ask(func(value1 interface{}) {
		ss.Ask(func(value2 interface{}) {
			if value1 == 2 && value2 == 1 {
				ss.Add(3)
			}
		})
		loop1 = append(loop1, value1)
	})

	if len(loop1) != 3 {
		panic("SortedSet_AskWithAnAdd failed")
	}

	if loop1[0] != 1 {
		panic("SortedSet_AskWithAnAdd failed")
	}

	if loop1[1] != 2 {
		panic("SortedSet_AskWithAnAdd failed")
	}

	if loop1[2] != 3 {
		panic("SortedSet_AskWithAnAdd failed")
	}
}

func TestSortedSet_FirstInNestedAsk(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	counter := 0
	ss.Ask(func(value1 interface{}) {

		count := 0
		for v := ss.First(); v != nil && count < 2; v, _ = ss.Next() {
			// count++
		}

		ss.Ask(func(value2 interface{}) {
			fmt.Println("value2", value2)
			counter++
		})
	})

	if counter != 9 {
		panic("SortedSet_FirstInNestedAsk failed")
	}
}

func TestSortedSet_FirstInNestedAny(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	counter := 0
	ss.Any(func(value1 interface{}) bool {

		fmt.Println("value1", value1)
		count := 0
		for v := ss.First(); v != nil && count < 2; v, _ = ss.Next() {
			count++
		}

		a := ss.Any(func(value2 interface{}) bool {
			counter++
			return value2 == 100 && value1 == 100
		})
		return a
	})

	if counter != 9 {
		panic("SortedSet_FirstInNestedAsk failed")
	}
}

func TestSortedSet_RandomValue(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	if ss.RandomValue() == nil {
		panic("SortedSet_RandomValue failed")
	}
}

func TestSortedSet_RandomValueWhere(t *testing.T) {
	ss := NewSortedSet(1, 2, 3)
	if ss.RandomValueWhere(func(value interface{}) bool {
		return value.(int) > 1 && value.(int) < 3
	}) != 2 {
		panic("SortedSet_RandomValueWhere failed")
	}
}
