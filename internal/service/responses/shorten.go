package responses

type ShortenResponse struct {
	ShortUrl string `json:"short_url"`
}

func NewShortenResponse(shortUrl string) ShortenResponse {
	return ShortenResponse{
		ShortUrl: shortUrl,
	}
}
