package controller

import (
	"net/http"

	"github.com/ESD-Laboratory/Software-Development-Backend/internal/app/dto"
	"github.com/ESD-Laboratory/Software-Development-Backend/internal/config"
	"github.com/ESD-Laboratory/Software-Development-Backend/internal/entity"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

//Register is a function to register user
func (u *UserHandler) Register(c *fiber.Ctx) error {
	//Body Request
	req := dto.RegisterRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Data Request!",
		})
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(encryptPassword),
	}

	//check if user already registered
	var existUser entity.User
	err = config.DB.Where("email = ?", user.Email).First(&existUser).Error

	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User already Exist!",
		})
	}

	err = config.DB.Create(&user).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Create User Fail!",
		})
	}

	response := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})
}
