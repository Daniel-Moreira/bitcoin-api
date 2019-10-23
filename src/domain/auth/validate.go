package auth

func Validate () string {
	if len(user.UserID) < 4 {
		return "UserID must have at least 4 characters"
	}
	if len(user.Password) < 4 {
		return "Password must have at least 4 characters"
  }
  
  return nil
}