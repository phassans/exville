package rocket

const (
	apiPath       = "api/v1"
	login         = "login"
	createChannel = "channels.create"
)

type (
	ErrorResponse struct {
		Status  string `json:"status"`
		Error   int    `json:"error"`
		Message string `json:"message"`
	}

	AdminCredentials struct {
		AuthToken string `json:"authToken"`
		UserId    string `json:"userId"`
	}

	Err struct {
		Success   bool   `json:"success"`
		Error     string `json:"error"`
		ErrorType string `json:"errorType"`
	}
)
