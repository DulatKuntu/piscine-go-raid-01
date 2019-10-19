package main

import (
	"fmt"

	student ".."
)

func main() {
	student.Raid1b(5, 3)
	fmt.Println()
	student.Raid1b(5, 1)
	fmt.Println()
	student.Raid1b(1, 1)
	fmt.Println()
	student.Raid1b(1, 5)
}
