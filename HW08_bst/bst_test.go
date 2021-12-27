package HW08_bst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBST_Add(t *testing.T) {
	tr := NewBST()
	//Create
	tr.Insert(20)
	require.Equal(t, 20, tr.root.Val)
	require.Nil(t, tr.root.L, "L must be nil")
	require.Nil(t, tr.root.R, "R must be nil")

	//Insert Left
	tr.Insert(10)
	require.Equal(t, 10, tr.root.L.Val)
	require.NotNil(t, tr.root.L, "L must be not nil")
	require.Nil(t, tr.root.R, "R must be nil")

	//Insert Right
	tr.Insert(30)
	require.Equal(t, 30, tr.root.R.Val)
	require.NotNil(t, tr.root.R, "L must be not nil")
	require.NotNil(t, tr.root.L, "L must be not nil")

	//Search
	require.Nil(t, tr.Search(0))
	require.Equal(t, 10, tr.Search(10).Val)
	require.Equal(t, 20, tr.Search(20).Val)
	require.Equal(t, 30, tr.Search(30).Val)

	//TODO
	require.Equal(t, []int{10, 20, 30}, tr.List())
}

func TestBST_Draw(t *testing.T) {
	tr := NewBST()
	values := []int{40, 50, 60, 45, 20, 10, 36, 30, 22, 28, 12, 8}
	for _, v := range values {
		tr.Insert(v)
	}
	tr.Draw()
	println("---------------------------------------------------------------------------")
	tr.Delete(60)
	tr.Draw()
	println("---------------------------------------------------------------------------")
	tr.Delete(50)
	tr.Draw()
	println("---------------------------------------------------------------------------")
	tr.Delete(20)
	tr.Draw()
	require.IsIncreasing(t, tr.List())
}
