package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is ", id)

	movie, err := app.model.DB.Get(id)
	/*movie := models.Movie{
		ID:          id,
		Title:       "Some Movie",
		Description: "some desc",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
	}*/

	app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
