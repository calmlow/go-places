package places

import (
	"github.com/calmlow/go-places/internal/config"
	"github.com/calmlow/go-places/internal/types"
)

func GetPlaces() ([]types.Place, error) {
	localConfig, err := config.ReadYamlConfigFile()
	return filter(&localConfig.Places, false), err
}

func GetHiddenPlaces() ([]types.Place, error) {
	localConfig, err := config.ReadYamlConfigFile()

	return filter(&localConfig.Places, true), err
}

func GetFullConfig() (config.LocalConfig, error) {
	return config.ReadYamlConfigFile()
}

func filter(places *[]types.Place, onlyHidden bool) []types.Place {
	var filteredPlaces []types.Place
	for _, place := range *places {
		if place.Hidden == onlyHidden {
			filteredPlaces = append(filteredPlaces, place)
		}
	}
	return filteredPlaces
}
