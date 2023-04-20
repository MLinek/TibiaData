package pkg

import (
	"github.com/MLinek/TibiaData/internal"
	"github.com/MLinek/TibiaData/internal/models"
)

type GuildsFetcher struct{}

func (f GuildsFetcher) ByName(name string) (models.Guild, error) {
	api := internal.API{}
	return api.Guild(name)
}
