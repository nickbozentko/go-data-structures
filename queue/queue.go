package queue

type queue[T any] struct {
	queue []T
}

func NewQueue[T any]() queue[T] {
	return queue[T]{
		queue: []T{},
	}
}

// @TODO
func (q *queue[T]) Add(e T) {

}

// @TOOD
func (q *queue[T]) Poll() (T, bool) {
	var r T
	return r, false
}
