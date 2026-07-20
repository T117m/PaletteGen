package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/T117m/PaletteGen/internal/handler"
	"github.com/T117m/PaletteGen/views"
	"github.com/a-h/templ"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var (
		index   = views.Index()
		static  = http.FileServer(http.Dir("./static"))
		uploads = http.FileServer(http.Dir("./uploads"))
	)

	http.Handle("/static/", http.StripPrefix("/static/", static))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", uploads))

	go func() {
		<-sigChan
		fmt.Println("Received Ctrl+C, closing...")

		os.Exit(0)
	}()

	http.Handle("/", templ.Handler(index))
	http.HandleFunc("POST /upload", handler.Upload)
	http.HandleFunc("POST /generate", handler.Generate)

	fmt.Println("Running on :7777")
	http.ListenAndServe(":7777", nil)
}
