package facebook


//Facebook is a mock type containing basic Info to fufill this learning stage
type Facebook struct {
	Title string	`xml:"1" json:"1" yaml:"1"`
	Feeds string	`xml:"2" json:"2" yaml:"2"`
}

// Feed returns the latest Facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here is my cool new selfie",
		"What is code?",
	}
}




