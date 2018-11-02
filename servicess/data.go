package servicess

import (
	"fmt"
)

type data struct {
	name string
}

func (w *data) say() {
	fmt.Println("this is data")
}
