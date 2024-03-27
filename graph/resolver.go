package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/e-mbrown/rollWOD/graph/model"
	"github.com/e-mbrown/rollWOD/pkg/seed"
	"github.com/e-mbrown/rollWOD/pkg/services"
)

type Resolver struct {
	// Ideal query searches by name
	characteristics map[string]*model.Characteristic
	sects           map[string]*model.Sect
	disciplines     map[string]*model.Discipline
	clans           map[string]*model.Clan
	genInfo         map[string]*model.GeneralInfo
	discAbilities   map[string]*model.DiscAbilities

	//Undetermined
	titles          map[string]*model.Title
	traditions      []*model.Tradition
}

func NewResolver() Config {
	r := Resolver{}

	r.genInfo = services.EntrytoModelGenInfo(seed.EntryMap)
	r.characteristics = services.ChartoModelChar(seed.InfoMap)

	sects, titles := services.SecttoModelSect(seed.SectMap)
	r.titles = titles
	r.sects = sects

	//ToDo reformat
	disc, discab := services.DisciplinetoModelDiscipline(seed.DisciplineMap)
	r.disciplines = disc
	r.discAbilities =  discab

	traditions, gen := services.TradtoModelTrad(seed.TraditionMap)
	r.traditions = append(r.traditions, traditions...)
	for _, v := range gen {
		r.genInfo[v.Name] = v
	}

	// Never run clans before Sects and Disciplines are done.
	r.clans = services.ClantoModelClan(r.sects, r.disciplines, seed.ClanMap)

	return Config{
		Resolvers: &r,
	}
}
