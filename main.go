package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, host := range os.Args[1:] {
		resp, err := ntp.Query(host)
		if err == nil {
			fmt.Println(host, resp.ClockOffset)
		}
	}
}
