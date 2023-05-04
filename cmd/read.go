package main

import (
	"fmt"
	"spells.tvblackman1.ru/lib/requests"
)

func main() {
	req := requests.NewRequest("asd")
	fmt.Println(req.String())
}
