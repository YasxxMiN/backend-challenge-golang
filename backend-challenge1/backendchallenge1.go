package backendchallenge1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func fetchJson(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func maxPathSum(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += max(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PrintTriangle() (result int, err error) {
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"
	data, err := fetchJson(url)
	if err != nil {
		return result, err
	}

	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		return result, err
	}

	result = maxPathSum(triangle)

	return result, err

}
