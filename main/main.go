package main

import "glitch-gin/cmd"

func main() {
	cmd.Start()
	if sqlDB, err := cmd.DB.DB(); err == nil {
		sqlDB.Close()
	}
}
