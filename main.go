package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// uncomment to check whether all digits are correctly displayed on terminal

	// for line := range zero {
	// 	for dig := range digits {
	// 		fmt.Printf("%s %s", digits[dig][line], strings.Repeat(" ", 5))
	// 	}
	// 	fmt.Println()
	// }

	var (
		clock = [12]digit{}
		t     time.Time
		h     int
		m     int
		s     int
		n     int
		ms1   int
		ms2   int
		ms3   int
		on    bool = true
	)

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

		time.Sleep(time.Millisecond * 100)
	}
}
