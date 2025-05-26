package handlers

import (
	"fmt"
	"kotoshop/models"
	"kotoshop/postgres"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Signup godoc
// @Summary      Регистрирует нового пользователя
// @Description  Регистрирует пользователя через почту и пароль
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param user body models.User true "Данные пользователя"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/auth/signup [post]
func Signup(c *gin.Context) {
	var user models.User 

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err, 
		})

		return
	}

	userPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	user.Password = string(userPassword)

	if err := postgres.DB.Create(&user).Error; err != nil {
    log.Print("Ошибка при создании пользователя:", err) 
		
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Ошибка при создании пользователя: %s", err),
		})
    return
}

	accessToken, accessErr := createAccessToken(user.ID, user.Role)

	if accessErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error on creating token",
		})

		return
	} 

	c.JSON(http.StatusAccepted, gin.H{
		"message": "user successfully signed up",
		"token": accessToken,
	})
}


// Login godoc
// @Summary      Аутентифицирует пользователя
// @Description  Аутентифицирует пользователя через почту и пароль
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param user body models.UserLoginReq true "Данные пользователя"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/auth/login [post]
func Login(c *gin.Context) {
	var user models.User 

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err, 
		})

		return
	}

	var foundUser models.User 

	if err := postgres.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials",})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials",})

		return
	}

	accessToken, accessErr := createAccessToken(foundUser.ID, foundUser.Role)

	if accessErr != nil  {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error on creating token",
		})

		return
	} 

	c.JSON(http.StatusAccepted, gin.H{
		"message": "user successfully signed up",
		"token": accessToken,
	})
}

// Profile godoc
// @Summary      Возвращает данные о пользователе
// @Description  Возвращает данные о пользователе при корректном токене
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param  Authorization header string true "Access Token"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/auth/profile [get]
func Profile(c *gin.Context) {
	userID := c.GetUint("userID")

	var user models.User 
	if err := postgres.DB.Select("email, first_name, last_name, phone_number").First(&user, userID).Error; err != nil {
		log.Printf("error on selecting database: %v", err )
		c.AbortWithStatus(http.StatusUnauthorized)

		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
		"first_name": user.FirstName,
		"last_name":user.LastName,
		"phone_number":user.PhoneNumber,
	})
}

func UpdateUser (c *gin.Context) {
	userID := c.GetUint("userID")

	var req struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.ShouldBindJSON(&req); err != nil{
		log.Printf("error on parsing request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"error on parsing request",
		})
		return 
	}

	var user models.User 

	if err := postgres.DB.First(&user, userID).Error; err != nil {
		log.Printf("error on getting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on getting user",
		})
		return
	}

	if err := postgres.DB.Model(&user).Updates(models.User{FirstName: req.FirstName, LastName: req.LastName, PhoneNumber: req.PhoneNumber}).Error; err != nil {
		log.Printf("error on updating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"error on updating user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"user updated successfully",
	})
}