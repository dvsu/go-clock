package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const usage = ".\\go-clock.exe -s [true|false]"

func main() {

	var (
		l          int = 12
		max        int = 12
		cursor     int
		clock      = [12]digit{}
		clockSlide = [12]digit{}
		t          time.Time
		h          int
		m          int
		s          int
		n          int
		ms1        int
		ms2        int
		ms3        int
		on         bool = true
		slide      bool = false
		forward    bool = false
	)

	args := os.Args[1:]

	if len(args) == 2 {
		if args[0] == "-s" {
			switch {
			case strings.ToLower(args[1]) == "true":
				slide = true
			case strings.ToLower(args[1]) == "false":
				slide = false
			default:
				fmt.Println(usage)
				return
			}
		} else {
			fmt.Println(usage)
			return
		}
	}

	// uncomment to check whether all digits are correctly displayed on terminal

	// for line := range zero {
	// 	for dig := range digits {
	// 		fmt.Printf("%s %s", digits[dig][line], strings.Repeat(" ", 5))
	// 	}
	// 	fmt.Println()
	// }

	for {
		// clear terminal screen and move cursor to top left corner
		fmt.Print("\033[0;0H")
		// get the latest time
		t = time.Now()
		h, m, s, n = t.Hour(), t.Minute(), t.Second(), t.Nanosecond()
		ms1 = n / 100000000
		ms2 = (n / 10000000) - (ms1 * 10)
		ms3 = (n / 1000000) - (ms1 * 100) - (ms2 * 10)

		switch on {
		case true:
			clock = [...]digit{
				digits[h/10],
				digits[h%10],
				separator,
				digits[m/10],
				digits[m%10],
				separator,
				digits[s/10],
				digits[s%10],
				dot,
				digits[ms1],
				digits[ms2],
				digits[ms3],
			}
		case false:
			clock = [...]digit{
				digits[h/10],
				digits[h%10],
				blank,
				digits[m/10],
				digits[m%10],
				blank,
				digits[s/10],
				digits[s%10],
				dot,
				digits[ms1],
				digits[ms2],
				digits[ms3],
			}
		}

		if slide {
			cursor = max - l
			if forward {
				for i := 0; i < cursor; i++ {
					clockSlide[i] = blank
				}
				for i := cursor; i < max; i++ {
					clockSlide[i] = clock[i-cursor]
				}
				l += 1
			} else {
				for i := cursor; i < max; i++ {
					clockSlide[i-cursor] = clock[i]
				}
				for i := l; i < max; i++ {
					clockSlide[i] = blank
				}
				l -= 1
			}

			switch l {
			case 0:
				forward = true
			case 12:
				forward = false
			}

			for line := range zero {
				for dig := range clockSlide {
					fmt.Printf("%s %s", clockSlide[dig][line], strings.Repeat(" ", 3))
				}
				fmt.Println()
			}

		} else {
			for line := range zero {
				for dig := range clock {
					fmt.Printf("%s %s", clock[dig][line], strings.Repeat(" ", 3))
				}
				fmt.Println()
			}

			switch {
			case s%2 == 0:
				on = false
			default:
				on = true
			}
		}

		time.Sleep(time.Millisecond * 100)
	}
}
