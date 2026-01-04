//L2.2
// Что выведет программа?
// 2
// 1
// Объяснить порядок выполнения defer функций и итоговый вывод.
// defer в обоих случаях  вызывается при выходе из функции,
// однако в случае функции anotherTest при return x
// создается копия текущего значения
// т.о. инкрементирование в defer не влияет на вывод;
// return в test() (x int) возвращает сам параметр, не копию,
// так что изменение x через defer отражается в выводе

package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
