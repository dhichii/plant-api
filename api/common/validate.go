package common

func ValidateByRole(role, claimsRole string) bool {
	if claimsRole == "super" {
		return true
	}
	if role != claimsRole {
		return false
	}
	return true
}

func ValidateById(id int, claimsId uint, role string) bool {
	if role == "super" {
		return true
	}
	if int(claimsId) != id {
		return false
	}
	return true
}
