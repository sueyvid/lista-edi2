package list

import (
	"fmt"
	"testing"
)

var size int
var lists [1]IList

func createList(size int) {
	arraylist := &ArrayList{}
	(*arraylist).Init(size)
	lists = [1]IList{arraylist}
}

func deleteList() {
	lists[0] = nil
}

func SetupTest() func() {
	size = 10
	createList(size)

	return func() {
		deleteList()
	}
}

func TestAdd(t *testing.T) {
	defer SetupTest()()
	for _, list := range lists {
		for i := 0; i < size; i++ {
			list.AddEnd(i)
			if list.Size() != i+1 {
				t.Errorf("Tamanho de %T é %d, mas nós esperamos que ele seja %d", list, list.Size(), i+1)
			}
		}
		fmt.Println(list)
	}
}
