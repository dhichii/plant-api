package user

import (
	"plant-api/api/v1/user/response"
	"plant-api/business/user"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Create new user and store into database
func (repo *repository) Create(user user.User) (uint, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// Get all users
func (repo *repository) GetAll() ([]response.User, error) {
	users := []response.User{}
	if err := repo.db.Find(&users, "deleted_at IS NULL").Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get user by given id. It's return nil if not found
func (repo *repository) Get(id int) (*response.User, error) {
	user := response.User{}
	if err := repo.db.First(&user, id, "deleted_at IS NULL").Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Get user by given email. It's return nil if not found
func (repo *repository) GetByEmail(email string) (*user.User, error) {
	user := user.User{}
	if err := repo.db.First(&user, "email", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update user and store it into database
func (repo *repository) Update(id int, user user.User) error {
	if err := repo.db.Where("id", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
