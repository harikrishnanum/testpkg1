package foo

import (
	"fmt"

	"github.com/harikrishnanum/testpkg1/config"
)

func Foo() {
	fmt.Printf("Key1: %s\n", config.Conf.Key1)
}
