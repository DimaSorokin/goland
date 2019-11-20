package main

import "fmt"

type UserId int

func main() {
	idx := 1
	var uid UserId = 42
	// даже если базовый тип одинаковый, разные типы несовместимы
	// cannot use uid (type UserID) as type int64 in assignment
	//myId := idx

	myId := UserId(idx)

	fmt.Println(uid, myId)
}
