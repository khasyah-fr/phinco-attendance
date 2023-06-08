package entities

type OnboardingMessage struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
