package users

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"required"`
	IsActive bool   `json:"is_active"`
}

type ChangePasswordRequest struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}