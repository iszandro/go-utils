package utils

import (
	"errors"
	"reflect"
)

var (
	NotSliceErr      = errors.New("collection value is not a Slice.")
	NilMapFuncErr    = errors.New("map function is nil.")
	NilSelectFuncErr = errors.New("select function is nil.")
	ElemNotFoundErr  = errors.New("element not found.")
)

// MapFunc is the function to be called by Map.
// It receives each element of the collection and returns the value which will be
// added in the new slice.
type MapFunc func(obj interface{}) interface{}

// SelectFunc is the function to be called by Select.
// It receives each element of the collection and returns bool value. If true is returned,
// the element will be added to the returned slice.
type SelectFunc func(obj interface{}) bool

// Combination returns a slice of all combinations of elements from all slices.
// If slices is not a slice, then NotSliceErr is returned.
func Combination(slices ...interface{}) ([]interface{}, error) {
	slicesValue := reflect.ValueOf(slices)
	if slicesValue.Kind() != reflect.Slice {
		return make([]interface{}, 0), NotSliceErr
	}

	combinations := make([]interface{}, 0)
	singleCombination := make([]interface{}, slicesValue.Len())

	err := combination(singleCombination, &combinations, 0, &slicesValue)
	return combinations, err
}

// Compact returns a copy of the specified collection with all nil elements removed.
// If slices is not a slice, then NotSliceErr is returned.
func Compact(collection interface{}) ([]interface{}, error) {
	collectionValue := reflect.ValueOf(collection)
	if collectionValue.Kind() != reflect.Slice {
		return make([]interface{}, 0), NotSliceErr
	}

	compact := make([]interface{}, 0, collectionValue.Len())
	for i := 0; i < collectionValue.Len(); i++ {
		if item := collectionValue.Index(i).Interface(); item != nil {
			compact = append(compact, item)
		}
	}

	return compact, nil
}

// IsIncluded returns true if the specified element is present in the specified collection, otherwise returns false.
// If collection is not a slice, then NotSliceErr is returned.
// If element is not found, then ElemNotFoundErr is returned.
func IsIncluded(collection interface{}, obj interface{}) (bool, error) {
	collectionValue := reflect.ValueOf(collection)
	if collectionValue.Kind() != reflect.Slice {
		return false, NotSliceErr
	}

	for i := 0; i < collectionValue.Len(); i++ {
		if item := collectionValue.Index(i).Interface(); reflect.DeepEqual(item, obj) {
			return true, nil
		}
	}

	return false, ElemNotFoundErr
}

// Map calls the specified mapFunc once for each element in the collection.
// It returns a new slice containing the values returned by the mapFunc.
// If collection is not a slice, then NotSliceErr is returned.
// If mapFunc is nil, then NilMapFuncErr is returned.
func Map(collection interface{}, mapFunc MapFunc) ([]interface{}, error) {
	collectionValue := reflect.ValueOf(collection)
	if collectionValue.Kind() != reflect.Slice {
		return make([]interface{}, 0), NotSliceErr
	}

	if mapFunc == nil {
		return make([]interface{}, 0), NilMapFuncErr
	}

	newColl := make([]interface{}, collectionValue.Len())
	for i := 0; i < collectionValue.Len(); i++ {
		newColl[i] = mapFunc(collectionValue.Index(i).Interface())
	}

	return newColl, nil
}

// Select calls the specified selectFunc once for each element in the collection.
// It returns a new slice containing all elements of the collection for which the specified selectFunc returns true.
// If collection is not a slice, then NotSliceErr is returned.
// If mapFunc is nil, then NilSelectFuncErr is returned.
func Select(collection interface{}, selectFunc SelectFunc) ([]interface{}, error) {
	collectionValue := reflect.ValueOf(collection)
	if collectionValue.Kind() != reflect.Slice {
		return make([]interface{}, 0), NotSliceErr
	}

	if selectFunc == nil {
		return make([]interface{}, 0), NilSelectFuncErr
	}

	newColl := make([]interface{}, 0, collectionValue.Len())
	for i := 0; i < collectionValue.Len(); i++ {
		if obj := collectionValue.Index(i).Interface(); selectFunc(obj) {
			newColl = append(newColl, obj)
		}
	}

	return newColl, nil
}

func combination(singleCombination []interface{}, combinations *[]interface{}, depthLevel int, slicesValue *reflect.Value) error {
	slice := slicesValue.Index(depthLevel).Interface()
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Slice {
		return NotSliceErr
	}

	for i := 0; i < sliceValue.Len(); i++ {
		singleCombination[depthLevel] = sliceValue.Index(i).Interface()

		if nextLevel := depthLevel + 1; nextLevel < slicesValue.Len() {
			combination(singleCombination, combinations, nextLevel, slicesValue)
		} else {
			*combinations = append(*combinations, append([]interface{}{}, singleCombination...))
		}
	}

	return nil
}
