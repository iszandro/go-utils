package arraylist

import (
	"errors"
	"fmt"
	"reflect"
)

type ArrayList struct {
	slice []interface{}
}

// New returns a new *ArrayList
func New() *ArrayList {
	return new(ArrayList)
}

// Add appends the specified elements to the end of this list.
func (a *ArrayList) Add(objs ...interface{}) {
	a.slice = append(a.slice, objs...)
}

// AddAt inserts the specified elements at the specified position in this list.
// If pos is more than the list size or less than 0, then index out of range
// error is returned. Nil otherwise.
func (a *ArrayList) AddAt(pos int, objs ...interface{}) error {
	if err := a.checkRangeForAddAt(pos); err != nil {
		return err
	}

	switch pos {
	case 0:
		a.AddFirst(objs...)
		break
	case a.Size():
		a.Add(objs...)
		break
	default:
		a.addAt(pos, objs...)
	}

	return nil
}

// AddFirst inserts the specified elements to the beginning of this list.
func (a *ArrayList) AddFirst(objs ...interface{}) {
	a.slice = append(objs, a.slice...)
}

// Clear removes all of the elements from this list.
func (a *ArrayList) Clear() {
	a.slice = nil
}

// Get returns the element at the specified position in this list.
// It returns the element at the specified position if exists, otherwise returns nil.
// Can return index out of range error.
func (a *ArrayList) Get(pos int) (interface{}, error) {
	if err := a.checkRange(pos); err != nil {
		return nil, indexOutOfRangeErr(pos, a.Size())
	}

	return a.slice[pos], nil
}

// IndexOf returns the index (0-based) of the first occurrence of the specified element in this list.
// It can return -1 if this list does not contain the specified element.
func (a *ArrayList) IndexOf(obj interface{}) int {
	for i, o := range a.slice {
		if reflect.DeepEqual(o, obj) {
			return i
		}
	}

	return -1
}

// IsEmpty returns true if this list containes no elements.
func (a *ArrayList) IsEmpty() bool {
	return a.Size() == 0
}

// LastIndexOf returns the index (0-based) of the last occurrence of the specified element in this list.
// It can return -1 if this list does not contain the specified element.
func (a *ArrayList) LastIndexOf(obj interface{}) int {
	for i := a.Size() - 1; i > -1; i-- {
		if o := a.slice[i]; reflect.DeepEqual(o, obj) {
			return i
		}
	}

	return -1
}

// Remove removes the first occurrence of the specified element from this list.
// If element not found, it returns an element not found error.
func (a *ArrayList) Remove(obj interface{}) error {
	for i, o := range a.slice {
		if reflect.DeepEqual(o, obj) {
			return a.RemoveAt(i)
		}
	}

	return elementNotFoundErr(obj)
}

// RemoveAt removes the element at the specified position (0-based) in this list.
// It can return index out of range error.
func (a *ArrayList) RemoveAt(pos int) error {
	if err := a.checkRange(pos); err != nil {
		return err
	}

	a.slice[pos] = nil
	a.slice = append(a.slice[:pos], a.slice[pos+1:]...)
	return nil
}

// Size returns the number of elements in this list.
func (a *ArrayList) Size() int {
	return len(a.slice)
}

// Slice returns a slice containing all of the elements in this list.
// To avoid references, the returned slice is a copy of this list.
func (a *ArrayList) Slice() []interface{} {
	return append([]interface{}{}, a.slice...)
}

func (a *ArrayList) addAt(pos int, elements ...interface{}) {
	a.slice = append(append(append([]interface{}{}, a.slice[:pos]...), elements...), a.slice[pos:]...)
}

func (a *ArrayList) checkRangeForAddAt(pos int) error {
	if pos > a.Size() || pos < 0 {
		return indexOutOfRangeErr(pos, a.Size())
	}

	return nil
}

func (a *ArrayList) checkRange(pos int) error {
	if pos > a.Size()-1 || pos < 0 {
		return indexOutOfRangeErr(pos, a.Size())
	}

	return nil
}

func elementNotFoundErr(obj interface{}) error {
	errStr := fmt.Sprintf("%v element was not found in this list.", obj)
	return errors.New(errStr)
}

func indexOutOfRangeErr(pos, listSize int) error {
	errStr := fmt.Sprintf("Index %d is out of range from a list size of %d", pos, listSize)
	return errors.New(errStr)
}
