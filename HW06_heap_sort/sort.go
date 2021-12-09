package HW06_heap_sort

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func ShellSort(arr []int) {
	dist := 2
	for gap := len(arr) / dist; gap > 0; gap = gap / dist {
		// Insertion sort i+gap
		for i := 0; i+gap < len(arr); i++ {
			j := i + gap
			tmp := arr[j]

			for ; j-gap >= 0 && arr[j-gap] > tmp; j = j - gap {
				arr[j] = arr[j-gap]
			}

			arr[j] = tmp
		}
	}
}

// ShellSortFrank O(N^3/2)
func ShellSortFrank(arr []int) {
	for k := 1; ; k++ {
		gap := 2*(len(arr)/(2<<k)) + 1
		// Insertion sort i+gap
		for i := 0; i+gap < len(arr); i++ {
			j := i + gap
			tmp := arr[j]

			for ; j-gap >= 0 && arr[j-gap] > tmp; j = j - gap {
				arr[j] = arr[j-gap]
			}

			arr[j] = tmp
		}

		if gap <= 1 {
			return
		}
	}
}

// ShellSortSedgewick - O(N^4/3)
func ShellSortSedgewick(arr []int) {
	//1,8,23,77,281
	gaps := []int{1}
	if len(arr) < 46 {
		goto sort
	}
	gaps = append(gaps, 8)

	for k := 2; ; k++ {
		gap := 2<<(2*k-1) + 3*(2<<(k-2)) + 1
		if gap >= len(arr)/2 {
			break
		}
		gaps = append([]int{gap}, gaps...)
	}

sort:
	for _, gap := range gaps {
		for i := 0; i+gap < len(arr); i++ {
			j := i + gap
			tmp := arr[j]

			for ; j-gap >= 0 && arr[j-gap] > tmp; j = j - gap {
				arr[j] = arr[j-gap]
			}

			arr[j] = tmp
		}
	}
}
