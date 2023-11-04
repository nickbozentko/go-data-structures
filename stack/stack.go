package stack

type stack[T any] struct {
	stack []T
}

func NewStack[T any]() stack[T] {
	return stack[T]{
		stack: make([]T, 0),
	}
}

func (s *stack[T]) Push(e T) {
	s.stack = append(s.stack, e)
}

func (s *stack[T]) Pop() (T, bool) {
	if len(s.stack) == 0 {
		var empty T
		return empty, false
	}
	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return e, true
}

func (s *stack[T]) Peek() (T, bool) {
	if len(s.stack) == 0 {
		var empty T
		return empty, false
	}
	e := s.stack[len(s.stack)-1]
	return e, true
}

func (s *stack[T]) Length() int {
	return len(s.stack)
}
