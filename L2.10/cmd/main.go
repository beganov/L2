// L2.10
// Реализовать упрощённый аналог UNIX-утилиты sort (сортировка строк).
// Программа должна читать строки (из файла или STDIN) и выводить их отсортированными.
// Обязательные флаги (как в GNU sort):
// -k N — сортировать по столбцу (колонке) №N (разделитель — табуляция по умолчанию).
// Например, «sort -k 2» отсортирует строки по второму столбцу каждой строки.
// -n — сортировать по числовому значению (строки интерпретируются как числа).
// -r — сортировать в обратном порядке (reverse).
// -u — не выводить повторяющиеся строки (только уникальные).
// Дополнительные флаги:
// -M — сортировать по названию месяца (Jan, Feb, ... Dec),
// т.е. распознавать специфический формат дат.
// -b — игнорировать хвостовые пробелы (trailing blanks).
// -c — проверить, отсортированы ли данные; если нет, вывести сообщение об этом.
// -h — сортировать по числовому значению с учётом суффиксов
//  (например, К = килобайт, М = мегабайт — человекочитаемые размеры).
// Программа должна корректно обрабатывать комбинации флагов
// (например, -nr — числовая сортировка в обратном порядке, и т.д.).
// Необходимо предусмотреть эффективную обработку больших файлов.
// Код должен проходить все тесты, а также проверки go vet и golint
// (понимание, что требуются надлежащие комментарии, имена и структура программы).

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/beganov/custom_sort/internal/custom_sort"
)

func main() {
	args := os.Args[1:]
	var normalisedArgs []string
	for _, arg := range args {
		kflag := false
		if !strings.HasPrefix(arg, "-") {
			normalisedArgs = append(normalisedArgs, arg)
			continue
		}
		for i, ch := range arg {
			if ch == '-' && i < 2 {
				continue
			}
			if ch == '=' {
				normalisedArgs = append(normalisedArgs, arg[i+1:])
				break
			}
			if ch == 'k' {
				kflag = true
				continue
			}
			if ch >= '0' && ch <= '9' && kflag {
				normalisedArgs = append(normalisedArgs, "-k")
				normalisedArgs = append(normalisedArgs, string(ch))
			} else {
				normalisedArgs = append(normalisedArgs, "-"+string(ch))
			}

		}
		if kflag {
			normalisedArgs = append(normalisedArgs, "-k")
		}
	}

	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	blank := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	month := flag.Bool("m", false, "сортировать по названию месяца")
	check := flag.Bool("c", false, "проверить, отсортированы ли данные")
	human := flag.Bool("h", false, "сортировать по числовому значению с учётом суффиксов")
	column := flag.Int("k", 1, "сортировать по столбцу №N")

	var err error
	err = flag.CommandLine.Parse(normalisedArgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка парсинга аргументов: %v\n", err)
		flag.Usage()
		os.Exit(1)

	}
	var inputArray []string
	if len(flag.Args()) > 0 {
		inputArray, err = custom_sort.Input(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка парсинга: %v\n", err)
			flag.Usage()
			os.Exit(1)
		}
	} else {
		fmt.Println("hehe")
		inputArray = custom_sort.STDIN(flag.Args())
	}
	fmt.Println(custom_sort.Init(*numeric, *reverse, *unique, *blank, *month, *check, *human, *column, inputArray))
}
