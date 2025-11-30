package services

import (
	"errors"
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Order("created_at desc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {

	var user models.User
	err := database.DB.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe usuario con id %d", id)
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	var existing models.User
	err := database.DB.Where("email = ?", user.Email).First(&existing).Error

	if user.RoleId != 0 {
		var role models.Role
		err := database.DB.First(&role, user.RoleId).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("no existe rol con id %d", user.RoleId)
			}
			return nil, err
		}
	}

	if err == nil {
		return nil, fmt.Errorf("ya existe un usuario con el email %s", user.Email)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(id uint, userData *models.User) (*models.User, error) {

	user, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if userData.RoleId != 0 {
		var role models.Role
		err := database.DB.First(&role, userData.RoleId).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("no existe rol con id %d", userData.RoleId)
			}
			return nil, err
		}
	}

	if userData.Email != "" && userData.Email != user.Email {
		var existingUser models.User
		err := database.DB.Where("email = ?", userData.Email).First(&existingUser).Error
		if err == nil && existingUser.ID != id {
			return nil, fmt.Errorf("ya existe un usuario con el email %s", userData.Email)
		}
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		user.Email = userData.Email
	}

	if userData.Name != "" {
		user.Name = userData.Name
	}
	if userData.RoleId != 0 {
		user.RoleId = userData.RoleId
	}

	if userData.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 14)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashed)
	}

	if err := database.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {

	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
