// L2.3
// Что выведет программа?
// nil
// false
// Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
// В Go непустой интерфейс хранит два указателя: на таблицу типа (itab) и на данные(data).
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

type empty interface {
	Foo()
}

func main() {
	err := Foo()
	var err2 empty
	fmt.Println(err)
	fmt.Println(err == nil)
	fmt.Println(err2 == nil)
}
