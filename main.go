package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	if len(os.Args) < 2 {
		return ErrNoCmd
	}

	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	switch cmd := os.Args[1]; cmd {
	case "get":
		hue, err := getHue(cfg)
		if err != nil {
			return err
		}

		fmt.Println(hue)
	case "set":
		if len(os.Args) < 3 {
			return ErrIDNotSpecified
		}
		id := os.Args[2]

		scene, err := cfg.Scenes.get(id)
		if err != nil {
			return err
		}

		if err := putHueState(cfg, scene.LightID, scene.State); err != nil {
			return err
		}

		hue, err := getHue(cfg)
		if err != nil {
			return err
		}

		fmt.Println(hue)
	}

	return nil
}
