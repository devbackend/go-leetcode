package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblemMinStack(t *testing.T) {
	t.Parallel()

	t.Run("example 1", func(t *testing.T) {
		ms := Constructor()

		ms.Push(-2)
		ms.Push(0)
		ms.Push(-3)

		require.Equal(t, -3, ms.GetMin())

		ms.Pop()

		require.Equal(t, 0, ms.Top())
		require.Equal(t, -2, ms.GetMin())
	})

	t.Run("example 2", func(t *testing.T) {
		ms := Constructor()

		ms.Push(0)
		ms.Push(1)
		ms.Push(0)

		require.Equal(t, 0, ms.GetMin())

		ms.Pop()

		require.Equal(t, 0, ms.GetMin())
	})
}
