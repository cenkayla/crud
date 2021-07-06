package internal

import (
	"net/http"

	"github.com/cenkayla/crud/model"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
}

func (a *App) Open() error {
	dsn := "postgres://destr0:12345m@localhost:5432/crud?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	a.DB = db
	a.SetRouters()
	return a.DB.AutoMigrate(&model.Student{})
}

func (a *App) SetRouters() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/students", a.getAllStudents).Methods("GET")
	r.HandleFunc("/students", a.createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", a.updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", a.deleteStudent).Methods("DELETE")
	return r
}
