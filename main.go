package main

import (
	"fmt"
	"time"
)

func main() {
	var s int
	fmt.Println("Введите количество секунд...")
	fmt.Scanln(&s)
	for i := 0; i < s; i++ {
		si := s - i
		if si < 10 {
			fmt.Printf("00:0%d\r", si)
		} else {
			fmt.Printf("00:%d\r", si)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Время истекло!")
}
