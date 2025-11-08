package handlers

import (
	"net/http"

	"task-api/internal/database"
	"task-api/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Summary Create new task
// @Description Create task for logged in user
// @Tags Tasks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param task body models.Task true "Task data"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Router /api/tasks [post]
func CreateTask(c *gin.Context) {
	userID := c.GetUint("user_id") // ✅ Ambil user_id dari token

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
		UserID:      userID, // ✅ Pastikan diset
	}

	database.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	userID := c.GetUint("user_id")

	var tasks []models.Task
	database.DB.Where("user_id = ?", userID).Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}

// GetTasks godoc
// @Summary Get user tasks
// @Description Get all tasks of logged in user
// @Tags Tasks
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Task
// @Router /api/tasks [get]
func GetTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not yours"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update task by ID
// @Tags Tasks
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param task body models.Task true "Updated task"
// @Success 200 {object} models.Task
// @Failure 404 {object} map[string]string
// @Router /api/tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not yours"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task updated", "task": task})
}

// DeleteTask godoc
// @Summary Delete task
// @Description Remove task by ID
// @Tags Tasks
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not yours"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
