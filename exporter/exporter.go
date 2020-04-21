package exporter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"sme"
	"strings"
	"gopkg.in/yaml.v2"
)




//Export writes a new .txt file to the current working directory
func Export(platform sme.SocialMedia, filename string) error{
	if strings.HasSuffix(filename, "txt") {
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil{
			return err
		}
	
		for _, feed := range platform.Feed() {
		n, err := f.WriteString(feed + "\n")
		if err != nil{
			return err
		}
		
		fmt.Printf("Wrote %d bytes\n",n)
		}
		defer f.Close()

	}else if strings.HasSuffix(filename, "xml") {
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil{
			return err
		}

		xmldata, err := xml.MarshalIndent(platform, " ", "  ")
		if err != nil{
			return err
		}
		
		n, err := f.WriteString(string(xmldata) + "\n"+ "\n")
		if err != nil{
			return err
		}

		fmt.Printf("Wrote %d bytes\n", n)

		defer f.Close()

	}else if strings.HasSuffix(filename, "json"){

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	jsondata, err := json.Marshal(platform)
	if err != nil {
		return err
	}

	n, err := f.WriteString(string(jsondata) + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes\n", n)
	defer f.Close()

	}else if strings.HasSuffix(filename, "yaml"){
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	yamldata, err := yaml.Marshal(platform)
	if err != nil {
		return err
	}

	n, err := f.WriteString(string(yamldata) + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes\n", n)
	defer f.Close()
	}else{
		log.Fatal("This file format is not supported at the moment")

	}

	return nil
}


