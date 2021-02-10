package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Container struct {
	XMLName xml.Name  `xml:"container"`
	Files   RootFiles `xml:"rootfiles"`
}

type RootFiles struct {
	XMLName     xml.Name `xml:"rootfiles"`
	CurrentFile RootFile `xml:"rootfile"`
}

type RootFile struct {
	XMLName xml.Name `xml:"rootfile"`
	// xml thingy different cause its an attribute not a tag
	FullPath string `xml:"full-path,attr"`
}

func Load() {

}

func getContentPath(path string) (string, error) {
	// path := "books/the-confessions-by-jean-jacques-rousseau"
	path += "/META-INF/container.xml"
	xmlFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer xmlFile.Close()

	text, _ := ioutil.ReadAll(xmlFile)
	// log.Printf("%q", text)

	var container Container
	xml.Unmarshal(text, &container)

	return container.Files.CurrentFile.FullPath, nil
}
