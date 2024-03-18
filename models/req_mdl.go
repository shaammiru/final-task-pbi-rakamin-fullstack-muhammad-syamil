package models

type UserRegister struct {
	Username string `json:"username" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdate struct {
	Username string `json:"username" validate:"omitempty,alphanum"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PhotoCreate struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
	UserID   uint   `json:"user_id"`
}

type PhotoUpdate struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"omitempty,url"`
	UserID   uint   `json:"user_id"`
}
