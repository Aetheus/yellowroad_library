package book_crud

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/utils/gin_tools"
	"net/http"
)

func FetchSingleBook(bookRepo book_repo.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		book_id, err := gin_tools.GetIntParam("book_id", c)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		book, findErr := bookRepo.FindById(book_id)
		if findErr != nil {
			c.JSON(api_response.ConvertErrWithCode(findErr))
			return
		}

		c.JSON(api_response.SuccessWithCode(book))
		return
	}
}

func FetchBooks(repository book_repo.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := gin_tools.GetIntQueryOrDefault("page",1,c)
		perpage := gin_tools.GetIntQueryOrDefault("perpage",15,c)

		//TODO: actually get some search options
		results, err := repository.Paginate(page,perpage, book_repo.SearchOptions{})
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(api_response.SuccessWithCode(results))
		return
	}
}

func UpdateBook (book_repo book_repo.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		book, err := book_repo.FindById(book_id)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		bookForm := entities.BookForm{}
		bindErr := c.BindJSON(&bookForm)
		if bindErr != nil {
			c.JSON(api_response.ConvertErrWithCode(app_error.Wrap(bindErr)))
			return
		}

		bookForm.Apply(&book)
		err = book_repo.Update(&book)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{ "book" : book })
		return
	}
}