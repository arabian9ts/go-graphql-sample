package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var userQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userType,
				Description: "Get User By ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						user, err := NewUserRepository().Find(int64(id))
						if user == nil {
							return nil, err
						}
						return user, nil
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "Get User List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return NewUserRepository().FindAll(), nil
				},
			},
		},
	},
)

func executeUserQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("err = %v", result.Errors)
	}
	return result
}

var userSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    userQuery,
		Mutation: nil,
	},
)

var messageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var messageQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type:        messageType,
				Description: "Get Messages By UserID",
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["user_id"].(int)
					if ok {
						user, err := NewMessageRepository().Find(int64(id))
						if user == nil {
							return nil, err
						}
						return user, nil
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(messageType),
				Description: "Get Message List",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return NewMessageRepository().FindAll(), nil
				},
			},
		},
	},
)

func executeMessageQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("err = %v", result.Errors)
	}
	return result
}

var messageSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    messageQuery,
		Mutation: nil,
	},
)
