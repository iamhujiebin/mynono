package servicess

import (
	"fmt"
)

type worker struct {
	name string
}

func (w *worker) say() {
	fmt.Println("this is worker")
}
