/*
	pever- A simple battery status monitor
*/

package main

import(
	"fmt"
	"os"
)

const(
	YellowColor = "\033[1;33m%s\033[0m"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){

	fmt.Printf(YellowColor, "\n Welcome to Pever. Here are your power statistics!\n")

	f, err := os.Open("/sys/class/power_supply/BAT0/capacity")
	check(err)

	b1 := make([]byte, 0x4)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("\n Current battery percentage: %s%% \n", string(b1[:n1-1]))

	f, err = os.Open("/sys/class/power_supply/BAT0/status")
	check(err)

	b1 = make([]byte, 0xc)
	n1, err = f.Read(b1)
	check(err)
	fmt.Printf(" Current battery status: %s", string(b1[:n1]))
}
