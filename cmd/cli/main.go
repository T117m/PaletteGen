package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/T117m/PaletteGen/internal/core"
	"github.com/T117m/PaletteGen/internal/utils"
)

const (
	DescK          = "Количество цветов в палитре"
	DescAlgo       = "Алгоритм генерации палитры"
	AvailableAlgos = "[dominant/d | mpa | median-cut/mc | k-means/km | octree/ot ]"
)

func main() {
	var (
		kFlag = flag.Int("k", 5, DescK)

		algoFlag = flag.String("a", "dominant", fmt.Sprint(
			DescAlgo,
			" ",
			AvailableAlgos,
		))
	)

	flag.Parse()

	var (
		k    = *kFlag
		algo = core.Dominant
	)

	if k <= 0 {
		fmt.Printf("Пожалуйста, укажите нормальное количество цветов")
		return
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
	case "octree", "ot":
		algo = core.OcTree
	default:
		fmt.Printf(
			"Неизвестный алгоритм: %s\nДоступные алгоритмы: %s",
			*algoFlag,
			AvailableAlgos,
		)
		return
	}

	if len(flag.Args()) < 1 {
		log.Fatal("Предоставьте хотя бы одно изображение")
	}

	for _, filePath := range flag.Args() {
		img, err := utils.LoadImage(filePath)
		if err != nil {
			fmt.Printf("Ошибка при загрузке изображения: %s", err)
			continue
		}

		utils.DisplayImage(filePath)
		fmt.Println()

		utils.PrintPalette(algo(img, k))
		fmt.Println()
	}
}
