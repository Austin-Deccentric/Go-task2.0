package twitter


//Twitter defines the social media platform Twitter
type Twitter struct {
	Title string `xml:"1" json:"1" yaml:"1"`
	Feeds string `xml:"2" json:"2" yaml:"2"`
}

// Feed returns the latest Twitter posts
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Hey, here's my cool new selfie",
		"What! is you a bitch?",
	}
}
