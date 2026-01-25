package anagramms

import (
	"slices"
	"strings"
)

func GroupAnagramms(input []string) map[string][]string {
	resultMap := make(map[string][]string, len(input))
	for _, j := range input {
		lowerJ := strings.ToLower(j) // Все слова нужно привести к нижнему регистру.
		runeJ := []rune(lowerJ)
		slices.Sort(runeJ)
		sortedJ := string(runeJ) // в качестве ключа используем отсортированную строку
		resultMap[sortedJ] = append(resultMap[sortedJ], lowerJ)
		//добавляем приведенную к нижнему регистру строку в результирующий массив
	}
	for i, j := range resultMap {
		if len(j) < 2 { //удаляем те значения, у которых нет анаграмм
			delete(resultMap, i)
		} else {
			if i != j[0] {
				delete(resultMap, i) //заменяем ключ-сортированную строку на первое встреченное слово множества
				resultMap[j[0]] = j
				slices.Sort(resultMap[j[0]])
				//сортируем значение — срез из всех слов, принадлежащих этому множеству анаграмм
			}
		}
	}
	return resultMap
}
