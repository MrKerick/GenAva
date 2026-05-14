package main

import (
	"avatar-generator/pkg/avatar"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	configFile, err :=
		os.ReadFile("./configs/config.json")

	if err != nil {
		panic(err)
	}

	var config avatar.Config

	err = json.Unmarshal(
		configFile,
		&config,
	)

	if err != nil {
		panic(err)
	}

	generator := avatar.New(config)

	result, err :=
		generator.Generate(
			"example@gmail.com",
			"avatar.png",
		)

	if err != nil {
		panic(err)
	}

	fmt.Println("File:", result.FilePath)
	fmt.Println("Color:", result.HexColor)
	fmt.Println("Email:", result.Email)
}