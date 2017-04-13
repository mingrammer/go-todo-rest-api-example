package main

func main() {
	config := GetConfig()

	app := &App{}
	app.Initialize(config)
	app.Run(":3000")
}
