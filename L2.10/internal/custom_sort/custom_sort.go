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

package custom_sort

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IndexedInput struct {
	valueModified string
	value         string
}

func Init(n, r, u, b, m, c, h bool, k int, input []string) []string {
	ii := make([]IndexedInput, 0, len(input))
	output := make([]string, 0, len(input))
	for _, j := range input {
		ii = append(ii, IndexedInput{valueModified: j, value: j})
	}
	if n {
		ii = CustomNumericSort(ii)
	} else {
		ii = CustomSort(ii)
	}
	if k != 1 {
		ii = CustomCollumnSort(ii, k)
	}
	if b {
		ii = CustomTrailingBlanksSort(ii)
	}
	if u {
		ii = CustomUnicSort(ii)
	}
	if m {
		ii = CustomMonthSort(ii)
	}
	if h {
		ii = CustomHumanSort(ii)
		n = true
	}
	if n {
		ii = CustomNumericSort(ii)
	} else {
		ii = CustomSort(ii)
	}
	if r {
		ii = CustomReverseSort(ii)
	}
	for i, j := range ii {
		if c {
			if input[i] != j.value {
				return []string{"Строки не отсортированы"}
			}
		} else {
			output = append(output, j.value)
		}
	}
	if c && (len(ii) == len(input)) {
		return []string{""}
	}
	return output
}

func CustomSort(input []IndexedInput) []IndexedInput {
	slices.SortFunc(input, func(a IndexedInput, b IndexedInput) int {
		if a.valueModified == b.valueModified {
			return 0
		}
		if a.valueModified < b.valueModified {
			return -1
		}
		return 1
	})
	return input
}

func CustomReverseSort(input []IndexedInput) []IndexedInput {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
	return input
}

func CustomHumanSort(input []IndexedInput) []IndexedInput {
	for i := range input {
		input[i].valueModified = strings.Replace(input[i].valueModified, "T", "KKKK", -1)
		input[i].valueModified = strings.Replace(input[i].valueModified, "G", "KKK", -1)
		input[i].valueModified = strings.Replace(input[i].valueModified, "M", "KK", -1)
		str := strings.Replace(input[i].valueModified, "K", "", -1)
		_, err := strconv.Atoi(str)
		if err != nil {
			float, err2 := strconv.ParseFloat(str, 64)
			if err2 != nil {
				input[i].valueModified = ""
			}
			var pow float64 = 1
			for j := 0; j < strings.Count(input[i].valueModified, "K"); j++ {
				pow = pow * 1000
			}
			input[i].valueModified = fmt.Sprintf("%v", float*pow)
		} else {
			input[i].valueModified = strings.Replace(input[i].valueModified, "K", "000", -1)
		}
	}
	return CustomNumericSort(input)
}

func CustomUnicSort(input []IndexedInput) []IndexedInput {
	UnicMap := make(map[IndexedInput]int, len(input))
	UnicOut := make([]IndexedInput, 0, len(UnicMap))
	for _, j := range input {
		if _, ok := UnicMap[j]; !ok {
			UnicMap[j]++
			UnicOut = append(UnicOut, j)
		}
	}
	return UnicOut
}

func CustomNumericSort(input []IndexedInput) []IndexedInput {
	slices.SortFunc(input, func(a IndexedInput, b IndexedInput) int {
		aValue, okA := new(big.Float).SetString(a.valueModified)
		bValue, okB := new(big.Float).SetString(b.valueModified)
		if !okA && !okB {
			return 0
		}
		if !okA {
			return -1
		}
		if !okB {
			return 1
		}
		if aValue == bValue {
			return 0
		}
		return aValue.Cmp(bValue)
	})
	return input
}

func CustomMonthSort(input []IndexedInput) []IndexedInput {
	for i, j := range input {
		input[i].valueModified = Monthed(j.valueModified)
	}
	return input
}

func CustomCollumnSort(input []IndexedInput, n int) []IndexedInput {
	n--
	if n < 0 {
		return input
	}
	for i, j := range input {
		split := strings.Split(j.valueModified, "\t")
		if n >= len(split) {
			input[i].valueModified = ""
		} else {
			input[i].valueModified = split[n]
		}
	}
	return input
}

func CustomTrailingBlanksSort(input []IndexedInput) []IndexedInput {
	for i, j := range input {
		input[i].valueModified = strings.TrimRight(j.valueModified, " ")
	}
	return input
}

func Monthed(input string) string {
	month := []rune(strings.ToUpper(input))[0:3]
	switch month[0] {
	case 'J':
		if month[1] == 'A' && month[2] == 'N' {
			return "1"
		}
		if month[1] == 'U' && month[2] == 'N' {
			return "6"
		}
		if month[1] == 'U' && month[2] == 'L' {
			return "7"
		}
	case 'F':
		if month[1] == 'E' && month[2] == 'B' {
			return "2"
		}
	case 'M':
		if month[1] == 'A' && month[2] == 'R' {
			return "3"
		}
		if month[1] == 'A' && month[2] == 'Y' {
			return "5"
		}
	case 'A':
		if month[1] == 'P' && month[2] == 'R' {
			return "4"
		}
		if month[1] == 'U' && month[2] == 'G' {
			return "8"
		}
	case 'S':
		if month[1] == 'E' && month[2] == 'P' {
			return "9"
		}
	case 'O':
		if month[1] == 'C' && month[2] == 'T' {
			return "A"
		}
	case 'N':
		if month[1] == 'O' && month[2] == 'V' {
			return "B"
		}
	case 'D':
		if month[1] == 'E' && month[2] == 'C' {
			return "C"
		}
	}
	return ""
}

func Input(inputPath string) ([]string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	output := make([]string, 0, 32)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}
	return output, nil
}

func STDIN(input []string) []string {
	sc := bufio.NewScanner(os.Stdin)
	output := make([]string, 0, 32)
	for sc.Scan() {
		line := sc.Text()
		output = append(output, line)
	}
	return output
}
