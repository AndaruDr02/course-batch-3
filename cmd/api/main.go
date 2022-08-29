package main

import (
	"course/internal/database"
	"course/internal/exercise/usecase"
	"course/internal/middleware"
	userUc "course/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.NewDabataseConn()
	exerciseUcs := usecase.NewExerciseUsecase(db)
	userUcs := userUc.NewUserUsecase(db)
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"message": "hello world",
		})
	})
	// exercise
	r.GET("/exercises/:id", middleware.WithAuthentication(userUcs), exerciseUcs.GetExercise)
	r.GET("/exercises/:id/score", middleware.WithAuthentication(userUcs), exerciseUcs.CalculateScore)

	// create new exercise
	r.POST("/exercises", middleware.WithAuthentication(userUcs), exerciseUcs.CreateExercise)

	// create questions of the exercise
	r.POST("/exercises/:id/questions", middleware.WithAuthentication(userUcs), exerciseUcs.CreateQuestions)

	// create Answer the question of the exercises
	r.POST("/exercises/:id/questions/:questionId/answer", middleware.WithAuthentication(userUcs), exerciseUcs.CreateAnswer)

	// user
	r.POST("/register", userUcs.Register)
	r.POST("/login", userUcs.Login)
	r.Run(":8080")
}
