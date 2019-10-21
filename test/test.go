package main

import (
	"fmt"

	student ".."
)

func main() {
	student.Raid1e(5, 0)
	fmt.Println()
	student.Raid1e(0, 1)
	fmt.Println()
	student.Raid1e(0, 1)
	fmt.Println()
	student.Raid1e(6, 0)
	fmt.Println()
	student.Raid1e(0, -6)
}
