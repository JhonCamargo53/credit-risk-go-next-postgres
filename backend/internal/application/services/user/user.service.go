package user

import (
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo ports.UserRepository
	roleRepo ports.RoleRepository
}

func NewUserService(userRepo ports.UserRepository, roleRepo ports.RoleRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAllOrderedByCreatedDesc()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("no existe usuario con id %d", id)
	}
	return user, nil
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	// Validar rol
	if user.RoleId != 0 {
		role, err := s.roleRepo.FindByID(user.RoleId)
		if err != nil {
			return nil, err
		}
		if role == nil {
			return nil, fmt.Errorf("no existe rol con id %d", user.RoleId)
		}
	}

	// Validar email único
	existing, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("ya existe un usuario con el email %s", user.Email)
	}

	// Hashear password si existe
	if user.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashed)
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(id uint, userData *models.User, requesterId uint) (*models.User, error) {
	user, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if id == requesterId {
		return nil, fmt.Errorf("un usuario no puede actualizarse a sí mismo")
	}

	// Validar rol si viene
	if userData.RoleId != 0 {
		role, err := s.roleRepo.FindByID(userData.RoleId)
		if err != nil {
			return nil, err
		}
		if role == nil {
			return nil, fmt.Errorf("no existe rol con id %d", userData.RoleId)
		}
	}

	// Validar email único
	if userData.Email != "" && userData.Email != user.Email {
		existingUser, err := s.userRepo.FindByEmail(userData.Email)
		if err != nil {
			return nil, err
		}
		if existingUser != nil && existingUser.ID != id {
			return nil, fmt.Errorf("ya existe un usuario con el email %s", userData.Email)
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

	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	if _, err := s.GetUserByID(id); err != nil {
		return err
	}
	return s.userRepo.Delete(id)
}
