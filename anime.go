package main

import (
	"io"
	"net/http"
)

func random_anime_msg() string {
	response, err := http.Get("https://anime-gifs.aaron-lun.workers.dev/random")

	if err != nil {
		return (err.Error())
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return (err.Error())
	}
	return (string(responseData))
}

func main() {
	random_anime_msg()
}
