package graph

import (
	"context"

	"github.com/e-mbrown/rollWOD/graph/model"
	"github.com/e-mbrown/rollWOD/pkg/seed"
)

// Seed is the resolver for the seed field.
func (r *mutationResolver) Seed(ctx context.Context, input *string) (string, error) {
	baseAbilities := seed.InfoMap
	// backgrounds := seed.BackgroundMap

	for k, v := range baseAbilities {
		for ik, iv := range v {
				entry := makeModelC(k, ik,iv)
			_, err := r.Mutation().CreateCharacteristic(ctx, entry)
			if err != nil {
				return "false", err
			}
		}
	}

	return "true", nil
}



func makeModelC(typ string, ik string, data seed.Characteristic) model.NewCharacteristic{
	entry := model.NewCharacteristic{
		Name:        ik,
		Type:        model.CharType(typ),
		Description: data.BaseDesc,
		ValDesc:     data.ValDescString(),
	}
	return entry
}