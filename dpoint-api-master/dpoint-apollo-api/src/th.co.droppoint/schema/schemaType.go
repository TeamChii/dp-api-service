package Schema

import "github.com/graphql-go/graphql"

var PagingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Paging",
		Fields: graphql.Fields{
			"pageNo": &graphql.Field{
				Type: graphql.Int,
			},
			"pageSize": &graphql.Field{
				Type: graphql.Int,
			},
			"totalRecord": &graphql.Field{
				Type: graphql.Int,
			},
			"totalPage": &graphql.Field{
				Type: graphql.Int,
			},
			"orderBy": &graphql.Field{
				Type: graphql.String,
			},
			"sortBy": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AlertMessageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AlertMessage",
		Fields: graphql.Fields{
			"messageCode": &graphql.Field{
				Type: graphql.String,
			},
			"messageDesc": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var CustomerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "customerMaster",
	Fields: graphql.Fields{
		"cust_id": &graphql.Field{
			Type: graphql.Int,
		},
		"cust_name": &graphql.Field{
			Type: graphql.String,
		},
		"cust_alt_name": &graphql.Field{
			Type: graphql.String,
		},
		"cust_language": &graphql.Field{
			Type: graphql.Int,
		},
		"lead_by": &graphql.Field{
			Type: graphql.String,
		},
		"cust_status": &graphql.Field{
			Type: graphql.Int,
		},
		"cust_dob": &graphql.Field{
			Type: graphql.String,
		},
		"cust_sex": &graphql.Field{
			Type: graphql.Int,
		},
		"cust_region": &graphql.Field{
			Type: graphql.Int,
		},
		"cust_mobile": &graphql.Field{
			Type: graphql.String,
		},
		"cust_mobile_ext": &graphql.Field{
			Type: graphql.String,
		},
		"cust_email": &graphql.Field{
			Type: graphql.String,
		},
		"cust_email_ext": &graphql.Field{
			Type: graphql.String,
		},
		"cust_facebook_id": &graphql.Field{
			Type: graphql.String,
		},
		"cust_line_id": &graphql.Field{
			Type: graphql.String,
		},
		"cust_twitter_id": &graphql.Field{
			Type: graphql.String,
		},
		"cust_google_id": &graphql.Field{
			Type: graphql.String,
		},
		"create_by": &graphql.Field{
			Type: graphql.String,
		},
		"update_by": &graphql.Field{
			Type: graphql.String,
		},
		"create_date": &graphql.Field{
			Type: graphql.String,
		},
		"update_date": &graphql.Field{
			Type: graphql.String,
		},
	}})

var MerchantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "merchantMaster",
	Fields: graphql.Fields{
		"mc_id": &graphql.Field{
			Type: graphql.Int,
		},
		"mc_name": &graphql.Field{
			Type: graphql.String,
		},
		"mc_email": &graphql.Field{
			Type: graphql.String,
		},
		"mc_email_ext": &graphql.Field{
			Type: graphql.String,
		},
		"mc_facebook_id": &graphql.Field{
			Type: graphql.String,
		},
		"mc_line_id": &graphql.Field{
			Type: graphql.String,
		},
		"mc_twitter_id": &graphql.Field{
			Type: graphql.String,
		},
		"mc_ig_id": &graphql.Field{
			Type: graphql.String,
		},
		"mc_google_id": &graphql.Field{
			Type: graphql.String,
		},
		"mc_phone": &graphql.Field{
			Type: graphql.String,
		},
		"mc_phone_ext": &graphql.Field{
			Type: graphql.String,
		},
		"mc_mob": &graphql.Field{
			Type: graphql.String,
		},
		"mc_mob_ext": &graphql.Field{
			Type: graphql.String,
		},
		"mc_region": &graphql.Field{
			Type: graphql.String,
		},
		"mc_language": &graphql.Field{
			Type: graphql.String,
		},
		"mc_country": &graphql.Field{
			Type: graphql.String,
		},
		"mc_address_1": &graphql.Field{
			Type: graphql.String,
		},
		"mc_address_2": &graphql.Field{
			Type: graphql.String,
		},
		"mc_key_contact": &graphql.Field{
			Type: graphql.String,
		},
		"mc_lat": &graphql.Field{
			Type: graphql.Int,
		},
		"mc_long": &graphql.Field{
			Type: graphql.Int,
		},
		"mc_status": &graphql.Field{
			Type: graphql.String,
		},
		"mc_business_cat": &graphql.Field{
			Type: graphql.String,
		},
		"mc_business_size": &graphql.Field{
			Type: graphql.String,
		},
		"mc_sale_volume": &graphql.Field{
			Type: graphql.Int,
		},
		"mc_currency": &graphql.Field{
			Type: graphql.String,
		},
		"mc_is_head_office": &graphql.Field{
			Type: graphql.String,
		},
		"mc_ref": &graphql.Field{
			Type: graphql.Int,
		},
		"create_by": &graphql.Field{
			Type: graphql.String,
		},
		"create_date": &graphql.Field{
			Type: graphql.String,
		},

		"update_by": &graphql.Field{
			Type: graphql.String,
		},
		"update_date": &graphql.Field{
			Type: graphql.String,
		},
		"mc_tax": &graphql.Field{
			Type: graphql.String,
		},
		"mc_detail": &graphql.Field{
			Type: graphql.String,
		},
		"package_id": &graphql.Field{
			Type: graphql.Int,
		},
		"package": &graphql.Field{
			Type: PackageType,
		},
	}})

var ContainerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "containerMaster",
	Fields: graphql.Fields{
		"container_id": &graphql.Field{
			Type: graphql.Int,
		},
		"container_type_id": &graphql.Field{
			Type: graphql.Int,
		},
		"containerType": &graphql.Field{
			Type: ContainerTypeType,
		},
		"mc_id": &graphql.Field{
			Type: graphql.Int,
		},
		"layout_template": &graphql.Field{
			Type: graphql.String,
		},
		"layout_color": &graphql.Field{
			Type: graphql.String,
		},
		"container_subject": &graphql.Field{
			Type: graphql.String,
		},
		"container_detail": &graphql.Field{
			Type: graphql.String,
		},
		"image_ref": &graphql.Field{
			Type: graphql.String,
		},
		"container_term_conditions": &graphql.Field{
			Type: graphql.String,
		},
		"expire_mode": &graphql.Field{
			Type: graphql.String,
		},
		"total_point": &graphql.Field{
			Type: graphql.Int,
		},
		"expire_date": &graphql.Field{
			Type: graphql.String,
		},
		"issue_date": &graphql.Field{
			Type: graphql.String,
		},
		"reward": &graphql.Field{
			Type: graphql.NewList(ContainerRewardType),
		},
	}})

var ContainerTypeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "containerTypeMaster",
	Fields: graphql.Fields{
		"container_type_id": &graphql.Field{
			Type: graphql.Int,
		},
		"container_type_name": &graphql.Field{
			Type: graphql.String,
		},
		"container_type_code": &graphql.Field{
			Type: graphql.String,
		},
		"ref_id": &graphql.Field{
			Type: graphql.Int,
		},
	}})
var ContainerRewardType = graphql.NewObject(graphql.ObjectConfig{
	Name: "containerRewardMaster",
	Fields: graphql.Fields{
		"container_id": &graphql.Field{
			Type: graphql.Int,
		},
		"point_amt": &graphql.Field{
			Type: graphql.Int,
		},
		"reward_detail": &graphql.Field{
			Type: graphql.String,
		},
	}})
var MerchantMapCustomerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "merchantMappcustomerMaster",
	Fields: graphql.Fields{
		"mc_id": &graphql.Field{
			Type: graphql.Int,
		},
		"cust_id": &graphql.Field{
			Type: graphql.Int,
		},
		"container_id": &graphql.Field{
			Type: graphql.Int,
		},
		"container": &graphql.Field{
			Type: ContainerType,
		},
		"expire_date": &graphql.Field{
			Type: graphql.String,
		},
		"cust_tag": &graphql.Field{
			Type: graphql.String,
		},
		"cust_frg": &graphql.Field{
			Type: graphql.String,
		},
		"cust_status": &graphql.Field{
			Type: graphql.String,
		},
		"cust_first_visit_date": &graphql.Field{
			Type: graphql.String,
		},
		"cust_last_visit_date": &graphql.Field{
			Type: graphql.String,
		},
		"create_by": &graphql.Field{
			Type: graphql.String,
		},
		"create_date": &graphql.Field{
			Type: graphql.String,
		},
		"update_by": &graphql.Field{
			Type: graphql.String,
		},
		"update_date": &graphql.Field{
			Type: graphql.String,
		},
		"issue_date": &graphql.Field{
			Type: graphql.String,
		},
	}})
var CustomerMapPointType = graphql.NewObject(graphql.ObjectConfig{
	Name: "customerMappointMaster",
	Fields: graphql.Fields{
		"cust_id": &graphql.Field{
			Type: graphql.Int,
		},
		"mc_id": &graphql.Field{
			Type: graphql.Int,
		},
		"card_type": &graphql.Field{
			Type: graphql.String,
		},
		"point_amt": &graphql.Field{
			Type: graphql.Int,
		},
	}})
var PackageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "packageMaster",
	Fields: graphql.Fields{
		"package_id": &graphql.Field{
			Type: graphql.Int,
		},
		"package_name": &graphql.Field{
			Type: graphql.String,
		},
		"max_container_amt": &graphql.Field{
			Type: graphql.Int,
		},
		"create_date": &graphql.Field{
			Type: graphql.String,
		},
		"create_by": &graphql.Field{
			Type: graphql.String,
		},
		"update_date": &graphql.Field{
			Type: graphql.String,
		},
		"update_by": &graphql.Field{
			Type: graphql.String,
		},
	}})
var RequestPointType = graphql.NewObject(graphql.ObjectConfig{
	Name: "requestPointMaster",
	Fields: graphql.Fields{
		"request_point_id": &graphql.Field{
			Type: graphql.Int,
		},
		"mc_id": &graphql.Field{
			Type: graphql.Int,
		},
		"cust_id": &graphql.Field{
			Type: graphql.Int,
		},
		"container_id": &graphql.Field{
			Type: graphql.Int,
		},
		"request_status": &graphql.Field{
			Type: graphql.String,
		},
		"request_date": &graphql.Field{
			Type: graphql.String,
		},
		"reqeust_type": &graphql.Field{
			Type: graphql.String,
		},
		"request_message": &graphql.Field{
			Type: graphql.String,
		},
	}})
