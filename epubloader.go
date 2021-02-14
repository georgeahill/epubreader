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

type XMLPackage struct {
	XMLName xml.Name `xml:"package"`
	CurrentMetadata MetaData `xml:"metadata"`
	CurrentManifest Manifest `xml:"mainfest"`
}

type MetaData struct{
	XMLName xml.Name `xml:"metadata"`
	// what information do we want from the meta data
}

type Manifest struct{
	XMLName	xml.Name `xml:"manifest"`
	// ask george about if this is the correct way to make a slice
	Chapters []Item `xml:"item"`
}

type Item struct{
	XMLName xml.Name `xml: "item"`
	Href string `xml:"href,attr"`
	Id string `xml:"id,attr"`
	MediaType string `xml:"media-type,attr"`
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

func getDataFromPath(path string) (XMLPackage,error){
	// sends in the path from getContentPath
	xmlFile, err := os.Open(path)
	if err != nil{
		log.Fatal(err)
		return "",err
	}
	defer xmlFile.Close()

	text, _ := ioutil.ReadAll(xmlFile)
	var xmlPackage XMLPackage
	xml.Unmarshal(text, &xmlPackage)

	return xmlPackage,nil

}

//Is this returning the list of data? 
func getChapters(contentData XMLPackage) ([]Item,error){

	return xmlPackage.Manifest.Item, nil
}