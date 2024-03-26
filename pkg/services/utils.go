package services

import (
	"fmt"

	"github.com/e-mbrown/rollWOD/graph/model"
	"github.com/e-mbrown/rollWOD/pkg/seed"
	guuid "github.com/google/uuid"
)

func GenUUID() *string {
	id := fmt.Sprintf(`%v`, guuid.New())
	return &id

}

// entrytoModelEntry is a util function that converts []seed.Entry
// to []model.GeneralInfo
func EntrytoModelGenInfo(data map[string]seed.Entry) []*model.GeneralInfo {
	entries := make([]*model.GeneralInfo, len(data))

	for _, v := range data {
		entries = append(
			entries,
			&model.GeneralInfo{
				ID:          GenUUID(),
				Name:        v.Name,
				Description: v.Description,
			})
	}

	return entries
}

func ChartoModelChar(data map[string]map[string]seed.Characteristic) []*model.Characteristic {
	entries := make([]*model.Characteristic, len(data))

	for k, v := range seed.InfoMap {
		for ik, iv := range v {
			entries = append(entries,
				&model.Characteristic{
					ID:          GenUUID(),
					Name:        ik,
					Type:        model.CharType(k),
					Description: iv.BaseDesc,
					DescbyVal:   seed.ValDescString(iv.ValDesc),
				})
		}
	}
	return entries
}

// SecttoModelSect & DisciplinetoModelDiscipline needs to take a resolver so that duplicate Sects arent created when making clans
func SecttoModelSect(data map[string]seed.Sect) ([]*model.Sect, []*model.Title) {
	sects := []*model.Sect{}
	titles := []*model.Title{}

	for _, v := range seed.SectMap {
		tempT := []*model.Title{}

		for _, iv := range v.Titles {
			tempT = append(
				titles,
				&model.Title{
					ID:          GenUUID(),
					Name:        iv.Name,
					Description: iv.Description,
				})
		}

		sects = append(
			sects,
			&model.Sect{
				ID:          GenUUID(),
				Name:        v.Name,
				Description: v.Description,
				Practices:   v.Practices,
				Rituals:     v.Rituals,
				Titles:      tempT,
			})
			titles = append(titles, tempT...)
	}

	return sects, titles
}

func TradtoModelTrad(data []seed.Traditions) ([]*model.Tradition, []*model.GeneralInfo) {
	entries := []*model.GeneralInfo{}
	trads := []*model.Tradition{}

	for _, v := range data {
		tempT := make([]*model.GeneralInfo, len(v.Tradition))
		for i, v := range v.Tradition {
			tempT[i] = &model.GeneralInfo{
					ID:          GenUUID(),
					Name:        v.Name,
					Description: v.Description,
				}
		}

		trads = append(
			trads,
			&model.Tradition{
				ID:          GenUUID(),
				Name:        v.Name,
				Description: v.Description,
				Traditions:  tempT,
			},
		)
		entries = append(entries, tempT...)
	}

	return trads, entries
}

func DisciplinetoModelDiscipline(data map[string]seed.Discipline) ([]*model.Discipline, []*model.DiscAbilities) {
	entries := []*model.Discipline{}
	discAb := []*model.DiscAbilities{}

	for _, disc := range data {
		tempAb := []*model.DiscAbilities{}
		for k, v := range disc.Abilities {
			tempAb = append(tempAb, &model.DiscAbilities{
				ID: GenUUID(),
				Name:        k,
				Description: v.BaseDesc,
				Lvl:         v.Lvl,
				System:      seed.ValDescString(v.System),
			})
		}

		entries = append(entries, &model.Discipline{
			ID: GenUUID(),
			Name:        disc.Name,
			Description: disc.Description,
			Abilities:   tempAb,
		})
		discAb = append(discAb, tempAb...)
	}
	return entries, discAb
}

// Always run after Sects and disciplines are made.
func ClantoModelClan(allSects []*model.Sect, allDisc []*model.Discipline, clanMap map[string]seed.Clan) []*model.Clan {
	entries := []*model.Clan{}
	dCache := make(map[string]*model.Discipline, len(allDisc))
	sCache := make(map[string]*model.Sect, len(allSects))

	for _, v := range clanMap {
		sect := []*model.Sect{}
		disc := []*model.Discipline{}
		for _, s := range v.AssociatedSect {
			v, ok := sCache[s]
			if !ok {
				v = RangeSect(allSects, s)
				sCache[s] = v
			}

			sect = append(sect, v)
		}

		for _, d := range v.Discipline {
			v, ok := dCache[d.Name]
			if !ok {
				v = RangeDiscipline(allDisc, d.Name)
				dCache[d.Name] = v
			}

			disc = append(disc, v)
		}

		entries = append(
			entries,
			&model.Clan{
				ID:             GenUUID(),
				Name:           v.Name,
				Description:    v.Description,
				Appearance:     v.Appearance,
				AssociatedSect: sect,
				Haven:          v.Haven,
				Background:     v.Background,
				Character:      v.Character,
				Discipline:     disc,
				Weakness: v.Weakness,
				Organizations: &v.Organizations,
				Strongholds: v.Strongholds,
				IsHighClan: &v.IsHighClan,
			},
		)
	}
	return entries
}

func MakeModelDiscipline(data []seed.Discipline) []*model.Discipline {
	entries := make([]*model.Discipline, len(data))

	return entries
}

func RangeSect(allSect []*model.Sect, name string) *model.Sect {
	for _, sect := range allSect {
		if sect.Name == name {
			return sect
		}
	}

	return nil
}
func RangeDiscipline(allDisc []*model.Discipline, name string) *model.Discipline {
	for _, disc := range allDisc {
		if disc.Name == name {
			return disc
		}
	}

	return nil
}

/*
 Issue
	When making clans we have to associate sects with them as well as disciplines.
	If sects and disciplines are populated first we can use the resolver to get all sects and disciplines and associate them with one another.
*/
