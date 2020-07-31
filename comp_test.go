package slicecomp

import "testing"


func TestEquals(t *testing.T) {
    arr1 := []int{1, 2, 3}
    arr2 := []int{1, 2, 3}
    equals := Equals(arr1, arr2)
    if ! equals {
        t.Error("FAILURE: Equals returned false for two equal slices.")
    }
}


func TestNotEquals(t *testing.T) {
    arr1 := []int{1, 2, 3}
    arr2 := []int{1, 3, 5}
    equals := Equals(arr1, arr2)
    if equals {
        t.Error("FAILURE: Equals returned true for two non-equal slices.")
    }
}


func TestDifferentLengths(t *testing.T) {
    arr1 := []int{1, 2, 3}
    arr2 := []int{1, 2, 3, 4, 5}
    equals := Equals(arr1, arr2)
    if equals {
        t.Error("FAILURE: Equals returned true for two different-length slices.")
    }
}


func TestNil(t *testing.T) {
    var arr1 []int
    var arr2 []int
    equals := Equals(arr1, arr2)
    if ! equals {
        t.Error("FAILURE: Equals does not consider nil slices equal.")
    }
}
