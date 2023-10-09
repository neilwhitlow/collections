package doublylinkedlist_test

import (
	"strings"
	"testing"

	dll "github.com/neilwhitlow/collections/doublylinkedlist"
	"github.com/stretchr/testify/assert"
)

func TestEnsuredInitialization(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var l dll.DoublyLinkedList[string]
		l.AddFirst("Zeta")
		assert.Equal(t, 1, l.Len(), "Length should be 1")
	})
}

func TestAddFirst(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		l.AddFirst("Zeta")
		l.AddLast("Alpha")
		l.AddFirst("Gamma")
		assert.Equal(t, 3, l.Len(), "Length should be 3")

		assert.Equal(t, "Gamma", l.First().Value, "Zeta should be first")
	})
}

func TestFirst(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		l.AddFirst("Zeta")
		l.AddLast("Alpha")
		assert.Equal(t, 2, l.Len(), "Length should be 2")

		assert.Equal(t, "Zeta", l.First().Value, "Zeta should be first")
	})
}

func TestAddLast(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		l.AddFirst("Zeta")
		l.AddLast("Alpha")
		l.AddLast("Delta")
		assert.Equal(t, 3, l.Len(), "Length should be 3")

		assert.Equal(t, "Delta", l.Last().Value, "Delta should be last")
	})
}

func TestLast(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		l.AddFirst("Zeta")
		l.AddLast("Alpha")
		assert.Equal(t, 2, l.Len(), "Length should be 2")

		assert.Equal(t, "Alpha", l.Last().Value, "Alpha should be last")
	})
}

func TestFirstNil(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		assert.Equal(t, (*dll.Node[string])(nil), l.First(), "First should be nil.")
	})
}

func TestLastNil(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		assert.Equal(t, (*dll.Node[string])(nil), l.Last(), "Last should be nil.")
	})
}

func TestContains(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		l.AddFirst("Zeta")
		l.AddFirst("Gamma")
		l.AddFirst("Beta")
		l.AddLast("Alpha")

		assert.True(t, l.Contains("Beta"), "Beta should be found")
		assert.False(t, l.Contains("Epsilon"), "Epsilon should not be found")
	})
}

func TestNextIteration(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		l.AddFirst("Zeta")
		l.AddFirst("Gamma")
		l.AddFirst("Beta")
		l.AddLast("Alpha")

		var sb strings.Builder
		sb.WriteString("[")
		for currentNode := l.First(); currentNode != nil; currentNode = currentNode.Next() {
			sb.WriteString(currentNode.Value)
			if currentNode.Next() != nil {
				sb.WriteString(",")
			}
		}
		sb.WriteString("]")
		expected := "[Beta,Gamma,Zeta,Alpha]"
		assert.Equal(t, expected, sb.String(), "Unexpected order")
	})
}

func TestPrevIteration(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := dll.New[string]()
		l.AddFirst("Zeta")
		l.AddFirst("Gamma")
		l.AddFirst("Delta")
		l.AddFirst("Beta")
		l.AddFirst("Alpha")

		var sb strings.Builder
		sb.WriteString("[")
		for currentNode := l.Last(); currentNode != nil; currentNode = currentNode.Prev() {
			sb.WriteString(currentNode.Value)
			if currentNode.Prev() != nil {
				sb.WriteString(",")
			}
		}
		sb.WriteString("]")
		expected := "[Zeta,Gamma,Delta,Beta,Alpha]"
		assert.Equal(t, expected, sb.String(), "Unexpected order")
	})
}
