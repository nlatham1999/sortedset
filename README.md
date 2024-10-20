SortedSet is a set of elements that can be sorted with a provided sorter function.

Items added will not cause the set to be resorted. This must be done as an extra step.

# sortedset

```go
import "github.com/nlatham1999/sortedset"
```

## Example Usage

```go
	type coord struct {
		x, y int
	}
	ss := NewSortedSet(coord{1, 2}, coord{2, 3}, coord{3, 4})
	ss.SortDesc(func(e interface{}) interface{} {
		c := e.(coord)
		return c.x
	})
```

```go
	ss := NewSortedSet(1, 3)
	ss.InsertBefore(2, 3)
```

```go
	ss := NewSortedSet(1, 2, 3)
	next, _ := ss.Next(1)
```

## Index

- [Variables](<#variables>)
- [type SortedSet](<#SortedSet>)
  - [func NewSortedSet\(values ...interface\{\}\) \*SortedSet](<#NewSortedSet>)
  - [func \(ss \*SortedSet\) Add\(value interface\{\}\) error](<#SortedSet.Add>)
  - [func \(ss \*SortedSet\) After\(value interface\{\}\) \(interface\{\}, error\)](<#SortedSet.After>)
  - [func \(ss \*SortedSet\) Before\(value interface\{\}\) \(interface\{\}, error\)](<#SortedSet.Before>)
  - [func \(ss \*SortedSet\) Contains\(value interface\{\}\) bool](<#SortedSet.Contains>)
  - [func \(ss \*SortedSet\) Current\(\) interface\{\}](<#SortedSet.Current>)
  - [func \(ss \*SortedSet\) Difference\(set \*SortedSet\) \*SortedSet](<#SortedSet.Difference>)
  - [func \(ss \*SortedSet\) Empty\(\) bool](<#SortedSet.Empty>)
  - [func \(ss \*SortedSet\) First\(\) interface\{\}](<#SortedSet.First>)
  - [func \(ss \*SortedSet\) InsertAfter\(value, after interface\{\}\) error](<#SortedSet.InsertAfter>)
  - [func \(ss \*SortedSet\) InsertBefore\(value, before interface\{\}\) error](<#SortedSet.InsertBefore>)
  - [func \(ss \*SortedSet\) Intersection\(set \*SortedSet\) \*SortedSet](<#SortedSet.Intersection>)
  - [func \(ss \*SortedSet\) Last\(\) interface\{\}](<#SortedSet.Last>)
  - [func \(ss \*SortedSet\) Len\(\) int](<#SortedSet.Len>)
  - [func \(ss \*SortedSet\) List\(\) \[\]interface\{\}](<#SortedSet.List>)
  - [func \(ss \*SortedSet\) Next\(\) \(interface\{\}, error\)](<#SortedSet.Next>)
  - [func \(ss \*SortedSet\) Previous\(\) \(interface\{\}, error\)](<#SortedSet.Previous>)
  - [func \(ss \*SortedSet\) Remove\(value interface\{\}\) error](<#SortedSet.Remove>)
  - [func \(ss \*SortedSet\) SortAsc\(f func\(e interface\{\}\) interface\{\}\)](<#SortedSet.SortAsc>)
  - [func \(ss \*SortedSet\) SortDesc\(f func\(e interface\{\}\) interface\{\}\)](<#SortedSet.SortDesc>)
  - [func \(ss \*SortedSet\) SymmetricDifference\(set \*SortedSet\) \*SortedSet](<#SortedSet.SymmetricDifference>)
  - [func \(ss \*SortedSet\) Union\(set \*SortedSet\) \*SortedSet](<#SortedSet.Union>)


## Variables

<a name="ErrItemExists"></a>

```go
var (
    ErrItemExists      = errors.New("item exists")
    ErrItemDoesntExist = errors.New("item doesn't exist")
)
```

<a name="SortedSet"></a>
## type SortedSet



```go
type SortedSet struct {
    // contains filtered or unexported fields
}
```

<a name="NewSortedSet"></a>
### func NewSortedSet

```go
func NewSortedSet(values ...interface{}) *SortedSet
```

create a new sorted set with the given values

<a name="SortedSet.Add"></a>
### func \(\*SortedSet\) Add

```go
func (ss *SortedSet) Add(value interface{}) error
```

adds a value to the end of the set

<a name="SortedSet.After"></a>
### func \(\*SortedSet\) After

```go
func (ss *SortedSet) After(value interface{}) (interface{}, error)
```

returns the next value in the set after the given value

<a name="SortedSet.Before"></a>
### func \(\*SortedSet\) Before

```go
func (ss *SortedSet) Before(value interface{}) (interface{}, error)
```

returns the value before the given value

<a name="SortedSet.Contains"></a>
### func \(\*SortedSet\) Contains

```go
func (ss *SortedSet) Contains(value interface{}) bool
```

returns if the value exists in the set

<a name="SortedSet.Current"></a>
### func \(\*SortedSet\) Current

```go
func (ss *SortedSet) Current() interface{}
```

returns the current value of the pointer

<a name="SortedSet.Difference"></a>
### func \(\*SortedSet\) Difference

```go
func (ss *SortedSet) Difference(set *SortedSet) *SortedSet
```

returns the difference of this set and the given set as a new set \(this \- set\)

<a name="SortedSet.Empty"></a>
### func \(\*SortedSet\) Empty

```go
func (ss *SortedSet) Empty() bool
```

returns if the set is empty

<a name="SortedSet.First"></a>
### func \(\*SortedSet\) First

```go
func (ss *SortedSet) First() interface{}
```

returns the first value in the set and sets the pointer to it

<a name="SortedSet.InsertAfter"></a>
### func \(\*SortedSet\) InsertAfter

```go
func (ss *SortedSet) InsertAfter(value, after interface{}) error
```

inserts a value after another value

<a name="SortedSet.InsertBefore"></a>
### func \(\*SortedSet\) InsertBefore

```go
func (ss *SortedSet) InsertBefore(value, before interface{}) error
```

inserts a value before another value

<a name="SortedSet.Intersection"></a>
### func \(\*SortedSet\) Intersection

```go
func (ss *SortedSet) Intersection(set *SortedSet) *SortedSet
```

returns the intersection of this set and the given set as a new set

<a name="SortedSet.Last"></a>
### func \(\*SortedSet\) Last

```go
func (ss *SortedSet) Last() interface{}
```

returns the last value in the set and sets the pointer to it

<a name="SortedSet.Len"></a>
### func \(\*SortedSet\) Len

```go
func (ss *SortedSet) Len() int
```

returns the length of the set

<a name="SortedSet.List"></a>
### func \(\*SortedSet\) List

```go
func (ss *SortedSet) List() []interface{}
```

returns the set as a slice

<a name="SortedSet.Next"></a>
### func \(\*SortedSet\) Next

```go
func (ss *SortedSet) Next() (interface{}, error)
```

moves the pointer to the next value in the set and returns it

<a name="SortedSet.Previous"></a>
### func \(\*SortedSet\) Previous

```go
func (ss *SortedSet) Previous() (interface{}, error)
```

moves the pointer to the previous value in the set and returns it

<a name="SortedSet.Remove"></a>
### func \(\*SortedSet\) Remove

```go
func (ss *SortedSet) Remove(value interface{}) error
```

removes a value from the set

<a name="SortedSet.SortAsc"></a>
### func \(\*SortedSet\) SortAsc

```go
func (ss *SortedSet) SortAsc(f func(e interface{}) interface{})
```

sorts the set in ascending order

<a name="SortedSet.SortDesc"></a>
### func \(\*SortedSet\) SortDesc

```go
func (ss *SortedSet) SortDesc(f func(e interface{}) interface{})
```

sorts the set in descending order

<a name="SortedSet.SymmetricDifference"></a>
### func \(\*SortedSet\) SymmetricDifference

```go
func (ss *SortedSet) SymmetricDifference(set *SortedSet) *SortedSet
```

returns the symmetric difference of this set and the given set as a new set \(this XOR set\)

<a name="SortedSet.Union"></a>
### func \(\*SortedSet\) Union

```go
func (ss *SortedSet) Union(set *SortedSet) *SortedSet
```

returns the union of this set and the given set as a new set

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
