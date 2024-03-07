package list

import (
	"fmt"
  	"errors"
)

type No struct {
	value int
	next *No
}

type linkedList struct {
	head *No
	size int
}

func (l *linkedList) eh_vazia() bool {
	if l.size == 0 {
		return true
	} else {
		return false
	}
}

func (l *linkedList) AddEnd(value int) {
	no := No{value, nil}
	pont := l.head
	prev := pont

	if l.eh_vazia() {
		l.head = &no
	} else {
		for pont != nil {
			prev = pont
			pont = pont.next
		}
		prev.next = &no
	}
	l.size++
}

func (l *linkedList) index_dentro(index int) bool {
	if index <= l.size-1 {
		return true
	} else {
		return false
	}
}

func (l *linkedList) GetByIndex(index int) (int, error) {
	if !l.index_dentro(index) {
		return 0, errors.New("Index fora da lista")
	}

	pont := l.head
	prev := pont
	for i := 0; i <= index; i++ {
		prev = pont
		pont = pont.next
	}
	return prev.value, errors.New("")
}
