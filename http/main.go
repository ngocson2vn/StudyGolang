package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 30 * time.Second}

func getJson(url string, target interface{}) error {
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return json.NewDecoder(resp.Body).Decode(target)
	}

	return fmt.Errorf("Status code: %d", resp.StatusCode)
}

type Response struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func post(apiURL string, reader *bytes.Reader) (*http.Response, error) {
	var httpClient = &http.Client{Timeout: 300 * time.Second}
	req, err := http.NewRequestWithContext(context.Background(), "POST", apiURL, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	// url := "https://dummyjson.com/products/1"
	// result := Response{}
	// err := getJson(url, &result)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	b, _ := json.Marshal(result)
	// 	fmt.Println(string(b))
	// }

	data := struct {
		Title       string `json:"title"`
		FolderToken string `json:"folder_token"`
	}{
		Title:       "Test",
		FolderToken: "Dummy",
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := bytes.NewReader(b)
	apiURL := "https://open.feishu.cn/open-apis/sheets/v3/spreadsheets"
	res, err := post(apiURL, reader)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Response: %v+", res)
}
