package main

import (
	"errors"
	"fmt"
)

func main() {

	res, _ := Soma(7, 5)

	fmt.Println(res)

}

func Soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("total maior que 10")
	}

	return res, nil
}
