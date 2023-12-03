package collections

type Stack[T comparable] []T

func (s *Stack[T]) Pop() T {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}
func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}
func (s Stack[T]) Peek() T {
	return s[len(s)-1]
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}
func (s *Set[T]) Add(value T) {
	(*s)[value] = struct{}{}
}
func (s *Set[T]) Remove(value T) {
	delete(*s, value)
}
func (s Set[T]) Intersection(otherSet Set[T]) []T {
	result := make([]T, 0)
	for v := range s {
		if otherSet.Has(v) {
			result = append(result, v)
		}
	}
	return result
}

type Number interface {
	int | float32 | float64 | int32 | int64
}

func Sum[T Number](list []T) T {
	sum := T(0)
	for _, n := range list {
		sum += n
	}
	return sum
}
