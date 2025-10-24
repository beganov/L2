// L2.3
// Что выведет программа?
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// fatal error: all goroutines are asleep - deadlock!
// Объяснить вывод программы.
// Попытка чтения из пустого канала по завершении работы отправителя приводит к блокировке
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for n := range ch {
		println(n)
	}
}
