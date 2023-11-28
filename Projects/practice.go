package main

import (
	"fmt"
)

func main() {

	//дефолт структуры данных инт, строки и константы
	v1, v2 := 3, 5
	const EarthRadius = 6371000
	_ = EarthRadius
	fmt.Println(v1, v2)

	str := "some text"
	fmt.Println(str[2:4])
	//----------------------

	//--Массивы-------------
	var x1 [5]int
	x2 := [...]int{1, 2, 3, 4}
	x3 := [3]int{1, 3, 5}
	x4 := [2]int64{1, 2}
	fmt.Println(x1, x2, x3, x4)
	//-----------------------

	//--Слайс или срез--------
	var list []int
	list1 := []int{1, 2, 3, 4}
	list = make([]int, 0, 5) // с помощью make можно создать новый слайс и обьявить его длину(len = 0 )
	// и емкость(capacity aka cap = 5)
	list = append(list, 5)
	fmt.Println(list, list1)
	//-------------------------

	//--отображение или map, или же хэш таблица
	age := map[string]int{
		"Макс": 30, // все ключи имеют одинаковый тип в мапе
		"Инна": 25, // все значение имеют одинаковый тип в мапе
	}
	fmt.Println(age)
	ages := age["Макс"]
	age["Наталья"] = 31 // добавлен ключ Наталья и значение 31
	fmt.Println(ages)   // 30
	// -----------------------

	//--Структуры--
	type Point struct {
		X int
		Y int
	}
	p := Point{
		X: 5,
		Y: 10,
	}
	fmt.Println(p.X, p.Y)
	//-------------------
}
func BoolToString(b bool) int {
	if b {
		return 1
	}
	return 0
}
