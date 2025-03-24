package graph_test

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/e-mbrown/rollWOD/graph"
	"github.com/e-mbrown/rollWOD/graph/model"
	"github.com/stretchr/testify/require"
)

func TestConpendium(t *testing.T) {
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})))

	t.Run("Testing Sect & Title creation", func(t *testing.T) {
		var resp struct {
			Time string
		}
		var respB struct {
			model.Sect
		}

		testSect := model.Sect{
			Name:        "Test",
			Description: "This is a test Sect",
			Practices:   "They test",
			Rituals:     "They test ritualistically",
			Titles: []*model.Title{
				{
					Name:        "Test Title",
					Description: "Im a title nested within Sect, as I should be",
				},
			},
		}

		// mutQuery := fmt.Sprintf(
		// 	`mutation f{ 
		// 		createSect(name: %s, description:%s, practices:%s, rituals:%s, titles: [{name:%s, description:%s}])} { time }`,
		// 	testSect.Name,
		// 	testSect.Description,
		// 	testSect.Practices,
		// 	testSect.Rituals,
		// 	testSect.Titles[0].Name,
		// 	testSect.Titles[0].Description,
		// )
		mutQuery :=`mutation f{createSect(name: "Test")} { time }`
	

		query := `
			query a() 
				sect {
					name
					description
					practices
					rituals
					titles
				}
			}
		`
		c.MustPost(mutQuery, &resp)
		c.MustPost(query, &respB)

		fmt.Print(resp, respB)
		require.Equal(t, testSect, respB)
	})

}
