package handlers

import (
	"net/http"

	"task-api/internal/database"
	"task-api/internal/models"

	"github.com/gin-gonic/gin"
)

// CREATE TASK
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

// GET ALL TASKS OF LOGGED USER
func GetTasks(c *gin.Context) {
	userID := c.GetUint("user_id")

	var tasks []models.Task
	database.DB.Where("user_id = ?", userID).Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}

// GET TASK BY ID (must belong to user)
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

// UPDATE TASK (only if belongs to user)
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

// DELETE TASK (only if belongs to user)
func DeleteTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not yours"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
