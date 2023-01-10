package main

import (
	"fmt"

	"github.com/batistondeoliveira/fcutil-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
