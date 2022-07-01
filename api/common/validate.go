package common

func ValidateByRole(role, claimsRole string) bool {
	if claimsRole == "super" || claimsRole == role {
		return true
	}
	return false
}

func ValidateById(id int, claimsId uint, claimsRole string) bool {
	if claimsRole == "super" || int(claimsId) == id {
		return true
	}
	return false
}
