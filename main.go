package main

import (
	"fmt"
	"rivercrossing/event"
)

func main() {
	fmt.Println(event.ViewEvent())
	event.SettInnKylling()
	fmt.Println(event.ViewEvent())
}
