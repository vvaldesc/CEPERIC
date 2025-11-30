package service

import (
	"errors"

	"github.com/ceperic/backend/internal/domain"
	"github.com/ceperic/backend/internal/repository"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(dto domain.CreateUserDTO) (*domain.User, error) {
	// Validar que el email no exista
	existingUser, err := s.repo.FindByEmail(dto.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	user := &domain.User{
		Email:       dto.Email,
		Name:        dto.Name,
		FirebaseUID: dto.FirebaseUID,
		Role:        dto.Role,
	}

	if user.Role == "" {
		user.Role = "user"
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(id uint, dto domain.UpdateUserDTO) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Actualizar campos
	if dto.Email != "" {
		user.Email = dto.Email
	}
	if dto.Name != "" {
		user.Name = dto.Name
	}
	if dto.Role != "" {
		user.Role = dto.Role
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
