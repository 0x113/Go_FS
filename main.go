package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type FilesResponse struct {
	Directory   string   `json:"directory"`
	Files       []string `json:"files"`
	Directories []string `json:"directories"`
	Error       string   `json:"error"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var root = os.Getenv("FILES_DIR")

func main() {
	port := os.Getenv("PORT")
	if root == "" {
		log.Fatal("FILES_DIR environment variable cannot be empty")
	}
	if port == "" {
		log.Fatal("PORT environment varaible cannot be empty")
	}
	// logger format
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// endpoints
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/file/serve/{path:.*}", serveFile).Methods("GET")
	r.HandleFunc("/{path:.*}", index).Methods("GET", "HEAD")
	http.Handle("/-/assets/", http.StripPrefix("/-/assets/", http.FileServer(http.Dir("./frontend/assets"))))
	http.Handle("/", r)

	log.Infof("Serving http on port: %s", port)
	http.ListenAndServe(":"+port, logRequest(accessControl(http.DefaultServeMux)))
}

func index(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	relPath := filepath.Join(root, path)

	t, err := template.New("index.html").Delims("[[", "]]").ParseFiles("index.html")
	if err != nil {
		errRes := errorResponse(relPath, err)
		t.Execute(w, errRes)
		return
	}
	files, dirs, err := scanDir(relPath)
	if err != nil {
		errRes := errorResponse(relPath, err)
		t.Execute(w, errRes)
		return
	}

	res := &FilesResponse{
		Directory:   relPath,
		Files:       files,
		Directories: dirs,
		Error:       "",
	}
	t.Execute(w, res)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	relPath := filepath.Join(root, path)
	// template
	t, err := template.New("index.html").Delims("[[", "]]").ParseFiles("index.html")
	if err != nil {
		errRes := errorResponse(relPath, err)
		t.Execute(w, errRes)
		return
	}
	// check if file exists
	if _, err := os.Stat(relPath); os.IsNotExist(err) {
		errRes := errorResponse(relPath, err)
		t.Execute(w, errRes)
		return
	}

	// if mp4 file
	if strings.HasSuffix(relPath, ".mp4") {
		w.Header().Add("Content-Type", "video/mp4")
	}

	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, relPath)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func errorResponse(relPath string, err error) *FilesResponse {
	return &FilesResponse{relPath, []string{}, []string{}, err.Error()}
}
