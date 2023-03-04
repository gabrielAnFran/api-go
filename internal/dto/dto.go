package dto

type CreateShoesInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Brand string  `json:"brand"`
	Size  float64 `json:"size"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
