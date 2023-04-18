package datastructures

const deleted = -1 // special value to mark deleted elements

// Array defines our array structure
type Array struct {
	data       []int
	size       int
	deletedCnt int
}

// Append a new value to our array
func (a *Array) Append(val int) {
	if a.deletedCnt > 0 {
		for i, v := range a.data {
			if v == deleted {
				a.data[i] = val
				a.deletedCnt--
				return
			}
		}
	}
	if a.size == len(a.data) {
		newData := make([]int, a.size*2)
		copy(newData, a.data)
		a.data = newData
	}
	a.data[a.size] = val
	a.size++
}

// Get a value based on the index
func (a *Array) Get(index int) (int, bool) {
	if index >= 0 && index < a.size {
		if a.data[index] != deleted {
			return a.data[index], true
		}
	}
	return 0, false
}

// Set a value onto the index
func (a *Array) Set(index, val int) bool {
	if index >= 0 && index < a.size {
		if a.data[index] != deleted {
			a.data[index] = val
			return true
		}
	}
	return false
}

// Remove implemented using a "lazy deletion" strategy. Instead of removing a value from the underlying slice and
// shifting all subsequent values left, we can simply mark the value as deleted by setting it to a special "deleted"
// value, and keep track of the number of deleted values separately. This allows us to maintain the original order of
// the values in the array while still freeing up space when necessary.
func (a *Array) Remove(index int) bool {
	if index < 0 || index >= a.size {
		return false
	}
	if a.data[index] != deleted {
		a.data[index] = deleted
		a.deletedCnt++
	}
	if a.deletedCnt > a.size/4 {
		newData := make([]int, a.size)
		j := 0
		for _, v := range a.data {
			if v != deleted {
				newData[j] = v
				j++
			}
		}
		a.data = newData
		a.size = j
		a.deletedCnt = 0
	}
	return true
}

// Size returns the array size
func (a *Array) Size() int {
	return a.size
}
