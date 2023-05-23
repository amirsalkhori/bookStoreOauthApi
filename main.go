package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Unix(2, 2))
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().UTC().Add(24 * time.Hour))
	fmt.Println(time.Now().UTC().Add(24 * time.Hour).Unix())

}
