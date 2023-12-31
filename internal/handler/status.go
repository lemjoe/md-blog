package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/lemjoe/md-blog/internal/models"
)

type StatusCode struct {
	Code        string
	Title       string
	Description string
}

func (h *Handler) SendCode(w http.ResponseWriter, r *http.Request, status StatusCode) {
	curUser := h.GetCurrentUser(w.Header().Get("userID"))

	lang := curUser.Settings["language"]
	translation := Localizer([]string{"homeButton"}, lang, h.bundle)

	intCode, err := strconv.Atoi(status.Code)
	if err != nil { // if there is an error
		log.Print("status conversion error: ", err) // log it
		fmt.Fprintln(w, "status conversion error: ", err)
		return
	}
	w.WriteHeader(intCode)

	t, err := template.ParseFiles("lib/templates/status.html") //parse the html file homepage.html
	if err != nil {                                            // if there is an error
		log.Print("template parsing error: ", err) // log it
		fmt.Fprintln(w, "template parsing error: ", err)
		return
	}
	StatusPageVars := models.PageVariables{ //store the date and time in a struct
		HomeButton:   translation["homeButton"],
		Title:        status.Code + " - " + status.Title,
		BodyLoudText: status.Title,
		BodyText:     status.Description,
		Theme:        curUser.Settings["theme"],
	}
	err = t.Execute(w, StatusPageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                    // if there is an error
		log.Print("template executing error: ", err) //log it
		fmt.Fprintln(w, "template executing error: ", err)
		return
	}
}

func (h *Handler) PageNotFound(w http.ResponseWriter, r *http.Request) {
	// Send 404 if not found
	h.SendCode(w, r, statusCodes[http.StatusNotFound])
}
