package slicecomp

func Equals(this, other []int) bool {
    if len(this) != len(other) {
        return false
    }

    for i, v := range this {
        if v != other[i] {
            return false
        }
    }
    return true
}
