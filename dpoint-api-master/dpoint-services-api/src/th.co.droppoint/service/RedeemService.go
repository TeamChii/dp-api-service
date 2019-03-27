package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func RedeemByIdMc(RedeemMaster model.RedeemReq) map[string]interface{} {

	var data []entity.RedeemEntity

	var count int
	err := config.DBsql().Where("mc_id = ?", RedeemMaster.Mc_id).Find(&data).Preload("MerchantEntity").Count(&count).Error

	fmt.Println(data)
	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoad(err, dataToJson, RedeemMaster.Paging, count)
}
func AddRedeem(RedeemReq model.RedeemReqList, ctx iris.Context) map[string]interface{} {

	CheckPointExpireAndResum(RedeemReq)

	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//	var resModel model.ResponseModel
	//id, _ := strconv.Atoi(utils.RandSeq(32))
	//RedeemMaster.Mc_id = id
	now := time.Now()
	createBy := utils.Decode(jwt.Raw).(string)
	var err2 error
	var err error
	redeemSucc := 0
	type LogRedeem struct {
		Container_Id int    `json:"container_id"`
		StatusCode   string `json:"statusCode"`
		MessageDesc  string `json:"messageDesc"`
	}
	var logRedeem []LogRedeem
	tx := config.DBsql().Begin()
	var mc_group entity.MerchantEntity

	config.DBsql().Where("mc_id = ? ", RedeemReq.Mc_id).Find(&mc_group)
	for index := 0; index < len(RedeemReq.Container); index++ {

		var sumcustomerPoint entity.SumCustomerPointEntity
		var check int
		config.DBsql().Where("mc_id = ? AND cust_id = ? AND container_id = ?", mc_group.Mc_group_id, RedeemReq.Cust_id, RedeemReq.Container[index].Container_id).
			Find(&sumcustomerPoint).Count(&check)

		var pointCurrent *int
		num := 0
		pointCurrent = &num
		var RedeemMaster entity.RedeemEntity
		var dataLog LogRedeem
		if check < 1 {
			dataLog.Container_Id = RedeemReq.Container[index].Container_id
			dataLog.StatusCode = "E"
			dataLog.MessageDesc = "Redeem Fail."
			logRedeem = append(logRedeem, dataLog)

		} else {
			pointCurrent = sumcustomerPoint.Current_point_amt
		}
		//if (RedeemReq.Container[index].Point_amt <= *pointCurrent && *pointCurrent > 0) && check > 0 {
		if ((RedeemReq.Container[index].Point_amt <= *pointCurrent && *pointCurrent > 0) || RedeemReq.Container[index].Container_Type_id == 3) && check > 0 {

			RedeemMaster.Mc_id = RedeemReq.Mc_id
			RedeemMaster.Cust_id = RedeemReq.Cust_id
			RedeemMaster.Container_id = RedeemReq.Container[index].Container_id
			RedeemMaster.Point_amt = RedeemReq.Container[index].Point_amt
			RedeemMaster.Create_date = &now
			RedeemMaster.Create_by = createBy
			//err := config.DBsql().Create(&RedeemMaster).Error
			err = tx.Create(&RedeemMaster).Error
			if err != nil {
				dataLog.Container_Id = RedeemReq.Container[index].Container_id
				dataLog.StatusCode = "E"
				dataLog.MessageDesc = "Redeem Fail."
				logRedeem = append(logRedeem, dataLog)
			}

			if err == nil {
				sumcustomerPoint.Cust_id = RedeemReq.Cust_id
				sumcustomerPoint.Mc_id = RedeemReq.Mc_id
				sumcustomerPoint.Container_id = RedeemMaster.Container_id
				sumcustomerPoint.Current_point_amt = &RedeemMaster.Point_amt

				pointCur := 1
				if RedeemReq.Container[index].Container_Type_id != 3 {
					pointCur = *pointCurrent - RedeemMaster.Point_amt
				}

				sumcustomerPoint.Current_point_amt = &pointCur
				err2 = tx.Model(&sumcustomerPoint).
					Where("mc_id = ? AND cust_id = ? AND container_id = ?", RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemMaster.Container_id).
					Update(&sumcustomerPoint).Error

				if err2 != nil {
					dataLog.Container_Id = RedeemReq.Container[index].Container_id
					dataLog.StatusCode = "E"
					dataLog.MessageDesc = "Redeem Fail."
					logRedeem = append(logRedeem, dataLog)
				} else {

					// -------------------- Do UpdatePointRedeemFlag in dp_tb_point -----------------------
					type PointObject struct {
						Point_id    int        `json:"point_id"`
						Point_amt   int        `json:"point_amt"`
						Expire_date *time.Time `json:"expire_date"`
						Redeem_amt  int        `json:"redeem_amt"`
						Expire_amt  int        `json:"expire_amt"`
						Sum         int        `json:"sum"`
					}
					var PointList []PointObject

					config.DBsql().
						Raw("SELECT * FROM ( SELECT p.point_id, p.point_amt, p.expire_date, p.redeem_amt, ( SELECT sum(point_amt) FROM dp_tb_point WHERE (point_amt > p.point_amt OR (expire_date < p.expire_date AND point_amt = p.point_amt) OR (expire_date = p.expire_date) AND (expire_date = p.expire_date AND point_id <= p.point_id)) AND mc_id = ? AND cust_id = ? AND container_id = ? AND expire_flag != 'Y' AND redeem_flag != 'Y' ORDER BY p.expire_date ASC ) AS sum FROM dp_tb_point p WHERE mc_id = ? AND cust_id = ? AND container_id = ? AND expire_flag != 'Y' AND redeem_flag != 'Y' ORDER BY p.expire_date ASC ) tb WHERE tb.sum - tb.point_amt < ?",
							RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[index].Container_id, RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[index].Container_id, RedeemReq.Container[index].Point_amt).
						Scan(&PointList)

					var pointRedeem = RedeemReq.Container[index].Point_amt

					updatePointFlagSucc := 0
					for i := 0; i < len(PointList); i++ {

						var PointMaster entity.PointEntity
						var availablePoint = PointList[i].Point_amt - PointList[i].Redeem_amt - PointList[i].Expire_amt

						if pointRedeem-availablePoint >= 0 {
							PointMaster.Redeem_flag = "Y"
							now := time.Now()
							PointMaster.Redeem_date = &now
							PointMaster.Redeem_amt = PointList[i].Redeem_amt + availablePoint

							pointRedeem = pointRedeem - availablePoint
						} else {
							PointMaster.Redeem_flag = "P"
							now := time.Now()
							PointMaster.Redeem_date = &now
							PointMaster.Redeem_amt = PointList[i].Redeem_amt + pointRedeem

							pointRedeem = 0
						}

						PointMaster.Point_id = PointList[i].Point_id

						errPoint := tx.Model(&PointMaster).Where("point_id = ?", PointList[i].Point_id).Update(&PointMaster).Error
						if errPoint != nil {
							dataLog.Container_Id = RedeemReq.Container[index].Container_id
							dataLog.StatusCode = "E"
							dataLog.MessageDesc = "Update point flag fail."
							logRedeem = append(logRedeem, dataLog)
						} else {
							updatePointFlagSucc += 1
						}
					}
					// -------------------- End do UpdatePointRedeemFlag in dp_tb_point -----------------------

					if updatePointFlagSucc == len(PointList) {
						redeemSucc++
					}
				}

				/*if check > 0{
					//fmt.Println("update sum")
					pointCur := 1
					if RedeemReq.Container[index].Container_Type_id!=3 {
						 pointCur = *pointCurrent - RedeemMaster.Point_amt
					}

					sumcustomerPoint.Current_point_amt = &pointCur
					err2 = tx.Model(&sumcustomerPoint).
						Where("mc_id = ? AND cust_id = ? AND container_id = ?", RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemMaster.Container_id).
						Update(&sumcustomerPoint).Error

					if err2 != nil {
						dataLog.Container_Id = RedeemReq.Container[index].Container_id
						dataLog.StatusCode = "E"
						dataLog.MessageDesc = "Redeem Fail."
						logRedeem = append(logRedeem, dataLog)
					} else {
						redeemSucc++
					}

				}else   {
					//fmt.Println("create sum")
					err2 = tx.Create(&sumcustomerPoint).Error

					if err2 != nil {
						dataLog.Container_Id = RedeemReq.Container[index].Container_id
						dataLog.StatusCode = "E"
						dataLog.MessageDesc = "Redeem Fail."
						logRedeem = append(logRedeem, dataLog)
					} else {
						redeemSucc++
					}
				}*/

			}

		} else {
			if check > 0 {
				dataLog.Container_Id = RedeemReq.Container[index].Container_id
				dataLog.StatusCode = "E"
				dataLog.MessageDesc = "Your point is not enough."
				logRedeem = append(logRedeem, dataLog)

			}
		}
		/*if tx.Commit().Error != nil {

		}*/
	}

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(logRedeem)
	json.Unmarshal(b, &dataToJson)

	if redeemSucc == len(RedeemReq.Container) {
		tx.Commit()
		return map[string]interface{}{
			"statusCode":  "S",
			"messageCode": "Success",
			"messageDesc": "Redeem Success",
			"log":         dataToJson,
		}
	} else {
		tx.Rollback()
		return map[string]interface{}{
			"statusCode":  "E",
			"messageCode": "Error",
			"messageDesc": "Redeem Fail",
			"log":         dataToJson,
		}
	}
}

func UpdateRedeem(RedeemMaster entity.RedeemEntity) map[string]string {

	//var resModel model.ResponseModel
	//account.UpdatedDate = time.Now()
	err := config.DBsql().Model(&RedeemMaster).Where("redeem_id = ?", RedeemMaster.Redeem_id).Update(&RedeemMaster).Error

	return utils.ResDataEdit(err)
}

func CheckPointExpireAndResum(RedeemReq model.RedeemReqList) bool {
	var checkSucc int
	var expireFlag = "Y"
	var redeemFlag = "Y"
	now := time.Now()
	tx := config.DBsql().Begin()
	for index := 0; index < len(RedeemReq.Container); index++ {
		var PointMaster []entity.PointEntity
		var count int
		config.DBsql().
			Where("mc_id = ? AND cust_id = ? AND container_id = ? AND expire_flag!=? AND redeem_flag!=?",
				RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[index].Container_id, expireFlag, redeemFlag).
			Find(&PointMaster).Count(&count)

		var sumCurrentAmount = 0
		var sumExpireAmount = 0
		var containerLevelStatus = true

		if count > 0 {
			for i := 0; i < len(PointMaster); i++ {
				var success = true
				if PointMaster[i].Expire_date != nil {
					if now.Sub(PointMaster[i].Expire_date.UTC()) > 0 {
						PointMaster[i].Expire_flag = "Y"
						PointMaster[i].Expire_amt = PointMaster[i].Point_amt - PointMaster[i].Redeem_amt

						err := tx.Model(&PointMaster[i]).
							Where("mc_id = ? AND cust_id = ? AND container_id = ?", RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[index].Container_id).
							Update(&PointMaster[i]).Error

						if err != nil {
							success = false
						} else {
							success = true
						}
					}
				}
				if success {
					sumCurrentAmount += PointMaster[i].Point_amt - PointMaster[i].Redeem_amt - PointMaster[i].Expire_amt
					sumExpireAmount += PointMaster[i].Expire_amt
				}
			}

			var sumCustomerPoint entity.SumCustomerPointEntity

			config.DBsql().Where("mc_id = ? AND cust_id = ? AND container_id = ?", RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[index].Container_id).Find(&sumCustomerPoint)

			if sumCurrentAmount != 0 || sumExpireAmount != 0 {
				sumCustomerPoint.Current_point_amt = &sumCurrentAmount
				sumCustomerPoint.Expired_point_amt = sumCustomerPoint.Expired_point_amt + sumExpireAmount
			}

			sumError := tx.Model(&sumCustomerPoint).Where("mc_id = ? AND cust_id = ? AND container_id = ?", RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[index].Container_id).Update(&sumCustomerPoint).Error
			if sumError != nil {
				containerLevelStatus = false
			}
		}

		if containerLevelStatus {
			checkSucc += 1
		}
	}

	if checkSucc == len(RedeemReq.Container) {
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		return false
	}
}

func UpdatePointRedeemFlag(RedeemReq model.RedeemReqList) bool {

	type PointObject struct {
		Point_id    int        `json:"point_id"`
		Point_amt   int        `json:"point_amt"`
		Expire_date *time.Time `json:"expire_date"`
		Redeem_amt  int        `json:"redeem_amt"`
		Expire_amt  int        `json:"expire_amt"`
		Sum         int        `json:"sum"`
	}
	var PointList []PointObject

	config.DBsql().
		Raw("SELECT * FROM ( SELECT p.point_id, p.point_amt, p.expire_date, p.redeem_amt, ( SELECT sum(point_amt) FROM dp_tb_point WHERE (point_amt > p.point_amt OR (expire_date < p.expire_date AND point_amt = p.point_amt) OR (expire_date = p.expire_date) AND (expire_date = p.expire_date AND point_id <= p.point_id)) AND mc_id = ? AND cust_id = ? AND container_id = ? AND expire_flag != 'Y' AND redeem_flag != 'Y' ORDER BY p.expire_date ASC ) AS sum FROM dp_tb_point p WHERE mc_id = ? AND cust_id = ? AND container_id = ? AND expire_flag != 'Y' AND redeem_flag != 'Y' ORDER BY p.expire_date ASC ) tb WHERE tb.sum - tb.point_amt < ?",
			RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[0].Container_id, RedeemReq.Mc_id, RedeemReq.Cust_id, RedeemReq.Container[0].Container_id, RedeemReq.Container[0].Point_amt).
		Scan(&PointList)

	var pointRedeem = RedeemReq.Container[0].Point_amt

	for i := 0; i < len(PointList); i++ {

		var PointMaster entity.PointEntity
		var availablePoint = PointList[i].Point_amt - PointList[i].Redeem_amt - PointList[i].Expire_amt

		if pointRedeem-availablePoint >= 0 {
			PointMaster.Redeem_flag = "Y"
			now := time.Now()
			PointMaster.Redeem_date = &now
			PointMaster.Redeem_amt = PointList[i].Redeem_amt + availablePoint

			pointRedeem = pointRedeem - availablePoint
		} else {
			PointMaster.Redeem_flag = "P"
			now := time.Now()
			PointMaster.Redeem_date = &now
			PointMaster.Redeem_amt = PointList[i].Redeem_amt + pointRedeem

			pointRedeem = 0
		}

		PointMaster.Point_id = PointList[i].Point_id

		errPoint := config.DBsql().Model(&PointMaster).Where("point_id = ?", PointList[i].Point_id).Update(&PointMaster).Error
		if errPoint != nil {
			println("no")
		} else {
			println("yes")
		}
	}

	return true
}
