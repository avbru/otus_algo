package HW05_data_structures

type Arr interface {
	Get(idx int) int
	Append(int)
	Add(n int, idx int)
	Delete(idx int) int
}

type Array struct {
	arr        []int
	resizeFunc func(*[]int)
}

func resizeSA(a *[]int) {
	if len(*a) < cap(*a) {
		return
	}

	newArray := make([]int, len(*a), len(*a)+1)
	copy(newArray, *a)
	*a = newArray
}

func resizeVector(a *[]int) {
	const factor = 10
	if len(*a) < cap(*a) {
		return
	}

	newArray := make([]int, len(*a), len(*a)+factor)
	copy(newArray, *a)
	*a = newArray
}

func resizeFactor(a *[]int) {
	const factor = 2

	if len(*a) == 0 {
		newArray := make([]int, 0, 1)
		*a = newArray
	}

	if len(*a) < cap(*a) {
		return
	}

	newArray := make([]int, len(*a), len(*a)*factor)
	copy(newArray, *a)
	*a = newArray
}

func (a *Array) Append(n int) {
	a.resizeFunc(&a.arr)
	a.arr = append(a.arr, n)
}

func (a *Array) Get(idx int) int {
	return a.arr[idx]
}

func (a *Array) Delete(idx int) int {
	el := a.arr[idx]
	if len(a.arr) > 1 {
		a.arr = append(a.arr[:idx], a.arr[idx+1:]...)
	} else {
		a.arr = a.arr[0:0]
	}

	return el
}

func (a *Array) Add(n, idx int) {
	a.resizeFunc(&a.arr)
	a.arr = append(a.arr[:idx+1], a.arr[idx:]...)
	a.arr[idx] = n
}
