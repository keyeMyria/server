package controllers

import (
	"context"
	"encoding/json"

	"github.com/douban-girls/server/app/gql"
	"github.com/douban-girls/server/app/utils"
	"github.com/graphql-go/graphql"
	"github.com/revel/revel"
)

type GraphQLController struct {
	*revel.Controller
}

// Fetch is Get method for fetch data ^_^
func (g *GraphQLController) Fetch() revel.Result {
	query := g.Params.Get("query")

	params := graphql.Params{
		Schema:        gql.GraphQLSchema,
		RequestString: query,
	}
	result := graphql.Do(params)
	return g.RenderJSON(result)
}

type pgd struct {
	Query         string
	Variables     map[string]interface{}
	OperationName string
}

func (g *GraphQLController) FetchByPost() revel.Result {
	var postedData pgd

	revel.INFO.Println(string(g.Params.JSON))

	if err := json.Unmarshal(g.Params.JSON, &postedData); err != nil {
		revel.INFO.Println(string(g.Params.JSON))
		return g.RenderJSON(utils.Response(500, nil, err))
	}

	params := graphql.Params{
		Schema:         gql.GraphQLSchema,
		RequestString:  postedData.Query,
		VariableValues: postedData.Variables,
		Context:        context.WithValue(context.Background(), "controller", g.Controller),
	}
	result := graphql.Do(params)
	return g.RenderJSON(result)
}
