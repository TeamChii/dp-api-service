package Schema

import "github.com/graphql-go/graphql"

/*
var LoginTypeRes = graphql.NewObject(graphql.ObjectConfig{
	Name: "loginRes",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: UserType,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"statusCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageAbbr": &graphql.Field{
			Type: graphql.String,
		},
		"messageDesc": &graphql.Field{
			Type: graphql.String,
		},
	}})
var LoginLearnerTypeRes = graphql.NewObject(graphql.ObjectConfig{
	Name: "loginLearnerRes",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: LearnerType,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"statusCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageAbbr": &graphql.Field{
			Type: graphql.String,
		},
		"messageDesc": &graphql.Field{
			Type: graphql.String,
		},
	}})
*/

var MerchantTypeRes = graphql.NewObject(graphql.ObjectConfig{
	Name: "merchantMasterRes",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.NewList(MerchantType),
		},
		"messageAbbr": &graphql.Field{
			Type: graphql.String,
		},
		"messageCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageDesc": &graphql.Field{
			Type: graphql.String,
		},
		"statusCode": &graphql.Field{
			Type: graphql.String,
		},
	}})
var ContainerTypeRes = graphql.NewObject(graphql.ObjectConfig{
	Name: "containerMasterRes",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.NewList(ContainerType),
		},
		"messageAbbr": &graphql.Field{
			Type: graphql.String,
		},
		"messageCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageDesc": &graphql.Field{
			Type: graphql.String,
		},
		"statusCode": &graphql.Field{
			Type: graphql.String,
		},
		"paging": &graphql.Field{
			Type: PagingType,
		},
	}})
var ContainerTypeTypeRes = graphql.NewObject(graphql.ObjectConfig{
	Name: "containerTypeMasterRes",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.NewList(ContainerTypeType),
		},
		"messageAbbr": &graphql.Field{
			Type: graphql.String,
		},
		"messageCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageDesc": &graphql.Field{
			Type: graphql.String,
		},
		"statusCode": &graphql.Field{
			Type: graphql.String,
		},
		"paging": &graphql.Field{
			Type: PagingType,
		},
	}})
var MerchantMapCustomerTypeRes = graphql.NewObject(graphql.ObjectConfig{
	Name: "merchantMappcustomerMasterRes",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: graphql.NewList(MerchantMapCustomerType),
		},
		"messageAbbr": &graphql.Field{
			Type: graphql.String,
		},
		"messageCode": &graphql.Field{
			Type: graphql.String,
		},
		"messageDesc": &graphql.Field{
			Type: graphql.String,
		},
		"statusCode": &graphql.Field{
			Type: graphql.String,
		},
		"paging": &graphql.Field{
			Type: PagingType,
		},
	}})
