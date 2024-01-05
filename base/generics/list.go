package generics

type List[T any] interface {
	Add(index int, val T)
	Append(val T) error
	Delete(index int) error
}

type LinkedList[T any] struct {
	head *node[T]
}

type node[T any] struct {
	data T
}

func (l *LinkedList[T]) Add(index int, val T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Append(val T) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

func UseList() {
	l := &LinkedList[string]{}
	//l.Add(1, 123)
	l.Add(1, "123")
	//l.Add(1, 123) //err

}
