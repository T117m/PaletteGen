package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/T117m/PaletteGen/views"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		e := "Не удалось разобрать форму: %s" + err.Error()
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	rK, rAlgo := r.PostFormValue("colors"), r.PostFormValue("algorithm")

	k, err := strconv.Atoi(rK)
	if err != nil {
		e := "Ошбка чтения количества цветов: %s" + err.Error()
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	if k <= 0 {
		http.Error(w, "Укажите нормальное количество цветов", http.StatusBadRequest)
		return
	}

	if rAlgo != "dominant" && rAlgo != "mpa" && rAlgo != "mediancut" && rAlgo != "kmeans" && rAlgo != "octree" {
		e := "Неизвестный алгоритм: %s" + rAlgo
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		e := "Поле 'image' не найдено или пустое: %s" + err.Error()
		http.Error(w, e, http.StatusBadRequest)
		return
	}
	defer file.Close()

	newFileName := fmt.Sprintf("uploads/%d%s", time.Now().UnixNano(), header.Filename)

	dst, err := os.Create(newFileName)
	if err != nil {
		e := "Ошибка создания файла: " + err.Error()
		http.Error(w, e, http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		e := "Ошибка сохранения файла" + err.Error()
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	views.FetchResult(newFileName, rK, rAlgo).Render(context.Background(), w)
}
