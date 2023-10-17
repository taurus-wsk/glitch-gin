package handler

import (
	"glitch-gin/cmd"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cmd.Start()
}

func main() {}
