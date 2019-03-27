package service

import (
	"log"
	_ "reflect"
	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func CheckIncreasPoint(obj model.LoginToOpenReq) map[string]interface{} {
	var dataToJson map[string]interface{}

	var cust_point_group entity.SumCustomerPointEntity
	current_time := time.Now()
	// log.Println(obj.Cust_id)
	// log.Println(obj.Mc_id)
	// log.Println(obj.Container_id)

	check := config.DBsql().
		Where("mc_id = ? And container_id = ? And cust_id = ?", obj.Mc_id, obj.Container_id, obj.Cust_id).
		Find(&cust_point_group).RecordNotFound()

	if !check {
		config.DBsql().
			Select("current_point_amt,last_archived_point_date").
			Where("mc_id = ? And container_id = ? And cust_id = ?", obj.Mc_id, obj.Container_id, obj.Cust_id).
			Find(&cust_point_group)

		diff_time := current_time.Sub(*cust_point_group.Last_archived_point_date)
		log.Println(diff_time)

		if diff_time.Hours() > float64(3.0) {
			log.Println("update")

			config.DBsql().Model(&cust_point_group).
				Where("mc_id = ? And container_id = ? And cust_id = ?", obj.Mc_id, obj.Container_id, obj.Cust_id).
				Update(
					map[string]interface{}{"current_point_amt": *cust_point_group.Current_point_amt + 1, "last_archived_point_date": time.Now()})
				// b, _ := json.Marshal()
				// json.Unmarshal(b, &dataToJson)
			return utils.ResDataLoadById(!check, dataToJson)

		} else {
			log.Println("no update")
			// b, _ := json.Marshal(d)
			// json.Unmarshal(b, &dataToJson)
			return utils.ResDataLoadById(check, dataToJson)

		}
	}
	return utils.ResDataLoadById(!check, dataToJson)
}
