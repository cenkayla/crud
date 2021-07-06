package internal

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cenkayla/crud/model"
	"github.com/gorilla/mux"
)

func (a *App) getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var all []model.Student
	err := a.DB.Find(&all).Error
	if err != nil {
		SendErr(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(all)
}

func (a *App) createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := model.Student{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		SendErr(w, err, http.StatusBadRequest)
		return
	}
	err = a.DB.Create(&s).Error
	if err != nil {
		SendErr(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *App) updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := model.Student{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		SendErr(w, err, http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	c := json.Number(id)
	if s.ID != c {
		SendErr(w, errors.New("invalid ID"), http.StatusBadRequest)
		return
	}
	s.ID = c
	err = a.DB.Save(&s).Error
	SendErr(w, err, http.StatusInternalServerError)
}

func (a *App) deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := model.Student{}
	id := mux.Vars(r)["id"]
	c := json.Number(id)
	err := a.DB.First(&s, c).Error
	if err != nil {
		SendErr(w, err)
		return
	}
	err = a.DB.Delete(s).Error
	SendErr(w, err)
}
