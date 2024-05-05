package vault

type AccountCreateRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type AccountCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Account AccountDetails `json:"account"`
}

type AccountDetails struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type AccountUpdateRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"` // Omitempty pour permettre la mise à jour partielle
	Description string `json:"description,omitempty"`
}

type AccountUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Account AccountDetails `json:"account"`
}

type AccountDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
