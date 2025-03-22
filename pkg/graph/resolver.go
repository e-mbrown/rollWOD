package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/e-mbrown/rollWOD/graph/model"
)

type Resolver struct {
	// Ideal query searches by name
	characteristics map[string]*model.Characteristic
	sects           map[string]*model.Sect
	disciplines     map[string]*model.Discipline
	clans           map[string]*model.Clan
	genInfo         map[string]*model.GeneralInfo
	discAbilities   map[string]*model.DiscAbilities
	archetypes      map[string]*model.Archetypes

	//Undetermined
	titles     map[string]*model.Title
	traditions []*model.Tradition
}

func NewResolver() Config {
	r := hydrateResolver(Resolver{})

	return Config{
		Resolvers: &r,
	}
}
