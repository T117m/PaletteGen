package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/T117m/PaletteGen/internal/core"
	"github.com/T117m/PaletteGen/internal/utils"
	"github.com/T117m/PaletteGen/views"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var (
		filePath = r.PostForm.Get("image")
		rK = r.PostForm.Get("colors")
		rAlgo = r.PostForm.Get("algorithm")
	)


	k, err := strconv.Atoi(rK)
	if err != nil {
		e := "Ошбка чтения количества цветов: %s" + err.Error()
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	algo := core.Dominant
	switch rAlgo {
	case "dominant":
		algo = core.Dominant
	case "mpa":
		algo = core.MPA
	case "mediancut":
		algo = core.MedianCut
	case "kmeans":
		algo = core.KMeans
	case "octree":
		algo = core.OcTree
	default:
		e := "Неизвестный алгоритм: %s" + rAlgo
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	img, err := utils.LoadImage(filePath)
	if err != nil {
		e := "Ошибка при загрузке изображения: " + err.Error()
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	p := algo(img, k)
	views.Result(filePath, p).Render(context.Background(), w)
}
