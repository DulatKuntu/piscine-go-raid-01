package student

import (
	"github.com/01-edu/z01"
)

func Raid1b(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	
	for y1 := 1; y1 <= y; y1++ {
		for x1 := 1; x1 <= x; x1++ {
			
			var isTop = y1 == 1
			var isBottom = y1 == y
			var isLeft = x1 == 1
			var isRight = x1 == x
			
			if (isLeft && isTop) || (isRight && isBottom && y != 1 && x != 1) {
				z01.PrintRune('/')
				continue
			} else if (isRight && isTop) || (isBottom && isLeft) {
				z01.PrintRune('\\')
				continue
			} else {
				z01.PrintRune('*')
				continue
			} 

			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}
