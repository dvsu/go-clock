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
		clock = [8]digit{}
		t     time.Time
		h     int
		m     int
		s     int
		on    bool = true
	)

	for {
		// clear terminal screen and move cursor to top left corner
		fmt.Print("\033[0;0H")
		// get the latest time
		t = time.Now()
		h, m, s = t.Hour(), t.Minute(), t.Second()

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
			}
		}

		for line := range zero {
			for dig := range clock {
				fmt.Printf("%s %s", clock[dig][line], strings.Repeat(" ", 3))
			}
			fmt.Println()
		}

		on = !on
		time.Sleep(time.Second)
	}
}
