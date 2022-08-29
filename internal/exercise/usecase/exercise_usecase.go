package usecase

import (
	"course/internal/domain"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseUsecase struct {
	db *gorm.DB
}

func NewExerciseUsecase(db *gorm.DB) *ExerciseUsecase {
	return &ExerciseUsecase{
		db: db,
	}
}

func (exerUseCase ExerciseUsecase) CreateAnswer(c *gin.Context) {
	type Answers struct {
		ExerciseID int
		QuestionID int
		UserID     int
		Answer     string `json:"answer"`
		CreatedAt  string
		UpdatedAt  string
	}
	paramID := c.Param("id")
	exerciseID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	paramID2 := c.Param("questionId")
	questionID, err := strconv.Atoi(paramID2)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid question id",
		})
		return
	}

	var exercise domain.Exercise
	err = exerUseCase.db.Where("id = ?", exerciseID).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	var question domain.Question
	err = exerUseCase.db.Where("id = ?", questionID).Take(&question).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "Question not found",
		})
		return
	}

	var createAnswerRequest Answers
	createAnswerRequest.ExerciseID = exerciseID
	createAnswerRequest.QuestionID = questionID
	createAnswerRequest.UserID = 1
	createAnswerRequest.CreatedAt = time.Now().Format(time.RFC3339)
	createAnswerRequest.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := c.ShouldBind(&createAnswerRequest); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if createAnswerRequest.Answer == "" {
		c.JSON(400, map[string]string{
			"message": "answer required",
		})
		return
	}

	if err := exerUseCase.db.Create(createAnswerRequest).Error; err != nil {
		c.JSON(500, map[string]string{
			"message": "cannot create answer",
		})
		return
	}

	c.JSON(201, map[string]interface{}{
		"message": "success to create answer",
	})
}

func (exerUseCase ExerciseUsecase) CreateQuestions(c *gin.Context) {
	type Questions struct {
		ExerciseID    int
		Body          string
		OptionA       string `json:"option_a"`
		OptionB       string `json:"option_b"`
		OptionC       string `json:"option_c"`
		OptionD       string `json:"option_d"`
		CorrectAnswer string `json:"correct_answer"`
		Score         int
		CreatorID     int
		CreatedAt     string
		UpdatedAt     string
	}
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = exerUseCase.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	// var createQuestionRequest domain.Question
	var createQuestionRequest Questions
	createQuestionRequest.ExerciseID = id
	createQuestionRequest.Score = 10
	createQuestionRequest.CreatorID = 1
	createQuestionRequest.CreatedAt = time.Now().Format(time.RFC3339)
	createQuestionRequest.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := c.ShouldBind(&createQuestionRequest); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if createQuestionRequest.Body == "" {
		c.JSON(400, map[string]string{
			"message": "body required",
		})
		return
	}

	if createQuestionRequest.OptionA == "" {
		c.JSON(400, map[string]string{
			"message": "option_a required",
		})
		return
	}

	if createQuestionRequest.OptionB == "" {
		c.JSON(400, map[string]string{
			"message": "option_b required",
		})
		return
	}

	if createQuestionRequest.OptionC == "" {
		c.JSON(400, map[string]string{
			"message": "option_c required",
		})
		return
	}

	if createQuestionRequest.OptionD == "" {
		c.JSON(400, map[string]string{
			"message": "option_d required",
		})
		return
	}

	if createQuestionRequest.CorrectAnswer == "" {
		c.JSON(400, map[string]string{
			"message": "correct_answer required",
		})
		return
	}

	if err := exerUseCase.db.Create(createQuestionRequest).Error; err != nil {
		c.JSON(500, map[string]string{
			"message": "cannot create questions",
		})
		return
	}

	c.JSON(201, map[string]interface{}{
		"message": "success to create questions",
	})
}

// create new exercise
func (exerUseCase ExerciseUsecase) CreateExercise(c *gin.Context) {
	type CreateExerciseRequest struct {
		Title       string
		Description string
	}

	var createRequest CreateExerciseRequest
	if err := c.ShouldBind(&createRequest); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if createRequest.Title == "" {
		c.JSON(400, map[string]string{
			"message": "title required",
		})
		return
	}

	if createRequest.Description == "" {
		c.JSON(400, map[string]string{
			"message": "description required",
		})
		return
	}

	exercise := domain.NewExercise(createRequest.Title, createRequest.Description)
	if err := exerUseCase.db.Create(exercise).Error; err != nil {
		c.JSON(500, map[string]string{
			"message": "cannot create exercise",
		})
		return
	}

	c.JSON(201, map[string]interface{}{
		"id":          exercise.ID,
		"title":       exercise.Title,
		"description": exercise.Description,
	})
}

func (exerUsecase ExerciseUsecase) GetExercise(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = exerUsecase.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "exercise not found",
		})
		return
	}
	c.JSON(200, exercise)
}

func (exerUsecase ExerciseUsecase) CalculateScore(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = exerUsecase.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	userID := int(c.Request.Context().Value("user_id").(float64))
	var answers []domain.Answer
	err = exerUsecase.db.Where("exercise_id = ? AND user_id = ?", id, userID).Find(&answers).Error
	if err != nil || len(answers) == 0 {
		c.JSON(200, map[string]interface{}{
			"score": 0,
		})
		return
	}
	mapQA := make(map[int]domain.Answer)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer
	}

	var score ScoreCount
	wg := new(sync.WaitGroup)
	for _, question := range exercise.Questions {
		newQuestion := question
		wg.Add(1)
		go func() {
			defer wg.Done()
			if strings.EqualFold(newQuestion.CorrectAnswer, mapQA[newQuestion.ID].Answer) {
				score.Inc(newQuestion.Score)
			}
		}()
	}
	wg.Wait()
	c.JSON(200, map[string]interface{}{
		"score": score.score,
	})
}

type ScoreCount struct {
	score int
	mu    sync.Mutex
}

func (sc *ScoreCount) Inc(value int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.score += value
}
