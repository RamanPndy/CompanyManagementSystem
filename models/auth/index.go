package auth

type CreateRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password"`
	Description string `json:"description"`
	IsActive    *bool  `json:"isActive"`
}

type Response struct {
	Token   *string `json:"token,omitempty"`
	Message *string `json:"message,omitempty"`
}
