package main

import "testing"

func TestLoad(t *testing.T) {
	if false {
		t.Error("Failed! Underscore Test")
	}
}

func TestGetContentPath_Success(t *testing.T) {
	expected := "OEBPS/content.opf"
	path := "books/the-confessions-by-jean-jacques-rousseau"
	output, err := getContentPath(path)
	if err != nil || output != expected {
		t.Errorf("%v != %v as expected. err: %v", output, expected, err)
	}
}
