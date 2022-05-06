package user

import "plant-api/business/user"

// Create response payload
type userResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Construct the response
func GetResponse(user user.User) *userResponse {
	return &userResponse{
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	}
}

// Construct the array response
func GetAllResponse(users []user.User) []userResponse {
	result := []userResponse{}
	for _, user := range users {
		response := GetResponse(user)
		result = append(result, *response)
	}
	return result
}
