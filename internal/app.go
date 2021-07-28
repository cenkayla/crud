package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/cenkayla/crud/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (a *App) Open() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error when loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})
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
