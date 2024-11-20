package btree

import (
	"fmt"
	"strings"
)

type Tree[T any] struct {
	Left  *Tree[T]
	Right *Tree[T]
	Value T
}

func (t Tree[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("<value: %v, ", t.Value))
	sb.WriteString(fmt.Sprintf("left: %v, ", t.Left))
	sb.WriteString(fmt.Sprintf("right: %v>", t.Right))

	return sb.String()
}
