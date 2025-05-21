package main

import (
	"fmt"
	"time"

	"ascii.juanmartincoder.com/particles"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	fire := particles.NewFire(10, 4)
	fire.Start()

	for {
		select {
		case <-ticker.C:
			fmt.Print("\033[H\033[2J")
			fire.Update()

			flame := fire.Show()
			for _, row := range flame {
				fmt.Printf("%s%s\n", Padding, row)
			}

			fmt.Println(AsciiPit)
		}
	}
}
