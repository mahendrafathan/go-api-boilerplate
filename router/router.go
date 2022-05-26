package router

import (
	"fmt"
	"net/http"

	"github.com/mahendrafathan/go-api-boilerplate/controllers"
	"github.com/mahendrafathan/go-api-boilerplate/repository"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	// middleware "api/src/middlewares"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", MiddlewareCheckAuth(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hallo")
	}))

	userRepo := repository.NewUserRepoImpl(db)
	UserController := controllers.NewUserController(userRepo)

	router.HandleFunc("/auth/login", UserController.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", UserController.SignupUser).Methods(http.MethodPost)

	bookRepo := repository.NewBookRepoImpl(db)
	BookController := controllers.NewBookController(bookRepo)

	router.HandleFunc("/book", MiddlewareCheckAuth(BookController.AddBookItem)).Methods(http.MethodPost)
	router.HandleFunc("/book/all", MiddlewareCheckAuth(BookController.GetAllBookItems)).Methods(http.MethodGet)
	router.HandleFunc("/book/{isbn}", MiddlewareCheckAuth(BookController.GetSingleBookItem)).Methods(http.MethodGet)
	router.HandleFunc("/book/{isbn}", MiddlewareCheckAuth(BookController.DeleteSingleBookItem)).Methods(http.MethodDelete)

	return router
}
