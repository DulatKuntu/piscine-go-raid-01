package main

import (
	"fmt"

	student ".."
)

func main() {
	student.Raid1c(5, 3)
	fmt.Println()
	student.Raid1c(5, 1)
	fmt.Println()
	student.Raid1c(1, 1)
	fmt.Println()
	student.Raid1c(1, 5)
}
