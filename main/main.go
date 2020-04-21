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

	err := exporter.Export(facebook.Fb,"fb_data.txt")
	check(err)

	err = exporter.Export(facebook.Fb, "fbdata.json")
	check(err)

	err = exporter.Export(facebook.Fb,"fb_data.xml")
	check(err)

	err = exporter.Export(facebook.Fb,"fb_data.yaml")
	check(err)

	err = exporter.Export(twitter.Tweet, "tweet_data.txt")
	check(err)

	err = exporter.Export(twitter.Tweet,"twitter_data.json")
	check(err)

	err = exporter.Export(twitter.Tweet, "twitter_data.xml")
	check(err)

	err = exporter.Export(twitter.Tweet, "twitter_data.yaml")
	check(err)

	err = exporter.Export(linkedin.Linked, "LinkedIn_data.txt")
	check(err)

	err = exporter.Export(linkedin.Linked,"LinkedIn_data.json")
	check(err)

	err = exporter.Export(linkedin.Linked,"LinkedIn_data.xml")
	check(err)

	err = exporter.Export(linkedin.Linked,"LinkedIn_data.yaml")
	check(err)

}

