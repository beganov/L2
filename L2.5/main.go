// L2.5
// Что выведет программа?
// error
// Объяснить вывод программы.
// Так как err - непустой интерфейс, то он не равен nil
// В Go непустой интерфейс хранит указатель на itab (в нем как минимум есть Inter
// с описанием интерфейса и его методов и fun с указателем на методы, так что если
// реализуется хоть какой то метод, то равенства с nil не будет) и на data.

package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
