package main

import (
	"fmt"

	"github.com/T117m/PaletteGen/internal/core"
	"github.com/T117m/PaletteGen/internal/utils"

	"flag"
	"log"
)

func main() {
	var (
		kFlag    = flag.Int("k", 5, "desired amount of colors in palette")
		algoFlag = flag.String("a", "dominant", "preferred generation algotithm. supported algorithms: dominant, mpa")
	)

	flag.Parse()

	var (
		k    = *kFlag
		algo = core.Dominant
	)

	if k <= 0 {
		log.Fatalf("Please, provide a valid amount of colors")
	}

	switch *algoFlag {
	case "dominant":
		algo = core.Dominant
	case "mpa":
		algo = core.MPA
		fmt.Println("MPA")
	default:
		log.Fatalf("Unknown algortihm %s", *algoFlag)
	}

	if len(flag.Args()) < 1 {
		log.Fatal("Provide at least one image path")
	}

	for _, filePath := range flag.Args() {
		palette := algo(filePath, k)

		fmt.Println()
		utils.PrintPalette(palette)

		fmt.Println()
		utils.DisplaImage(filePath)
	}
}
