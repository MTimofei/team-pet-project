package web

import (
	"encoding/json"
	"fmt"
	jsonpkg "l/internal/json"
	"l/pkg/myerr"
	"log"
	"net/http"
	"os"
	"time"
)

func page() (h http.Handler) {
	fs := http.FileServer(http.Dir("./dist"))
	h = http.StripPrefix("/", fs)
	return h
}

func handlerCreateFile(w http.ResponseWriter, r *http.Request) {

	log.Println(*r.URL)

	if r.Method != http.MethodPost {
		err := fmt.Errorf("нетоы")
		myerr.ServesError(w, err)
	}
	var jsontempl jsonpkg.Files
	err := json.NewDecoder(r.Body).Decode(&jsontempl)
	if err != nil {
		myerr.ServesError(w, err)
	}
	log.Println(jsontempl.Title)
	file, err := os.Create(fmt.Sprintf("stor/%s.txt", jsontempl.Title))
	if err != nil {
		myerr.ServesError(w, err)
	}
	defer file.Close()
	_, err = file.WriteString(jsontempl.Description)
	if err != nil {
		myerr.ServesError(w, err)
	}
}

func handlerGetFile(w http.ResponseWriter, r *http.Request) {

	log.Println(*r.URL)

	if r.Method != http.MethodGet {
		err := fmt.Errorf("")
		myerr.ServesError(w, err)
	}
	var jsonrequst jsonpkg.Request
	err := json.NewDecoder(r.Body).Decode(&jsonrequst)
	if err != nil {
		myerr.ServesError(w, err)
	}

	file, err := os.Open(fmt.Sprintf("stor/%s.txt", jsonrequst.Title))
	if err != nil {
		myerr.ServesError(w, err)
	}
	defer file.Close()
	filename := fmt.Sprintf("%s.txt", jsonrequst.Title)
	//i, _ := file.Stat()
	//w.Header().Set("Content-Type", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	//w.Header().Set("Content-Length", fmt.Sprintf("%d", i.Size()))

	http.ServeContent(w, r, filename, time.Now(), file)

}
