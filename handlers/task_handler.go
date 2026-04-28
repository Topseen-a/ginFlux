package handlers

import (
	"goGin/models"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	tasks  []models.Task
	nextID uint = 1
	mu     sync.Mutex
)

func GetTasks(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	var input models.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	newTask := models.Task{
		ID:          nextID,
		Title:       input.Title,
		Description: input.Description,
		Status:      models.StatusPending,
		CreatedAt:   time.Now(),
	}
	nextID++
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, gin.H{"data": newTask})
}

func GetSingleTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, t := range tasks {
		if t.ID == uint(id) {
			c.JSON(http.StatusOK, gin.H{"data": t})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input struct {
		Title       string        `json:"title"`
		Description string        `json:"description"`
		Status      models.Status `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == uint(id) {
			tasks[i].Title = input.Title
			tasks[i].Description = input.Description
			tasks[i].Status = input.Status

			c.JSON(http.StatusOK, gin.H{"data": tasks[i]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == uint(id) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}

func GetFilteredTasks(c *gin.Context) {
	statusQuery := c.Query("status")

	mu.Lock()
	defer mu.Unlock()

	if statusQuery == "" {
		c.JSON(http.StatusOK, gin.H{"data": tasks})
		return
	}

	filtered := []models.Task{}

	for _, t := range tasks {
		if string(t.Status) == statusQuery {
			filtered = append(filtered, t)
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": filtered})
}

func PatchTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input models.PatchTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == uint(id) {
			if input.Title != nil {
				tasks[i].Title = *input.Title
			}
			if input.Description != nil {
				tasks[i].Description = *input.Description
			}
			if input.Status != nil {
				tasks[i].Status = models.Status(*input.Status)
			}
			c.JSON(http.StatusOK, gin.H{"data": tasks[i]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}
