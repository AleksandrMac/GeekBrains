package lesson2

//Макалов Александр
/*
3. **Исполнитель «Калькулятор» преобразует целое число, записанное на экране. У
исполнителя две команды, каждой присвоен номер:
1. Прибавь 1.
2. Умножь на 2.
Первая команда увеличивает число на экране на 1, вторая увеличивает его в 2 раза. Сколько
существует программ, которые число 3 преобразуют в число 20:
а. С использованием массива.
b. *С использованием рекурсии.
*/

//CalcPerformer -
func CalcPerformer(start int8, end int8, root bool) (count int64) {
	if start == end {
		return
	}
	if root {
		count++
	}
	if start*2 <= end {
		count++
		count += CalcPerformer(start*2, end, false)
	}
	count += CalcPerformer(start+1, end, false)
	return
}
