package containers //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, l.Len(), 0)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push front", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)

		require.Equal(t, l.Front(), l.Back())
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Front().Next)
		require.Equal(t, l.Front().Value, 10)

		l.PushFront(20) //[20,10]
		require.Nil(t, l.Front().Next)
		require.NotNil(t, l.Front().Prev)

		require.Nil(t, l.Back().Prev)
		require.NotNil(t, l.Back().Next)

		require.Equal(t, l.Front().Value, 20)
		require.Equal(t, l.Back().Value, 10)

		front := *l.Front()
		back := *l.Back()

		require.Equal(t, front.Prev, l.Back())
		require.Equal(t, back.Next, l.Front())

	})

	t.Run("push back", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)
		require.Nil(t, l.Back().Prev)
		require.Nil(t, l.Back().Next)
		require.Equal(t, l.Back().Value, 10)
		require.Equal(t, l.Front().Value, 10)

		l.PushBack(20) //[10, 20]
		require.NotNil(t, l.Back().Next)
		require.Nil(t, l.Back().Prev)

		require.Equal(t, l.Back().Value, 20)
		require.Equal(t, l.Front().Value, 10)
	})

	t.Run("remove", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.Remove(l.Back())
		require.Equal(t, l.Len(), 0)

		l.PushFront(20)
		l.PushFront(10)    //[10,20]
		l.Remove(l.Back()) //[10]
		require.Equal(t, l.Len(), 1)
		require.Equal(t, l.Front(), l.Back())
		require.Equal(t, l.Front().Value, 10)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Front().Next)

		l.PushBack(20)
		l.PushBack(30)          //[10,20,30]
		l.Remove(l.Back().Next) //[10,30]
		require.Equal(t, l.Len(), 2)
		require.Equal(t, l.Front().Value, 10)
		require.Equal(t, l.Back().Value, 30)

		l.Remove(l.Front()) //[30]
		require.Equal(t, l.Len(), 1)
		require.Equal(t, l.Front().Value, 30)
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.MoveToFront(l.Front())
		require.Equal(t, l.Len(), 1)
		require.Nil(t, l.Front().Next, 1)
		require.Nil(t, l.Front().Prev, 1)
		require.Equal(t, l.Front().Value, 10)

		l.PushBack(20)          //[10,20]
		l.MoveToFront(l.Back()) //[20,10]
		require.Equal(t, l.Front().Value, 20)
		require.Equal(t, l.Back().Value, 10)
		require.Nil(t, l.Front().Next)

		l.PushBack(30)               //[20,10,30]
		l.MoveToFront(l.Back().Next) //[10,20,30]
		require.Equal(t, l.Front().Value, 10)
		require.Equal(t, l.Front().Prev.Value, 20)
		require.Equal(t, l.Back().Value, 30)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, l.Len(), 3)

		middle := l.Back().Next // 20
		l.Remove(middle)        // [10, 30]
		require.Equal(t, l.Len(), 2)

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, l.Len(), 7)
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Back(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, elems)
	})

}
