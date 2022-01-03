package HW11_12_topology_sort

func Kosaraju(arr [][]int) (components [][]int) {
	list := Tarjan(reverse(arr))
	used := make([]bool, len(arr))
	for i := 0; i < len(list); {
		var res []int
		dfs(arr, list[i], &res, used)
		components = append(components, res)
		i = i + len(res)
	}

	return components
}

func dfs(arr [][]int, n int, res *[]int, used []bool) {
	if used[n] {
		return
	}
	used[n] = true

	for _, ver := range arr[n] {
		dfs(arr, ver, res, used)
	}

	*res = append(*res, n)
}

func reverse(arr [][]int) [][]int {
	rev := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			rev[arr[i][j]] = append(rev[arr[i][j]], i)
		}
	}
	return rev
}
