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

//entrytoModelEntry is a util function that converts []seed.Entry
// to []model.GeneralInfo 
func EntrytoModelGenInfo(data []seed.Entry) []*model.GeneralInfo {
	entries := make([]*model.GeneralInfo, len(data))
	
	for _,v := range data {
		entries = append(
			entries,
			&model.GeneralInfo{
				ID: GenUUID(),
				Name: v.Name,
				Description: v.Description,
			})
	}

	return entries
}
