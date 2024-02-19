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
	titles          []*model.Title
	traditions      []*model.Tradition
	genInfo         []*model.GeneralInfo
}

func NewResolver() Config {
	r := Resolver{}

	for k, v := range seed.InfoMap {
		for ik, iv := range v {
			entry := &model.Characteristic{
				ID:          services.GenUUID(),
				Name:        ik,
				Type:        model.CharType(k),
				Description: iv.BaseDesc,
				DescbyVal:   iv.ValDescString(),
			}
			r.characteristics = append(r.characteristics, entry)
		}
	}

	for _, v := range seed.SectMap {
		titles := []*model.Title{}
		for _, iv := range v.Titles {
			entry := &model.Title{
				ID:          services.GenUUID(),
				Name:        iv.Name,
				Description: iv.Description,
			}
			r.titles = append(r.titles, entry)
			titles = append(titles, entry)
		}

		entry := &model.Sect{
			ID:          services.GenUUID(),
			Name:        v.Name,
			Description: v.Description,
			Practices:   v.Practices,
			Rituals:     v.Rituals,
			Titles:      titles,
		}

		r.sects = append(r.sects, entry)
	}

	for _, v := range seed.TraditionMap {
		trads := []*model.GeneralInfo{}

		for _, v := range v.Tradition {
			info := &model.GeneralInfo{
				ID:          services.GenUUID(),
				Name:        v.Name,
				Description: v.Description,
			}

			trads = append(trads, info)
			r.genInfo = append(r.genInfo, info)
		}

		entry := &model.Tradition{
			ID:          services.GenUUID(),
			Name:        v.Name,
			Description: v.Description,
			Traditions:  trads,
		}
		r.traditions = append(r.traditions, entry)
	}

	return Config{
		Resolvers: &r,
	}
}
