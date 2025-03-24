package graph

// TODO: Since seed is not publically available switch to sql calls
// or something
func hydrateResolver(r Resolver) Resolver {
	// r.genInfo = services.EntrytoModelGenInfo(seed.EntryMap)
	// r.archetypes = services.ArchtoModelArch(seed.ArchMap)
	// r.characteristics = services.ChartoModelChar(seed.InfoMap)

	// r.sects, r.titles = services.SecttoModelSect(seed.SectMap)

	// //ToDo reformat
	// disc, discab := services.DisciplinetoModelDiscipline(seed.DisciplineMap)
	// r.disciplines = disc
	// r.discAbilities = discab

	// traditions, gen := services.TradtoModelTrad(seed.TraditionMap)
	// r.traditions = append(r.traditions, traditions...)
	// for _, v := range gen {
	// 	r.genInfo[v.Name] = v
	// }

	// r.clans = services.ClantoModelClan(r.sects, r.disciplines, seed.ClanMap)
	return r
}
