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


func BenchmarkEquals(b *testing.B) {
    var arr1 = make([]int, 1000000)
    var arr2 = make([]int, 1000000)

    for i := 0; i < 1000000; i++ {
        arr1[i] = i
        arr2[i] = i
    }
    Equals(arr1, arr2)
}


func BenchmarkDifferentLengths(b *testing.B) {
    var arr1 = make([]int, 999999)
    var arr2 = make([]int, 1000000)

    for i := 0; i < 999999; i++ {
        arr1[i] = i
    }

    for i := 0; i < 1000000; i++ {
        arr2[i] = i
    }
    Equals(arr1, arr2)
}


func BenchmarkWorstCase(b *testing.B) {
    var arr1 = make([]int, 1000000)
    var arr2 = make([]int, 1000000)

    for i := 0; i < 1000000; i++ {
        arr1[i] = i
        arr2[i] = i
    }
    arr2[len(arr2)-1] = 235  // assign the last index to a random int
    Equals(arr1, arr2)
}
