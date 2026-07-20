ifeq ($(OS),Windows_NT)
    CLI := palgen.exe
    WEB := palgenserver.exe
else
    CLI := palgen
    WEB := palgenserver
endif

cli: cmd/cli/main.go
	go build -o build/$(CLI) cmd/cli/main.go

web: cmd/web/main.go
	templ generate
	go build -o build/$(WEB) cmd/web/main.go
