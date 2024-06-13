package response

type SignInResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}
