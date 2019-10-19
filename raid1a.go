package piscine

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	x1 := x
	y1 := y
	for y = y1; y > 0; y-- {
		for x = x1; x > 0; x-- {
			if x == x1 {
				if y == y1 {
					z01.PrintRune(rune('o'))
				} else if y == 1 {
					z01.PrintRune(rune('o'))
				} else {
					z01.PrintRune(rune('|'))
				}
			} else if x == 1 {
				if y == y1 {
					z01.PrintRune(rune('o'))
					z01.PrintRune(rune('\n'))
				} else if y == 1 {
					z01.PrintRune(rune('o'))
					z01.PrintRune(rune('\n'))
				} else {
					z01.PrintRune(rune('|'))
					z01.PrintRune(rune('\n'))
				}
			} else {
				if y == y1 {
					z01.PrintRune(rune('-'))
				} else if y == 1 {
					z01.PrintRune(rune('-'))
				} else {
					z01.PrintRune(rune(' '))
				}
			}
		}
	}
}
