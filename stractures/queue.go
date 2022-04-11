package structure

type Queue[T any] struct {
	size  int
	field []T
}

func NewEmptyQueue[T any]() *Queue[T] {
	return &Queue[T]{size: 0, field: make([]T, 0, 0)}
}

func (q *Queue[T]) Enqueue(obj T) {
	q.field = append(q.field, obj)
	q.size++
}

func (q *Queue[T]) Dequeue() (obj T) {
	if q.IsEmpty() {
		var nilVal T
		return nilVal
	}
	res := q.field[0]
	q.field = q.field[1:]
	q.size--
	return res
}

func (q *Queue[T]) Front() (obj T) {
	if q.IsEmpty() {
		var nilVal T
		return nilVal
	}
	return q.field[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
