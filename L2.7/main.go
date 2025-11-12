// L2.7
// Что выведет программа?
// Числа от 1 до 8 в случайном порядке
// Объяснить работу конвейера с использованием select.
// Select ожидает данные от одного из каналов - при этом один case не блокирует другие
// если несколько каналов готовы, то выбирает случайный
// если ни один канал не готов, блокируется до готовности хотя бы одного
// если канал закрыт, то чтение из закрытого канала немедленно возвращает zero-value,
// его отсеивает по ok (второму возвращаемому значению)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Print(v)
	}
}
