package main

import (
	"fmt"

	student ".."
)

func main() {
	student.Raid1e(5, 3)
	fmt.Println()
	student.Raid1e(5, 1)
	fmt.Println()
	student.Raid1e(1, 1)
	fmt.Println()
	student.Raid1e(1, 5)
	fmt.Println()
	student.Raid1e(-1, -6)
}
