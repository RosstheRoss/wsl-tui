package main

import (
	"io"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/huh"
	"github.com/locusts-r-us/locusts"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	HostIP  string   `toml:"host_ip"`
	WslArgs []string `toml:"extra_wsl_args"`
	SshArgs []string `toml:"extra_ssh_args"`
}

type Distribution struct {
	Name     string `toml:"name"`
	Username string `toml:"username"`
	Port     int    `toml:"port"`
}

type ConfigFile struct {
	Config
	Distributions []Distribution `toml:"distribution"`
}

func main() {
	locusts.IntroduceLocusts()

	// Create config file location
	path, err := xdg.ConfigFile("wsl-tui/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	configFile, err := os.Open(path)
	if err != nil {
		// Open local config file
		configFile, err = os.Open("config.toml")
		if err != nil {
			log.Fatal(err)
		}
	}

	configContents, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}
	configFile.Close()

	config := ConfigFile{}
	err = toml.Unmarshal(configContents, &config)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(config)

	var distro Distribution

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[Distribution]().
				Title("Select a distribution to launch!").
				Value(&distro).
				OptionsFunc(func() []huh.Option[Distribution] {
					options := make([]huh.Option[Distribution], len(config.Distributions))
					for i, distro := range config.Distributions {
						options[i] = huh.Option[Distribution]{
							Key:   distro.Name,
							Value: distro,
						}
					}
					return options
				}, &config,
				),
		),
	).WithTheme(huh.ThemeDracula()).Run()

	if err != nil {
		if err == huh.ErrUserAborted {
			os.Exit(130)
		}
		log.Fatal(err)
	}

	cmd, err := command(distro, config.Config)
	if err != nil {
		log.Fatal(err)
	}
	// Pass in the TTY shite
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(cmd.Args[0], " exited with error: ", err)
	}
}
