package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mahendrafathan/go-api-boilerplate/helpers"
	"github.com/mahendrafathan/go-api-boilerplate/models"
	"github.com/mahendrafathan/go-api-boilerplate/repository"

	"github.com/gorilla/mux"
)

type BookController struct {
	bookRepo repository.BookRepository
}

func NewBookController(repo repository.BookRepository) BookController {
	return BookController{
		bookRepo: repo,
	}
}

type BookResponseOutput struct {
	Data    interface{}
	Message string
}

func (b *BookController) AddBookItem(w http.ResponseWriter, r *http.Request) {
	Book := models.Book{}
	json.NewDecoder(r.Body).Decode(&Book)

	if len(Book.ISBN) < 3 {
		error.ApiError(w, http.StatusBadRequest, "ISBN should be at least 3 characters long!")
		return
	}

	if len(Book.Author) < 3 {
		error.ApiError(w, http.StatusBadRequest, "Author should be at least 3 characters long!")
		return
	}

	if result := b.bookRepo.Create(Book); result.Error != nil {
		error.ApiError(w, http.StatusInternalServerError, "Failed To Add new Book in database! \n"+result.Error.Error())
		return
	}

	helpers.RespondWithJSON(w, BookResponseOutput{
		Data:    Book,
		Message: "success",
	})
}

func (b *BookController) GetAllBookItems(w http.ResponseWriter, r *http.Request) {
	books := []models.Book{}
	if result := b.bookRepo.FindAll(&books); result.Error != nil {
		error.ApiError(w, http.StatusInternalServerError, "Failed To fetch Book in database! \n"+result.Error.Error())
		return
	}

	helpers.RespondWithJSON(w, BookResponseOutput{
		Data:    books,
		Message: "success",
	})
}

func (b *BookController) GetSingleBookItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book := models.Book{}
	result := b.bookRepo.FindOne(vars["isbn"], &book)
	if result.Error != nil {
		error.ApiError(w, http.StatusInternalServerError, "Failed To find Book in database! \n"+result.Error.Error())
		return
	}

	helpers.RespondWithJSON(w, BookResponseOutput{
		Data:    book,
		Message: "success",
	})
}

func (b *BookController) DeleteSingleBookItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book := models.Book{}
	if result := b.bookRepo.Delete(vars["isbn"], &book); result.Error != nil {
		error.ApiError(w, http.StatusInternalServerError, "Failed To delete Book in database! \n"+result.Error.Error())
		return
	}

	helpers.RespondWithJSON(w, BookResponseOutput{
		Message: "success",
	})
}
