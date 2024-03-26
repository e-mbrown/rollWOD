package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/e-mbrown/rollWOD/graph/model"
	"github.com/e-mbrown/rollWOD/pkg/seed"
	"github.com/e-mbrown/rollWOD/pkg/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	characteristics []*model.Characteristic
	sects           []*model.Sect
	disciplines     []*model.Discipline
	discAbilities   []*model.DiscAbilities
	titles          []*model.Title
	traditions      []*model.Tradition
	genInfo         []*model.GeneralInfo
	clans           []*model.Clan
}

func NewResolver() Config {
	r := Resolver{}

	r.genInfo = append(
		r.genInfo,
		services.EntrytoModelGenInfo(seed.EntryMap)...,
	)
	r.characteristics = append(
		r.characteristics,
		services.ChartoModelChar(seed.InfoMap)...,
	)

	sects, titles := services.SecttoModelSect(seed.SectMap)
	r.titles = append(r.titles, titles...)
	r.sects = append(r.sects, sects...)

	disc, discab := services.DisciplinetoModelDiscipline(seed.DisciplineMap)
	r.disciplines = append(r.disciplines, disc...)
	r.discAbilities = append(r.discAbilities, discab...)

	traditions, gen := services.TradtoModelTrad(seed.TraditionMap)
	r.traditions = append(r.traditions, traditions...)
	r.genInfo = append(r.genInfo, gen...)

	// Never run clans before Sects and Disciplines are done.
	r.clans = append(
		r.clans,
		services.ClantoModelClan(r.sects, r.disciplines, seed.ClanMap)...)

	return Config{
		Resolvers: &r,
	}
}
