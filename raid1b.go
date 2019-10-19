package student

import (
	"github.com/01-edu/z01"
)

func Raid1b(x, y int) {

	for y1 := 1; y1 <= y; y1++ {
		for x1 := 1; x1 <= x; x1++ {
			if y1 == 1 || y1 == y || x1 == 1 || x1 == x {
				if (x1 == 1 && y1 == 1) || (x1 == x && y1 == y && y != 1 && x != 1) {
					z01.PrintRune('/')
				} else if (x1 == x && y1 == 1) || (x1 == 1 && y1 == y) {
					z01.PrintRune('\\')
				} else {
					z01.PrintRune('*')
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
