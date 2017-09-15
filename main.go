package main

import (
	config "yellowroad_library/config"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes"
)
var configuration = config.Load("./config.json")
var container = containers.NewAppContainer(configuration)

func main() {
	routes.Init(container)
}

// func test() {
// 	router := gin.Default()
// 	//load middleware

// 	v1 := router.Group("/api/v1/")
// 	{
// 		// v1.GET("/greeting", greeting)
// 		// v1.Use(AuthMiddleware.AuthMiddlewareFactory(dbConn))
// 		v1.GET("/", func(c *gin.Context) {
// 			c.JSON(http.StatusOK, gin.H{"wtf ": "man"})
// 		})

// 		// v1.GET("/:id", FetchSingleTodo)
// 		// v1.PUT("/:id", UpdateTodo)
// 		// v1.DELETE("/:id", DeleteTodo)
// 	}

// 	router.Run()
// }

// func dummy(c *gin.Context) {
// 	var user models.User
// 	dbConn.Where("username = ?", "bob").First(&user)

// 	if token_claim, err := AuthMiddleware.GetTokenClaim(c); err == nil {
// 		token_claim.UserID
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"user":   user,
// 		"claims": token_claim,
// 	})
// }

// func greeting(c *gin.Context) {

// 	// _, err := userservice.RegisterUser("bob2", "eeeeeeee", "mail2@mail.com")

// 	user, token, err := userservice.LoginUser("bob2", "eeeeeeee")

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 	} else {
// 		val, _ := json.Marshal(user)
// 		c.JSON(http.StatusOK, gin.H{
// 			"user":  string(val),
// 			"token": token,
// 		})
// 	}

// }

// type Todo struct {
// 	gorm.Model
// 	Title     string `json:"title"`
// 	Completed int    `json:"completed"`
// }
// type TransformedTodo struct {
// 	ID        uint   `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }
// func CreateTodo(c *gin.Context) {
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))

// 	todo := Todo{
// 		Title:     c.PostForm("title"),
// 		Completed: completed,
// 	}

// 	db := Database()
// 	db.Save(&todo)
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status":     http.StatusCreated,
// 		"message":    "Todo item created successfully!",
// 		"resourceId": todo.ID,
// 	})
// }

// func FetchAllTodo(c *gin.Context) {
// 	var todos []Todo
// 	var _todos []TransformedTodo

// 	db := Database()
// 	db.Find(&todos)

// 	if len(todos) <= 0 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  http.StatusNotFound,
// 			"message": "No todo found!",
// 		})
// 		return
// 	}

// 	for _, item := range todos {
// 		completed := false
// 		if item.Completed == 1 {
// 			completed = true
// 		} else {
// 			completed = false
// 		}

// 		_todos = append(_todos, TransformedTodo{
// 			ID:        item.ID,
// 			Title:     item.Title,
// 			Completed: completed,
// 		})

// 	}

// }
