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

	err := exporter.Exporttxt(facebook.Fb,"fb_data.txt")
	check(err)

	err = exporter.Exportjson(facebook.Fb, "fbdata.json")
	check(err)

	err = exporter.Exportxml(facebook.Fb,"fb_data.xml")
	check(err)

	err = exporter.Exportyaml(facebook.Fb,"fb_data.yaml")
	check(err)

	err = exporter.Exporttxt(twitter.Tweet, "tweet_data.txt")
	check(err)

	err = exporter.Exportjson(twitter.Tweet,"twitter_data.json")
	check(err)

	err = exporter.Exportxml(twitter.Tweet, "twitter_data.xml")
	check(err)

	err = exporter.Exportyaml(twitter.Tweet, "twitter_data.yaml")
	check(err)

	err = exporter.Exporttxt(linkedin.Linked, "LinkedIn_data.txt")
	check(err)

	err = exporter.Exportjson(linkedin.Linked,"LinkedIn_data.json")
	check(err)

	err = exporter.Exportxml(linkedin.Linked,"LinkedIn_data.xml")
	check(err)

	err = exporter.Exportyaml(linkedin.Linked,"LinkedIn_data.yaml")
	check(err)

}

