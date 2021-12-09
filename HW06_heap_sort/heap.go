package HW06_heap_sort

func HeapSort(arr []int) {
	for root := len(arr)/2 - 1; root >= 0; root-- {
		heapify(arr, root, len(arr))
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func heapify(arr []int, root, size int) {
	l, r := root*2+1, root*2+2
	max := root

	if l < size && arr[l] > arr[max] {
		max = l
	}

	if r < size && arr[r] > arr[max] {
		max = r
	}

	if max != root {
		arr[max], arr[root] = arr[root], arr[max]
		heapify(arr, max, size)
	}
}
