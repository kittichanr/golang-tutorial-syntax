package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

	date := stuff.Date{}
	err := date.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1997)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day : %d/%d/%d\n", date.Month(), date.Day(), date.Year())
}
