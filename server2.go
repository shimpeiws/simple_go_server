package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/shimpeiws/simple_go_server/models"
	"github.com/unrolled/render"
)

func main() {
	ren := render.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		user := models.NewUser("Shimpei Takamatsu", 29)
		ren.JSON(w, http.StatusOK, user)
	})
	mux.HandleFunc("/index", func(w http.ResponseWriter, req *http.Request) {
		users := models.AllUsers()
		ren.JSON(w, http.StatusOK, users)
	})
	mux.HandleFunc("/show", func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.FormValue("id"))
		if err != nil {
			log.Println(err)
		}
		log.Println("id = " + strconv.Itoa(id))
		user := models.GetUser(id)
		ren.JSON(w, http.StatusOK, user)
	})
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
