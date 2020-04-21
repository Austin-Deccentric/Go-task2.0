package linkedin


//LinkedIn defines the social media platform LinkedIn
type LinkedIn struct {
	Title string `xml:"1" json:"1"`
	Feeds string `xml:"2" json:"2"`
}

//Linked is an instance of LinkedIn struct
var Linked *LinkedIn = &LinkedIn{
	Title: "LinkedIn Feeds",
	Feeds: "I just got a job at Google!",
}

// Feed returns the latest LinkedIn posts
func (l *LinkedIn) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}
