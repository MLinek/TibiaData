package main

import (
	"github.com/MLinek/TibiaData/internal"
	"github.com/MLinek/TibiaData/internal/models"
	"github.com/MLinek/TibiaData/pkg"
	"time"
)

type config struct {
	Lists []struct {
		Name   string   `json:"name"`
		Guilds []string `json:"guilds"`
	} `json:"lists"`
}

type export struct {
	Lists []exportList `json:"lists"`
}

type exportList struct {
	Name   string        `json:"name"`
	Guilds []exportGuild `json:"guilds"`
}
type exportGuild struct {
	Name    string         `json:"name"`
	Members []exportMember `json:"members"`
}

type exportMember struct {
	Name     string `json:"name"`
	Vocation string `json:"vocation"`
	Level    int    `json:"level"`
}

func main() {
	go startPolling()
	select {}
}

func startPolling() {
	for range time.Tick(60 * time.Second) {
		var err error
		cfg, err := loadConfig("config/guildonlinemonitor.json")
		if err == nil {
			err := updateExportedList(cfg)
			if err != nil {
				print(err.Error())
			}
		} else {
			print(err.Error())
		}
	}
}

func loadConfig(path string) (config, error) {
	var cfg config
	err := internal.ImportJsonFromFile(&cfg, path)
	if err != nil {
		return config{}, err
	}
	return cfg, err
}

func updateExportedList(cfg config) error {
	export, err := generateExport(cfg)
	if err != nil {
		return err
	}
	return internal.ExportJsonToFile(export, "dist/guildsonlinemonitor.json")
}

func generateExport(cfg config) (export, error) {
	var result = export{}
	for _, list := range cfg.Lists {
		rl, err := generateExportList(list.Name, list.Guilds)
		if err != nil {
			return export{}, err
		}
		result.Lists = append(result.Lists, rl)
	}
	return result, nil
}

func generateExportList(name string, guilds []string) (exportList, error) {
	var rl = exportList{
		Name:   name,
		Guilds: make([]exportGuild, len(guilds)),
	}
	for i, gName := range guilds {
		online, err := fetchGuildOnlinePlayers(gName)
		if err != nil {
			return exportList{}, err
		}
		rl.Guilds[i] = exportGuild{
			Name:    gName,
			Members: online,
		}
	}
	return rl, nil
}

func fetchGuildOnlinePlayers(name string) ([]exportMember, error) {
	fetcher := pkg.GuildsFetcher{}
	guild, err := fetcher.ByName(name)
	if err != nil {
		return []exportMember{}, err
	}
	online, err := guild.OnlinePlayers()
	if err != nil {
		return []exportMember{}, err
	}
	return membersToExport(online), nil
}

func membersToExport(members []models.Member) []exportMember {
	export := make([]exportMember, len(members))
	for i, m := range members {
		export[i] = exportMember{Name: m.Name, Vocation: m.Vocation, Level: m.Level}
	}
	return export
}
