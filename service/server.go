package service

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
	mx.HandleFunc("/add/{id1}/{id2}", addHandler).Methods("GET")
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	vars := mux.Vars(req)
	id1 := vars["id1"]
	id2 := vars["id2"]
	id_1, _ := strconv.Atoi(id1)
	id_2, _ := strconv.Atoi(id2)
	id := id_1 + id_2
	id_str := strconv.Itoa(id)
	formatter.JSON(w, http.StatusOK, struct{ Test string }{id1 + " + " + id2 + " = " + id_str})
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
