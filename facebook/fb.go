package facebook


//Facebook is a mock type containing basic Info to fufill this learning stage
type Facebook struct {
	Title string	`xml:"1" json:"1" yaml:"1"`
	Feeds string	`xml:"2" json:"2" yaml:"2"`
}

//Fb is an instance of the Facebook struct
var Fb  *Facebook  = &Facebook{
	Title: "Facebook feeds",
	Feeds: "Hey, here's my cool new selfie",
}


// Feed returns the latest Facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}




