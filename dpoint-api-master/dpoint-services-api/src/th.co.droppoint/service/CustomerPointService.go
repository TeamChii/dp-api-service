package service

import (
	"encoding/json"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func PointCheck(phone string) map[string]interface{} {
	type PointCustomer struct {
		Mc_id             int    `json:"mc_id"`
		Mc_name           string `json:"mc_name"`
		Container_Subject string `json:"container_subject"`
		Container_id      int    `json:"container_id"`
		Container_type_id int    `json:"container_type_id"`
		Current_point_amt int    `json:"current_point_amt"`
	}
	var pointCustomers []PointCustomer
	var err error
	err = config.DBsql().Raw(`select
	c2.container_id,mc.mc_id,mc.mc_name,
       c2.container_subject,c2.container_type_id,
       sum_cust.current_point_amt from dp_sum_customer_point sum_cust left join dp_ms_merchant mc
 on sum_cust.mc_id = mc.mc_id left join dp_ms_container c2 on sum_cust.container_id = c2.container_id
 left join dp_ms_customer customer on sum_cust.cust_id = customer.cust_id
where customer.cust_mobile=? and c2.container_type_id in (4,5,6)
order by  mc.mc_name asc , mc.mc_id , c2.container_type_id `, phone).
		Scan(&pointCustomers).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(pointCustomers)
	json.Unmarshal(b, &dataToJson)
	return utils.ResDataLoadNonPagin(err, dataToJson)

}
