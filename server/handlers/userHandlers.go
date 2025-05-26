package handlers

import (
	"kotoshop/models"
	"kotoshop/postgres"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary      Удаляет пользователя
// @Description  Удаляет пользователя по его id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param user_id query int true "user id"
// @Param Authorization header string true "Токен"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/user/delete [post]
func DeleteUser(c *gin.Context) {
	userRole := c.GetString("role")

	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error":"Not enough permissions",
		})
		return 
	}

	userId, err := strconv.Atoi(c.Query("user_id"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error on parsing user id",
		})
		
		return 
	}

	result := postgres.DB.Delete(&models.User{}, userId)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "could not delete user",
        })
        return
    }

    // Если ни одна запись не была удалена (RowsAffected == 0)
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "user not found",
        })
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"message":"user deleted successfully",
	})
}

// UpdateRule godoc
// @Summary      Меняет роль
// @Description	 Меняет роль пользователя
// @Tags         User
// @Accept       json
// @Produce      json
// @Param user_id query int true "user id"
// @Param role path string true "user role"
// @Param Authorization header string true "Токен"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/user/update_role [post]
func UpdateUserRole(c *gin.Context) {
	userRole := c.GetString("role")

	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error":"not enouth permissions",
		})
		return 
	}

	userID, err := strconv.Atoi(c.Query("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"error on parsing user id",
		})
		log.Print("error on parsing user id: ", err.Error())
		return 
	}

	newRole := c.Query("role")

	if err := postgres.DB.Model(&models.User{}).Where("user_id = %s", userID).Update("role", newRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on updating user role",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"role updated successfully",
	})
}

