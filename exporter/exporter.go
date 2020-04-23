package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"sme"
	"gopkg.in/yaml.v2"
)




//Exporttxt writes a new .txt file to the current working directory
func Exporttxt(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}
	defer f.Close()
	
	for _, feed := range platform.Feed() {
		n, err := f.WriteString(feed + "\n")
		if err != nil {
		return errors.New("An error occured writing to file" + err.Error())
		}
		
		fmt.Printf("Wrote %d bytes\n",n)
	}
	
	return nil
}


//Exportjson implements writing a json data to file
func Exportjson(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}
	defer f.Close()

	feeds:= platform.Feed()
	feedsmap:= make(map[int]interface{})
	for n, feed := range feeds{
		n++
		feedsmap[n] = feed
	}

	jsondata, err := json.Marshal(feedsmap)
	if err != nil {
		return errors.New("an error occured during encoding: " + err.Error())
	}

	n, err := f.WriteString(string(jsondata) + "\n")
	if err != nil {
		return errors.New("An error occured writing to file" + err.Error())
	}
	fmt.Printf("Wrote %d bytes\n", n)
	
	return nil
}


//Exportxml writes Xml data to file
func Exportxml(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}
	defer f.Close()

	feeds := platform.Feed()
	xmldata, err := xml.MarshalIndent(feeds, " ", "  ")
	if err != nil {
		return errors.New("an error occured during encoding: " + err.Error())
	}

	n, err := f.WriteString(string(xmldata) + "\n"+ "\n")
	if err != nil {
		return errors.New("An error occured writing to file" + err.Error())
	}

	fmt.Printf("Wrote %d bytes\n", n)
	return nil
}
	//Exportyaml exports the file to a yaml format
	func Exportyaml(platform sme.SocialMedia, filename string) error{
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}
	defer f.Close()

	feeds :=platform.Feed()

	yamldata, err := yaml.Marshal(feeds)
	if err != nil {
		return errors.New("an error occured during encoding: " + err.Error())
	}

	n, err := f.WriteString(string(yamldata) + "\n")
	if err != nil {
		return errors.New("An error occured writing to file" + err.Error())
	}
	fmt.Printf("Wrote %d bytes\n", n)
	
	return nil
}


