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

	tweets := new(twitter.Twitter)

 	linked := new(linkedin.LinkedIn)
	
 	facebook := new(facebook.Facebook)

	err := exporter.Exporttxt(facebook,"fb_data.txt")
	check(err)

	err = exporter.Exportjson(facebook, "fbdata.json")
	check(err)

	err = exporter.Exportxml(facebook,"fb_data.xml")
	check(err)

	err = exporter.Exportyaml(facebook,"fb_data.yaml")
	check(err)

	err = exporter.Exporttxt(tweets, "tweet_data.txt")
	check(err)

	err = exporter.Exportjson(tweets,"twitter_data.json")
	check(err)

	err = exporter.Exportxml(tweets, "twitter_data.xml")
	check(err)

	err = exporter.Exportyaml(tweets, "twitter_data.yaml")
	check(err)

	err = exporter.Exporttxt(linked, "LinkedIn_data.txt")
	check(err)

	err = exporter.Exportjson(linked,"LinkedIn_data.json")
	check(err)

	err = exporter.Exportxml(linked,"LinkedIn_data.xml")
	check(err)

	err = exporter.Exportyaml(linked,"LinkedIn_data.yaml")
	check(err)

}

