package Schema

import (
	"../../th.co.droppoint/resolver"
	"github.com/graphql-go/graphql"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"loadMerchantById": &graphql.Field{
			Type: MerchantTypeRes,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"authorization": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetMerchantByIdResolver,
		},
		"laodContainerByMcId": &graphql.Field{
			Type: ContainerTypeRes,
			Args: graphql.FieldConfigArgument{
				"mc_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageNo": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"orderBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"authorization": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetContainerByMcIdResolver,
		},
		"laodContainerType": &graphql.Field{
			Type: ContainerTypeTypeRes,
			Args: graphql.FieldConfigArgument{

				"pageNo": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"orderBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"authorization": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetContainerTypeResolver,
		},
		"laodMerchantMapCustomer": &graphql.Field{
			Type: MerchantMapCustomerTypeRes,
			Args: graphql.FieldConfigArgument{
				"mc_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"cust_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageNo": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"orderBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"authorization": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetMerchantMapCustomerResolver,
		},
		/*"searchCustomer": &graphql.Field{
			Type: CustomerTypeRes,
			Args: graphql.FieldConfigArgument{
				"searchString": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageNo": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sortBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"orderBy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"authorization": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetCustomerSearchResolver,
		},*/
	}})

var GetSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: query,
})
