package responses

type ExpandResponse struct {
	RedirectUrl string `json:"redirect_url"`
}

func NewExpandResponse(url string) ExpandResponse {
	return ExpandResponse{
		RedirectUrl: url,
	}
}
