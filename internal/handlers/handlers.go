package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/rizhyi/6-sprint-final/internal/service"
)

func HandlerMain(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func HandlerUpload(res http.ResponseWriter, req *http.Request) {
	// HTML parsing
	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(res, "Не удалось спарсить форму", http.StatusInternalServerError)
		return
	}

	// Getting file
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "Файл не найден", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Reading data from file
	buf := make([]byte, handler.Size)
	if _, err := file.Read(buf); err != nil {
		http.Error(res, "Не удалось прочитать файл", http.StatusInternalServerError)
		return
	}

	// Translate data
	convertedStr := service.Translate(string(buf))

	// Creating local file
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	ext := filepath.Ext(handler.Filename)
	outputFileName := fmt.Sprintf("output_%s%s", timestamp, ext)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		http.Error(res, "Не удалось создать файл", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Writing converted data to file
	if _, err := outputFile.Write([]byte(convertedStr)); err != nil {
		http.Error(res, "Не удалось записать данные в файл", http.StatusInternalServerError)
		return
	}

	// Result
	res.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(res, "Конвертированная строка:\n%s\n", convertedStr)
	fmt.Fprintf(res, "Результат сохранён в файле: %s\n", outputFileName)
}
