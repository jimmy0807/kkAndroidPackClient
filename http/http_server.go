package http

import (
	"kkdownloader/http/handler"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//HandleHTTPServer 水电费
func HandleHTTPServer(rootPath string) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/doPackage", handler.HandlerDoPackage).Methods("GET")
	r.HandleFunc("/api/getPackageList", handler.HandlerGetPackageList).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wait till you see me in action!"))
	})

	http.Handle(rootPath, r)
	return r
}

//Start 初始化
func Start(url string, handler http.Handler) {
	headsOpts := handlers.AllowedHeaders([]string{"Content-Type"})
	http.ListenAndServe(url, handlers.CORS(headsOpts)(handler))
}
