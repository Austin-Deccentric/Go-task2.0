package linkedin


//LinkedIn defines the social media platform LinkedIn
type LinkedIn struct {
	Title string `xml:"1" json:"1" yaml:"1"`
	Feeds string `xml:"2" json:"2" yaml:"2"`
}

// Feed returns the latest LinkedIn posts
func (l *LinkedIn) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}
