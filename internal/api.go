package internal

import (
	"encoding/json"
	"github.com/MLinek/TibiaData/internal/models"
	"strings"
)

type API struct {
}

func (api API) Guild(name string) (models.Guild, error) {
	f := &Fetcher{}
	body, err := f.fetch("https://api.tibiadata.com/v3/guild/" + strings.Replace(name, " ", "%20", -1))
	if err != nil {
		return models.Guild{}, err
	}
	var guilds models.Guilds
	err = json.Unmarshal(body, &guilds)
	if err != nil {
		return models.Guild{}, err
	}
	return guilds.Guilds.Guild, nil
}
