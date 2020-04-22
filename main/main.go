package main

import (
	"sme/exporter"
	"sme/facebook"
	"sme/linkedin"
	"sme/twitter"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	Tweets  := &twitter.Twitter{
		Title: "Twitter Timeline",
		Feeds: "Men are trash",
	}

 	Linked := &linkedin.LinkedIn{
	Title: "LinkedIn Feeds",
	Feeds: "I just got a job at Google!",
	}
	
 	Fb  := &facebook.Facebook{
	Title: "Facebook feeds",
	Feeds: "Hey, here is my cool new selfie",
}

	err := exporter.Exporttxt(Fb,"fb_data.txt")
	check(err)

	err = exporter.Exportjson(Fb, "fbdata.json")
	check(err)

	err = exporter.Exportxml(Fb,"fb_data.xml")
	check(err)

	err = exporter.Exportyaml(Fb,"fb_data.yaml")
	check(err)

	err = exporter.Exporttxt(Tweets, "tweet_data.txt")
	check(err)

	err = exporter.Exportjson(Tweets,"twitter_data.json")
	check(err)

	err = exporter.Exportxml(Tweets, "twitter_data.xml")
	check(err)

	err = exporter.Exportyaml(Tweets, "twitter_data.yaml")
	check(err)

	err = exporter.Exporttxt(Linked, "LinkedIn_data.txt")
	check(err)

	err = exporter.Exportjson(Linked,"LinkedIn_data.json")
	check(err)

	err = exporter.Exportxml(Linked,"LinkedIn_data.xml")
	check(err)

	err = exporter.Exportyaml(Linked,"LinkedIn_data.yaml")
	check(err)

}

