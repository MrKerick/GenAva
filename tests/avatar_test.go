package tests

import (
	"avatar-generator/pkg/avatar"
	"os"
	"testing"
)

func getConfig() avatar.Config {

	return avatar.Config{
		Width:          300,
		Height:         300,
		FontSize:       45,
		FontPath:       "../assets/font.ttf",
		ColorsPath:     "../configs/colors.json",
		OutputDir:      "../output",
		TruncateLength: 2,
	}
}

func TestGenerateAvatar(t *testing.T) {

	generator :=
		avatar.New(getConfig())

	result, err :=
		generator.Generate(
			"test@gmail.com",
			"test.png",
		)

	if err != nil {
		t.Fatal(err)
	}

	_, err =
		os.Stat(result.FilePath)

	if os.IsNotExist(err) {
		t.Fatal("file not created")
	}
}

func TestTruncate(t *testing.T) {

	generator :=
		avatar.New(getConfig())

	result, err :=
		generator.Generate(
			"example@gmail.com",
			"test2.png",
		)

	if err != nil {
		t.Fatal(err)
	}

	if result.Email != "ex..." {
		t.Fatal("truncate failed")
	}
}