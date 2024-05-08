package ds

type Queue[T any] struct {
	Items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Add(item T) {
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.Items) == 0 {
		return *new(T), false
	}

	item := q.Items[0]
	q.Items = q.Items[1:]

	return item, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if len(q.Items) == 0 {
		return *new(T), false
	}

	return q.Items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.Items)
}
