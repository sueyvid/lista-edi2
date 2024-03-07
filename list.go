package list

type IList interface {
	AddEnd(value int)
	GetByIndex(index int) (int, error)
	AddByIndex(index, value int) error
	Set(index, value int) error
	Size() int
	RemoveEnd() (int, error)
	RemoveByIndex(index int) (int, error)
}
