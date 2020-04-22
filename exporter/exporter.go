package exporter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"sme"
	"gopkg.in/yaml.v2"
)




//Exporttxt writes a new .txt file to the current working directory
func Exporttxt(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	for _, feed := range platform.Feed() {
		n, err := f.WriteString(feed + "\n")
		if err != nil {
			return err
		}
		defer f.Close()
		fmt.Printf("Wrote %d bytes\n",n)
	}
	
	return nil
}


//Exportjson implements writing a json data to file
func Exportjson(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	jsondata, err := json.Marshal(platform)
	if err != nil {
		return err
	}

	n, err := f.WriteString(string(jsondata) + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes\n", n)
	
	return nil
}


//Exportxml writes Xml data to file
func Exportxml(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	xmldata, err := xml.MarshalIndent(platform, " ", "  ")
	if err != nil {
		return err
	}

	n, err := f.WriteString(string(xmldata) + "\n"+ "\n")
	if err != nil {
		return err
	}

	fmt.Printf("Wrote %d bytes\n", n)
	return nil
}
	//Exportyaml exports the file to a yaml format
	func Exportyaml(platform sme.SocialMedia, filename string) error{
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	yamldata, err := yaml.Marshal(platform)
	if err != nil {
		return err
	}

	n, err := f.WriteString(string(yamldata) + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes\n", n)
	
	return nil
}


