package graph

import (
	"github.com/e-mbrown/rollWOD/pkg/seed"
	"github.com/e-mbrown/rollWOD/pkg/services"
)

func hydrateResolver(r Resolver) {
	r.genInfo = services.EntrytoModelGenInfo(seed.EntryMap)
	r.archetypes = services.ArchtoModelArch(seed.ArchMap)
	r.characteristics = services.ChartoModelChar(seed.InfoMap)

	sects, titles := services.SecttoModelSect(seed.SectMap)
	r.titles = titles
	r.sects = sects

	//ToDo reformat
	disc, discab := services.DisciplinetoModelDiscipline(seed.DisciplineMap)
	r.disciplines = disc
	r.discAbilities = discab

	traditions, gen := services.TradtoModelTrad(seed.TraditionMap)
	r.traditions = append(r.traditions, traditions...)
	for _, v := range gen {
		r.genInfo[v.Name] = v
	}

	// Never run clans before Sects and Disciplines are done.
	r.clans = services.ClantoModelClan(r.sects, r.disciplines, seed.ClanMap)
}
