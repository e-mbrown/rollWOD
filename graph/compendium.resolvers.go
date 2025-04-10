package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/e-mbrown/rollWOD/graph/model"
	"github.com/e-mbrown/rollWOD/pkg/services"
)

// TODO(maybe): Create query related argument constants so that they arent just hardcoded all over. Also look into N+1 problem.
func (r *clanResolver) Associatedsect(ctx context.Context, obj *model.Clan) ([]*model.Sect, error) {
	var res []*model.Sect
	pCtx := services.ForcedResolverParentContext(ctx)
	if pCtx.Args["expand"].(bool) {
		res = append(res, obj.Associatedsect...)
	}

	return res, nil
}

// Discipline is the resolver for the discipline field.
func (r *clanResolver) Discipline(ctx context.Context, obj *model.Clan) ([]*model.Discipline, error) {
	var res []*model.Discipline
	pCtx := services.ForcedResolverParentContext(ctx)
	if pCtx.Args["expand"].(bool) {
		res = append(res, obj.Discipline...)
	}

	return res, nil
}

// Subclan is the resolver for the subclan field.
func (r *clanResolver) Subclan(ctx context.Context, obj *model.Clan) ([]*model.Clan, error) {
	var res []*model.Clan
	pCtx := services.ForcedResolverParentContext(ctx)
	if pCtx.Args["expand"].(bool) {
		res = append(res, obj.Subclan...)
	}
	return res, nil
}

// Abilities is the resolver for the abilities field.
func (r *disciplineResolver) Abilities(ctx context.Context, obj *model.Discipline) ([]*model.DiscAbilities, error) {
	var res []*model.DiscAbilities
	pCtx := services.ForcedResolverParentContext(ctx)
	if pCtx.Args["expand"].(bool) {
		res = append(res, obj.Abilities...)
	}

	return res, nil
}

// Archetypes is the resolver for the archetypes field.
func (r *queryResolver) Archetypes(ctx context.Context) ([]*model.Archetypes, error) {
	res := []*model.Archetypes{}
	for _, v := range r.archetypes {
		res = append(res, v)
	}
	return res, nil
}

// Characteristics is the resolver for the characteristics field.
func (r *queryResolver) Characteristics(ctx context.Context) ([]*model.Characteristic, error) {
	res := []*model.Characteristic{}
	for _, v := range r.characteristics {
		res = append(res, v)
	}
	return res, nil
}

// Clans is the resolver for the clans field.
func (r *queryResolver) Clans(ctx context.Context, expand bool) ([]*model.Clan, error) {
	var res []*model.Clan

	for _, v := range r.clans {
		res = append(res, v)
	}

	return res, nil
}

// Disciplines is the resolver for the disciplines field.
func (r *queryResolver) Disciplines(ctx context.Context, expand bool) ([]*model.Discipline, error) {
	var res []*model.Discipline
	for _, v := range r.disciplines {
		res = append(res, v)
	}

	return res, nil
}

// Sects is the resolver for the sects field.
func (r *queryResolver) Sects(ctx context.Context, expand bool) ([]*model.Sect, error) {
	res := []*model.Sect{}
	for _, v := range r.sects {
		res = append(res, v)
	}
	return res, nil
}

// Titles is the resolver for the titles field.
func (r *queryResolver) Titles(ctx context.Context) ([]*model.Title, error) {
	res := []*model.Title{}
	for _, v := range r.titles {
		res = append(res, v)
	}
	return res, nil
}

// Traditions is the resolver for the traditions field.
func (r *queryResolver) Traditions(ctx context.Context) ([]*model.Tradition, error) {
	return r.traditions, nil
}

// GetSect is the resolver for the getSect field.
func (r *queryResolver) GetSect(ctx context.Context, name []*string, expand bool) ([]*model.Sect, error) {
	res := []*model.Sect{}

	for _, v := range name {
		val, ok := r.sects[*v]
		if ok {
			res = append(res, val)
		} else {
			graphql.AddErrorf(ctx, "Non-fatal error: parameter '%s' is not a sect", *v)
		}
	}

	if len(res) == 0 {
		return nil, graphql.GetErrors(ctx)
	}

	return res, nil
}

// GetTradition is the resolver for the getTradition field.
func (r *queryResolver) GetTradition(ctx context.Context, name *string) (*model.Tradition, error) {
	var res *model.Tradition
	for _, trad := range r.traditions {
		if trad.Name == *name {
			res = trad
		}
	}

	if res == nil {
		graphql.AddErrorf(ctx, "Non-fatal error: parameter '%s' is not a tradition", *name)
		return nil, graphql.GetErrors(ctx)
	}

	return res, nil
}

// GetDiscipline is the resolver for the getDiscipline field.
func (r *queryResolver) GetDiscipline(ctx context.Context, name []*string, expand bool) ([]*model.Discipline, error) {
	var res []*model.Discipline
	for _, v := range name {
		val, ok := r.disciplines[*v]
		if ok {
			res = append(res, val)
		} else {
			graphql.AddErrorf(ctx, "Non-fatal error: parameter '%s' is not a discipline", *v)
		}
	}

	if len(res) == 0 {
		return nil, graphql.GetErrors(ctx)
	}

	return res, nil
}

// GetClan is the resolver for the getClan field.
func (r *queryResolver) GetClan(ctx context.Context, name []*string, expand bool) ([]*model.Clan, error) {
	var res []*model.Clan
	for _, v := range name {
		val, ok := r.clans[*v]
		if ok {
			res = append(res, val)
		} else {
			graphql.AddErrorf(ctx, "Non-fatal error: parameter '%s' is not a discipline", *v)
		}
	}

	if len(res) == 0 {
		return nil, graphql.GetErrors(ctx)
	}

	return res, nil
}

// GenInfoByName is the resolver for the GenInfoByName field.
func (r *queryResolver) GenInfoByName(ctx context.Context, name *string) (*model.GeneralInfo, error) {
	res, ok := r.genInfo[strings.ToLower(*name)]
	if !ok {
		graphql.AddErrorf(ctx, "%s: is not a entry in the conpendium", *name)
		return nil, graphql.GetErrors(ctx)
	}
	return res, nil
}

// GenByID is the resolver for the GenByID field.
func (r *queryResolver) GenByID(ctx context.Context, id *string) (*model.GeneralInfo, error) {
	panic(fmt.Errorf("not implemented: GenByID - GenByID"))
}

// GetDiscAb is the resolver for the getDiscAb field.
func (r *queryResolver) GetDiscAb(ctx context.Context, name []*string) ([]*model.DiscAbilities, error) {
	var res []*model.DiscAbilities
	for _, v := range name {
		ab, ok := r.discAbilities[strings.ToLower(*v)]
		if !ok {
			graphql.AddErrorf(ctx, "%s: is not a entry in the discipline ability", *v)
			return nil, graphql.GetErrors(ctx)
		} else {
			res = append(res, ab)
		}
	}

	if len(res) == 0 {
		return nil, graphql.GetErrors(ctx)
	}

	return res, nil
}

// CharByType is the resolver for the charByType field.
func (r *queryResolver) CharByType(ctx context.Context, typeArg *model.CharType) ([]*model.Characteristic, error) {
	res := []*model.Characteristic{}
	for _, v := range r.characteristics {
		if v.Type == *typeArg {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		graphql.AddErrorf(ctx, "Non-fatal error: %s is not a charType", typeArg)
		return nil, graphql.GetErrors(ctx)
	}
	return res, nil
}

// Titles is the resolver for the titles field.
func (r *sectResolver) Titles(ctx context.Context, obj *model.Sect) ([]*model.Title, error) {
	var res []*model.Title
	arg := services.ForcedResolverParentContext(ctx).Args["expand"]
	if arg != nil && arg.(bool) {
		res = append(res, obj.Titles...)
	}

	return res, nil
}

// Clan returns ClanResolver implementation.
func (r *Resolver) Clan() ClanResolver { return &clanResolver{r} }

// Discipline returns DisciplineResolver implementation.
func (r *Resolver) Discipline() DisciplineResolver { return &disciplineResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Sect returns SectResolver implementation.
func (r *Resolver) Sect() SectResolver { return &sectResolver{r} }

type clanResolver struct{ *Resolver }
type disciplineResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type sectResolver struct{ *Resolver }
