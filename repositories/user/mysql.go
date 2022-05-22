package user

import (
	"plant-api/api/v1/user/response"
	"plant-api/business/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// Generate mysql DB repository
func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Type used to update user
type userModel user.User

// Create new user and store into database
func (repo *repository) Create(user user.User) error {
	// Hash the user password
	user.Password = hashAndSalt([]byte(user.Password))
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Get all users
func (repo *repository) GetAll() ([]response.User, error) {
	users := []response.User{}
	if err := repo.db.Where("deleted_at IS NULL").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get user by given id. It's return nil if not found
func (repo *repository) Get(id int) (*response.User, error) {
	user := response.User{}
	if err := repo.db.Where("deleted_at IS NULL").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Get user by given email. It's return nil if not found
func (repo *repository) GetByEmail(email string) (*user.User, error) {
	user := user.User{}
	if err := repo.db.Where("email", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

/*
Update user and store it into database
will hash the user password if it's not empty
*/
func (repo *repository) Update(id int, user user.User) error {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		temp := userModel{}
		repo.db.Find(&temp, id)
		user.Password = temp.Password
	}
	if err := repo.db.Model(&user).
		Where("id", id).
		Updates(
			userModel{
				Name:     user.Name,
				Email:    user.Email,
				Password: user.Password,
			}).
		Error; err != nil {
		return err
	}
	return nil
}

// Hash password and return string of hashed password
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("failed to hash password")
	}
	return string(hash)
}
