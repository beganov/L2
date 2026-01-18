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

	fmt.Println(normalisedArgs)
	var err error
	err = flag.CommandLine.Parse(normalisedArgs)
	if err != nil {
		fmt.Println(normalisedArgs)
	}
	var inputArray []string
	if len(flag.Args()) > 0 {
		inputArray, err = custom_sort.Input(flag.Args()[0])
		if err != nil {
			inputArray = custom_sort.STDIN(flag.Args())
		}
	} else {
		inputArray = custom_sort.STDIN(flag.Args())
	}
	fmt.Printf("n=%v, r=%v, u=%v, b=%v, m=%v, c=%v, h=%v, k=%d\n", *numeric, *reverse, *unique, *blank, *month, *check, *human, *column)
	fmt.Println(custom_sort.Init(*numeric, *reverse, *unique, *blank, *month, *check, *human, *column, inputArray))
}
