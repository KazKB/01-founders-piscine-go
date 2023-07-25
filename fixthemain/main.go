package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	var open bool

	if ptrDoor.state == "OPEN" {
		open = true
	} else {
		open = false
	}
	return open
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")

	var close bool

	if ptrDoor.state == "CLOSE" {
		close = true
	} else {
		close = false
	}
	return close
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
