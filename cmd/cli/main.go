package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/T117m/PaletteGen/internal/core"
	"github.com/T117m/PaletteGen/internal/utils"
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
	case "dominant", "d":
		algo = core.Dominant
	case "mpa":
		algo = core.MPA
	case "median-cut", "mc":
		algo = core.MedianCut
	case "k-means", "km":
		algo = core.KMeans
	default:
		log.Fatalf("Unknown algortihm %s", *algoFlag)
	}

	if len(flag.Args()) < 1 {
		log.Fatal("Provide at least one image path")
	}

	for _, filePath := range flag.Args() {
		img, err := utils.LoadImage(filePath)
		if err != nil {
			log.Printf("Error loading image: %s", err)
		}

		palette := algo(img, k)

		utils.PrintPalette(palette)
		fmt.Println()

		utils.DisplaImage(filePath)
		fmt.Println()
	}
}
