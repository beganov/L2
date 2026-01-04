// L2.9
//
// Написать функцию Go, осуществляющую примитивную распаковку строки, содержащей повторяющиеся символы/руны.
// Примеры работы функции:
// Вход: "a4bc2d5e"
// Выход: "aaaabccddddde"
//
// Вход: "abcd"
// Выход: "abcd" (нет цифр — ничего не меняется)
//
// Вход: "45"
// Выход: "" (некорректная строка, т.к. в строке только цифры — функция должна вернуть ошибку)
//
// Вход: ""
// Выход: "" (пустая строка -> пустая строка)
//
// Дополнительное задание
// Поддерживать escape-последовательности вида \:
// Вход: "qwe\4\5"
// Выход: "qwe45" (4 и 5 не трактуются как числа, т.к. экранированы)
//
// Вход: "qwe\45"
// Выход: "qwe44444" (\4 экранирует 4, поэтому распаковывается только 5)
//
// Требования к реализации
// Функция должна корректно обрабатывать ошибочные случаи (возвращать ошибку, например, через error),
// и проходить unit-тесты.
// Код должен быть статически анализируем (vet, golint).

package unpack

import (
	"errors"
	"strconv"
)

var (
	ErrOnlyNumbers       = errors.New("string contains only numbers")
	ErrFirstSymbolNumber = errors.New("first symbol is a number")
	ErrLastSlash         = errors.New("escape character at the end of string")
)

func StrUnpack(in string) (string, error) {
	slashFlag := false
	firstSymbolFlag := false
	inRune := []rune(in)
	outRune := make([]rune, 0, len(inRune))
	for j := 0; j < len(inRune); j++ {
		// Обработка обратного слеша
		if inRune[j] == '\\' && !slashFlag {
			slashFlag = true
			//второй обратный слэш записыается
			continue
		}
		// Если число и не экранировано
		_, err := strconv.Atoi(string(inRune[j]))
		if err == nil && !slashFlag {
			p := j
			for err == nil && j < len(inRune)-1 {
				// Пропускаем все числа
				j++
				_, err = strconv.Atoi(string(inRune[j]))
			}
			if err == nil {
				j++
			}
			z, _ := strconv.Atoi(string(inRune[p:j]))
			if len(outRune) > 0 { // Если были символы в строке
				for k := 0; k < z-1; k++ { // распаковываем, дублируя предыдущий символ
					outRune = append(outRune, outRune[len(outRune)-1])
				}
				if z == 0 { // Если 0 - распаковываем как отсутствие символа
					outRune = outRune[:len(outRune)-1]
				}
			} else {
				// Если символом первым идет число
				firstSymbolFlag = true
			}
			j--
		} else {
			// Если не число и не первый обратный слеш
			outRune = append(outRune, inRune[j])
			slashFlag = false
		}
	}
	// Ошибка, если в строке только цифры
	if len(outRune) == 0 && len(inRune) != 0 && firstSymbolFlag {
		return "", ErrOnlyNumbers
	}
	// Ошибка, если первый символ - число
	if firstSymbolFlag {
		return "", ErrFirstSymbolNumber
	}
	// Ошибка, если в конце -обратный слэш
	if slashFlag {
		return "", ErrLastSlash
	}
	return string(outRune), nil
}
