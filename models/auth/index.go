package auth

type Request struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

type Response struct {
	Token   *string `json:"token,omitempty"`
	Message *string `json:"message,omitempty"`
}
