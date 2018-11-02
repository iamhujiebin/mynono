package servicess

import (
	"fmt"
)

type utils struct {
	name string
}

func (w *utils) say() {
	fmt.Println("this is utils")
}
