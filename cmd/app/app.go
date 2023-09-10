package app

import "fmt"

func Run(detachMode string) {
	if detachMode == "true" {
		fmt.Println("Running in detach mode ...")
	} else {
		fmt.Println("Starting! WaterFlow ...")
	}
}
