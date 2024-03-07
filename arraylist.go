package list

import "errors"

type ArrayList struct {
	list []int
	size int
}

func (l *ArrayList) Init(size int) error {
	if size > 0 {
		l.list = make([]int, size)
		return nil
	} else {
		return errors.New("Não pode ser criado um ArrayList com tamanho <= 0")
	}
}

func (l *ArrayList) index_dentro(index int) bool {
	if index <= l.size-1 {
		return true
	}
	else {
		return false
	}
}

func (l *ArrayList) tem_espaco_disponivel() bool {
	if l.size < len(l.list) {
		return true
	}
	else {
		return false
	}
}

func (l *ArrayList) eh_vazia() bool {
	if l.size == 0 {
		return true
	}
	else {
		return false
	}
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) AddEnd(value int) {
	if l.tem_espaco_disponivel() {
		l.list[l.size] = value
		l.size++
	}
}

func (l *ArrayList) GetByIndex(index int) (int, error) {
	if not l.index_dentro(index) {
		return errors.New("Index fora da lista")
	}
	return l.list[index]
}

func (l *ArrayList) AddByIndex(index, value int) error {
	if not l.tem_espaco_disponivel() {
		return errors.New("Lista não tem espaço disponível")
	}

	for i := l.size; i > index; i-- {
		l.list[i] = l.list[i-1]
	}
	l.list[index] = value
	l.size++
}

func (l *ArrayList) Set(index, value int) error {
	if not l.index_dentro(index) {
		return errors.New("Index fora da lista")
	}
	l.list[index] = value
}

func (l *ArrayList) RemoveEnd() (int, error) {
	if l.eh_vazia() {
		return errors.New("Lista vazia, não há o que remover")
	}
	l.size--
}

func (l *ArrayList) RemoveByIndex(index int) (int, error) {
	if not index_dentro(index) {
		return errors.New("Index fora da lista")
	}

	for i := index; i < l.size-1; i++ {
		l.list[i] = l.list[i+1]
	}
	l.size--
}