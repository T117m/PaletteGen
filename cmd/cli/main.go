package main

import (
	"github.com/T117m/PaletteGen/internal/core"
	"github.com/T117m/PaletteGen/internal/utils"

	"log"
	"os"
	"strconv"
)

func main() {
	var (
		filePath string
		k        = 7
	)

	switch len(os.Args) {
	case 1:
		log.Fatal("No arguments provided")
	case 2:
		filePath = os.Args[1]
	case 3:
		ks, err := strconv.Atoi(os.Args[1])
		if err != nil || k < 0 {
			log.Fatalf("Please, provide a valid amount of colors")
		}
		k = ks
		filePath = os.Args[2]
	default:
		log.Fatal("Unhandled amount of input")
	}

	palette := core.Dominant(filePath, k)

	utils.PrintPalette(palette)
	utils.DisplaImage(filePath)
}

