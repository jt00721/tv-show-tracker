package main

import "github.com/jt00721/tv-show-tracker/config"

func main() {
	application := config.NewApp()

	application.Run()
}
