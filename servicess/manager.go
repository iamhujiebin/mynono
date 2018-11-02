package servicess

import (
	"fmt"
)

type manager struct {
	name string
}

func (w *manager) say() {
	fmt.Println("this is manager")
}
