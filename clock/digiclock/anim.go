package digiclock

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

//AnimateClock animates the digital clock according to the current time.
func AnimateClock() {
	for {
		//time storage
		currentTime := time.Now()
		hour := currentTime.Hour()
		minute := currentTime.Minute()
		second := currentTime.Second()

		//animate separators
		if second%2 == 0 {
			clock[2] = blank
			clock[5] = blank
		} else {
			clock[2] = separator
			clock[5] = separator
		}

		//animate clock
		clock[0] = digits[hour/10]
		clock[1] = digits[hour%10]
		clock[3] = digits[minute/10]
		clock[4] = digits[minute%10]
		clock[6] = digits[second/10]
		clock[7] = digits[second%10]

		//draw clock
		screen.MoveTopLeft()
		for i := 0; i < 5; i++ {
			fmt.Println()
			for j := 0; j < 8; j++ {
				fmt.Printf("%v\t", clock[j][i])
			}
		}

		time.Sleep(time.Second)
		screen.Clear()
	}
}
