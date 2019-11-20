package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var num0 int //default

	var num1 int = 1 //значение при инициализации

	var num2 = 10 //пропускаем тип

	fmt.Println(num0, num1, num2)

	//короткое обьевление / только для новых переменных
	num := 30
	num += 1
	fmt.Println(num)

	// int - платформозависимый тип, 32/64
	var i int = 10
	// автоматически выбранный int
	var autoInt = -10
	// int8, int16, int32, int64 var bigInt int64 = 1<<32 - 1
	// платформозависимый тип, 32/64
	var unsignedInt uint = 100500
	// uint8, unit16, uint32, unit64
	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(i, autoInt, unsignedInt, unsignedBigInt)

	// float32, float64
	var pi float32 = 3.141
	var e = 2.718
	goldenRatio := 1.618
	fmt.Println(pi, e, goldenRatio)

	// bool
	var b bool // false по-умолчанию
	var isOk bool = true
	var success = true
	cond := true
	fmt.Println(b, isOk, success, cond)

	var str string
	// со спец символами
	var hello string = "Привет\n\t"
	// без спец символов
	var world string = `Мир\n\t`
	fmt.Println(str, hello, world)

	//Одинарные кавычки в go используются для символов byte, который является по сути alias’ом для uint8. Либо же для rune, который представляет собой полноценный UTF-8 символ. Внутри это uint32.
	// одинарные кавычки для байт (uint8)
	//var rawBinary byte = `\x27`
	// rune (uint32) для UTF-8 символов
	//var someRune rune = `*`
	//fmt.Println(rawBinary, someRune)

	helloWorld := "Привет Мир"
	// конкатенация строк
	andGoodMorning := helloWorld + " и доброе утро!"
	fmt.Println(andGoodMorning)

	// получение длины строки
	byteLen := len(helloWorld)                    // 19 байт
	symbols := utf8.RuneCountInString(helloWorld) // 10 рун
	fmt.Println(byteLen, symbols)

	// конвертация в слайс байт и обратно
	var byteString = []byte(helloWorld)
	helloWorld = string(byteString)

	fmt.Println(byteString, helloWorld)
	//Константы
	const (
		hello_const = "Привет"
		e_const     = 2.718
	)
	//Константный инкремент
	const (
		zero  = iota
		_     // пустая переменная, пропуск iota
		three // = 3
	)
	fmt.Println("********************************")
	fmt.Println(zero, three)

	const (
		// нетипизированная константа
		year = 2017
		// типизированная константа
		yearTyped int = 2017
	)
	/**
	Нельзя прибавлять переменные с разными типами, int32 + int (panic)
	*/

}
