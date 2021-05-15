package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [4]string{"Kimchi", "Dongchimi", "Mull Kimchi", "Kkakdduki"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is Sexy"
}
