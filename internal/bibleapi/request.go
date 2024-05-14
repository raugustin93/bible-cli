package bibleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetVerse(book string, chapter string, verse string) (Verse, error) {
	fullUrl := versesURl
	fullUrl = strings.Replace(fullUrl, "${version}", "en-asv", 1)
	fullUrl = strings.Replace(fullUrl, "${book}", book, 1)
	fullUrl = strings.Replace(fullUrl, "${chapter}", chapter, 1)
	fullUrl = strings.Replace(fullUrl, "${verse}", verse, 1)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Verse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Verse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Verse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Verse{}, err
	}
	localVerse := Verse{}

	err = json.Unmarshal(data, &localVerse)
	if err != nil {
		return Verse{}, err
	}

	fmt.Printf("Verse: %v %v:%v \n", book, chapter, verse)
	fmt.Println(localVerse.Text)

	return localVerse, nil
}

func (c *Client) GetChapter(book string, chapter string) ([]Verse, error) {
	fullUrl := baseURl
	fullUrl = strings.Replace(fullUrl, "${version}", "en-asv", 1)
	fullUrl = strings.Replace(fullUrl, "${book}", book, 1)
	fullUrl = strings.Replace(fullUrl, "${chapter}", chapter, 1)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return []Verse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []Verse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return []Verse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Verse{}, err
	}
	localVerse := []Verse{}

	err = json.Unmarshal(data, &localVerse)
	if err != nil {
		return []Verse{}, err
	}

	fmt.Printf("%v chapter %v", book, chapter)
	for i, verse := range localVerse {
		fmt.Printf("%v - %v \n", i+1, verse)
	}

	return localVerse, nil
}
