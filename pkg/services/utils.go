package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/e-mbrown/rollWOD/graph/model"
	"github.com/e-mbrown/rollWOD/pkg/seed"
	guuid "github.com/google/uuid"
)

func GenUUID() *string {
	id := fmt.Sprintf(`%v`, guuid.New())
	return &id

}

// TODO: Maybe instead of returning context, specialize to take
// arg and return value
func ForcedResolverParentContext(ctx context.Context) graphql.FieldContext {
	return *graphql.GetFieldContext(ctx).Parent.Parent
}

// entrytoModelEntry is a util function that converts []seed.Entry
// to []model.GeneralInfo
func EntrytoModelGenInfo(data map[string]seed.Entry) map[string]*model.GeneralInfo {
	entries := map[string]*model.GeneralInfo{}

	for _, v := range data {
		name := strings.ToLower(v.Name)
		entries[name] = &model.GeneralInfo{
			ID:          GenUUID(),
			Name:        name,
			Description: seed.CleanDesc(v.Description),
		}
	}

	return entries
}

func ArchtoModelArch(data []seed.Archetypes) map[string]*model.Archetypes {
	entries := map[string]*model.Archetypes{}

	for _, arch := range data {
		name := strings.ToLower(arch.Name)
		entries[name] = &model.Archetypes{
			ID:          GenUUID(),
			Name:        name,
			Description: arch.Description,
			Sys:         arch.Sys,
		}
	}
	return entries
}

func ChartoModelChar(data map[string]map[string]seed.Characteristic) map[string]*model.Characteristic {
	entries := map[string]*model.Characteristic{}

	for k, v := range seed.InfoMap {
		for ik, iv := range v {
			name := strings.ToLower(ik)
			entries[name] = &model.Characteristic{
				ID:          GenUUID(),
				Name:        name,
				Type:        model.CharType(k),
				Description: seed.CleanDesc(iv.BaseDesc),
				Descbyval:   seed.ValDescString(iv.ValDesc),
			}
		}
	}
	return entries
}

// SecttoModelSect & DisciplinetoModelDiscipline needs to take a resolver so that duplicate Sects arent created when making clans
func SecttoModelSect(data map[string]seed.Sect) (map[string]*model.Sect, map[string]*model.Title) {
	sects, titles := map[string]*model.Sect{}, map[string]*model.Title{}
	for _, v := range seed.SectMap {
		tempT := []*model.Title{}

		for _, iv := range v.Titles {
			tName := strings.ToLower(iv.Name)
			title := &model.Title{
				ID:          GenUUID(),
				Name:        tName,
				Description: seed.CleanDesc(iv.Description),
			}
			tempT = append(tempT, title)
			titles[tName] = title
		}

		sects[v.Name] = &model.Sect{
			ID:          GenUUID(),
			Name:        v.Name,
			Description: seed.CleanDesc(v.Description),
			Practices:   v.Practices,
			Rituals:     v.Rituals,
			Titles:      tempT,
		}
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
				Description: seed.CleanDesc(v.Description),
			}
		}

		trads = append(
			trads,
			&model.Tradition{
				ID:          GenUUID(),
				Name:        v.Name,
				Description: seed.CleanDesc(v.Description),
				Traditions:  tempT,
			},
		)
		entries = append(entries, tempT...)
	}

	return trads, entries
}

func DisciplinetoModelDiscipline(data map[string]seed.Discipline) (map[string]*model.Discipline, map[string]*model.DiscAbilities) {
	entries := map[string]*model.Discipline{}
	abEntries := map[string]*model.DiscAbilities{}

	for _, disc := range data {
		name := strings.ToLower(disc.Name)
		tempAb := []*model.DiscAbilities{}
		for k, v := range disc.Abilities {
			abName := strings.ToLower(k)
			ab := &model.DiscAbilities{
				ID:          GenUUID(),
				Name:        abName,
				Description: seed.CleanDesc(v.BaseDesc),
				Lvl:         v.Lvl,
				System:      seed.ValDescString(v.System),
			}
			tempAb = append(tempAb, ab)

			abEntries[abName] = ab
		}

		entries[disc.Name] = &model.Discipline{
			ID:          GenUUID(),
			Name:        name,
			Description: seed.CleanDesc(disc.Description),
			Abilities:   tempAb,
		}

	}
	return entries, abEntries
}

// Always run after Sects and disciplines are made.
func ClantoModelClan(allSects map[string]*model.Sect, allDisc map[string]*model.Discipline, clanMap map[string]seed.Clan) map[string]*model.Clan {
	entries := map[string]*model.Clan{}

	for _, v := range clanMap {
		sect := []*model.Sect{}
		disc := []*model.Discipline{}
		for _, s := range v.AssociatedSect {
			sect = append(sect, allSects[(s)])
		}

		for _, d := range v.Discipline {
			disc = append(disc, allDisc[(d.Name)])
		}

		entries[v.Name] = &model.Clan{
			ID:             GenUUID(),
			Name:           v.Name,
			Description:    seed.CleanDesc(v.Description),
			Appearance:     v.Appearance,
			Associatedsect: sect,
			Haven:          v.Haven,
			Background:     v.Background,
			Character:      v.Character,
			Discipline:     disc,
			Weakness:       v.Weakness,
			Organizations:  &v.Organizations,
			Strongholds:    v.Strongholds,
			Ishighclan:     &v.IsHighClan,
		}
	}
	return entries
}
