package service

import (
	"encoding/json"
	"fmt"
	"time"

	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
		"th.co.droppoint/service/authentication"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func LoadDashboard(ReportReq model.ReportReq) map[string]interface{} {

	var data []model.ReportEnityModel
	var err error
	if ReportReq.Type == "day" {
		err = config.DBsql().
			Table("dp_rp_report").
			Select("mc_id, report_category, count(customer_id), max(unit) as unit").
			Where(`exists(
				select * from dp_ms_system_param as sys_param
				where sys_param.category_name='DASHBOARD' and sys_param.active='1'
				and sys_param.key_code=dp_rp_report.report_category ) AND mc_id = ?
				and date_date BETWEEN now()::date AND (now()::date + '1 days'::interval)`, ReportReq.Mc_id).
			Group("mc_id,report_category ").
			Find(&data).Error
	} else if ReportReq.Type == "week" {
		err = config.DBsql().
			Table("dp_rp_report").
			Select("mc_id, report_category, count(customer_id), max(unit) as unit").
			Where(`exists(
				select * from dp_ms_system_param as sys_param
				where sys_param.category_name='DASHBOARD' and sys_param.active='1'
				and sys_param.key_code=dp_rp_report.report_category ) AND mc_id = ?
				and date_date BETWEEN date_trunc('week', now()::date) AND (date_trunc('week', now()::date) + '7 days'::interval)`, ReportReq.Mc_id).
			Group("mc_id,report_category ").
			Find(&data).Error
	} else {
		err = config.DBsql().
			Table("dp_rp_report").
			Select("mc_id, report_category, count(customer_id), max(unit) as unit").
			Where(`exists(
				select * from dp_ms_system_param as sys_param
				where sys_param.category_name='DASHBOARD' and sys_param.active='1'
				and sys_param.key_code=dp_rp_report.report_category ) AND mc_id = ?
				and date_date BETWEEN date_trunc('month', now()::date) AND date_trunc('month', now()::date)+'1 month'::interval`, ReportReq.Mc_id).
			Group("mc_id,report_category ").
			Find(&data).Error
	}

	var dataResp model.ReportResp
	for index := 0; index < len(data); index++ {
		d := data[index]
		if data[index].Report_category == "DASHBOARD_NEW_CUSTOMER" {
			dataResp.New_customers = &d
		} else if data[index].Report_category == "DASHBOARD_POINT_TO_CUSTOMER" {
			dataResp.Point_to_customer = &d
		} else if data[index].Report_category == "DASHBOARD_REDEEMED_CUSTOMER" {
			dataResp.Redeemed_cutomers = &d
		} else if data[index].Report_category == "DASHBOARD_GIVE_POINT" {
			dataResp.Point_giving = &d
		} else if data[index].Report_category == "DASHBOARD_SPECIAL_REQUEST" {
			dataResp.Special_request = &d
		} else {
			dataResp.Special_request = &d
		}
	}

	var dataToJson map[string]interface{}
	b, _ := json.Marshal(dataResp)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPaginNonArray(err, dataToJson)
}

func LoadReportList() map[string]interface{} {
	var data []model.ReportEnityModel2
	err := config.DBsql().
		Table("dp_rp_report").
		Select("mc_id,report_category , max(sys_param_out.key_value) as name").
		Joins(`left join dp_ms_system_param sys_param_out
					on dp_rp_report.report_category = sys_param_out.key_code`).
		Where(`exists(
			select * from dp_ms_system_param sys_param
			where sys_param.category_name='REPORT_HISTORY' and sys_param.active='1'
			and sys_param.key_code=dp_rp_report.report_category )`).
		Group("mc_id,report_category").
		Find(&data).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func LoadReportByCategory(ReportDateReq model.ReportDateReq) map[string]interface{} {

	var data []entity.ReportEntity

	//fmt.Println(ReportDateReq.Start_month)
	//fmt.Println(ReportDateReq.End_month)

	timeFormat := "01-2006"

	start, _ := time.Parse(timeFormat, ReportDateReq.Start_month)

	end, _ := time.Parse(timeFormat, ReportDateReq.End_month)

	if ReportDateReq.Start_month == ReportDateReq.End_month {
		end = end.AddDate(0, 1, 0)
	}

	//fmt.Println(start.Format("2006-01-02 15:04:05"))
	//fmt.Println(end.Format("2006-01-02 15:04:05"))

	//startStr := start.Format("2006-01-02 15:04:05")
	//endStr := end.Format("2006-01-02 15:04:05")
	fmt.Println(start)
	fmt.Println(end)
	err := config.DBsql().
		Where(`dp_rp_report.mc_id = ? and
		dp_rp_report.report_category = ?
			  AND date_date BETWEEN ? AND ? `, ReportDateReq.Mc_id, ReportDateReq.Report_category, start, end).
		Order("dp_rp_report.date_date desc").
		Limit(50).
		Find(&data).Error

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}
func LoadCampaignMonitor(ReportReq model.ReportReq) map[string]interface{} {

	type ResultMain struct {
		Mc_id          int    `json:"mc_id"`
		Campaign_id    int    `json:"campaign_id"`
		Campaign_code  string `json:"campaign_code"`
		Campaign_name  string `json:"campaign_name"`
		Campagin_color string `json:"campagin_color"`
		Total          int    `json:"total"`
		Sended         int    `json:"sended"`
		Responsed      int    `json:"responsed"`
	}
	type ResultResp struct {
		Who_view      ResultMain `json:"who_view"`
		Top_customers ResultMain `json:"top_customers"`
		New_customers ResultMain `json:"new_customers"`
	}
	var data []ResultMain
	//var resp ResultResp
	var err error

	if ReportReq.Type == "this" {
		err = config.DBsql().Raw(`select
		x.mc_id,
		  x.campaign_id,
		  x.campaign_code,
			x.campagin_color,
			x.campaign_name,
		  (select count(*) from dp_mp_merchant_customer_campaign mcc_inner
		   where mcc_inner.mc_id=x.mc_id and mcc_inner.campaign_id=x.campaign_id
		  ) as total,
		  (select count(*) from dp_mp_merchant_customer_campaign mcc_inner
		   where mcc_inner.mc_id=x.mc_id and mcc_inner.campaign_id=x.campaign_id
			 and mcc_inner.response_status='Y'
		  ) as responsed,
			(select count(*) from dp_mp_merchant_customer_campaign mcc_inner
		   where mcc_inner.mc_id=x.mc_id and mcc_inner.campaign_id=x.campaign_id
			 and mcc_inner.send_status='Y'
		  ) as sended
		from (
		  select
			mcc.mc_id,
			mcc.campaign_id,
			mcc.cust_id,
			mcc.response_status,
			mc.campaign_code,
			mc.campaign_name,
				mc.campagin_color,
				mcc.send_date
		  from dp_mp_merchant_customer_campaign mcc
			left join dp_ms_campaign mc on mcc.campaign_id = mc.campaign_id
		)x 
		where x.send_date BETWEEN date_trunc('month', now()::date) AND date_trunc('month', now()::date)+'1 month'::interval
		and x.mc_id = ?
		group by x.mc_id,x.campaign_id,x.campaign_code,x.campagin_color,x.campaign_name;`, ReportReq.Mc_id).
			Scan(&data).Error

	} else {

		err = config.DBsql().Raw(`select
		x.mc_id,
		  x.campaign_id,
		  x.campaign_code,
			x.campagin_color,
			x.campaign_name,
		  (select count(*) from dp_mp_merchant_customer_campaign mcc_inner
		   where mcc_inner.mc_id=x.mc_id and mcc_inner.campaign_id=x.campaign_id
		  ) as total,
		  (select count(*) from dp_mp_merchant_customer_campaign mcc_inner
		   where mcc_inner.mc_id=x.mc_id and mcc_inner.campaign_id=x.campaign_id
			 and mcc_inner.response_status='Y'
		  ) as responsed,
			(select count(*) from dp_mp_merchant_customer_campaign mcc_inner
		   where mcc_inner.mc_id=x.mc_id and mcc_inner.campaign_id=x.campaign_id
			 and mcc_inner.send_status='Y'
		  ) as sended
		from (
		  select
			mcc.mc_id,
			mcc.campaign_id,
			mcc.cust_id,
			mcc.response_status,
			mc.campaign_code,
			mc.campaign_name,
				mc.campagin_color,
				mcc.send_date
		  from dp_mp_merchant_customer_campaign mcc
			left join dp_ms_campaign mc on mcc.campaign_id = mc.campaign_id
		)x 
		where x.send_date BETWEEN date_trunc('month', now()::date)-'1 month'::interval AND date_trunc('month', now()::date)
		and x.mc_id = ?
		group by x.mc_id,x.campaign_id,x.campaign_code,x.campagin_color,x.campaign_name;`, ReportReq.Mc_id).
			Scan(&data).Error

	}

	/*for index := 0; index < len(data); index++ {
		if data[index].Campaign_code == "WHO_VIEW" {
			resp.Who_view = data[index]
		} else if data[index].Campaign_code == "TOP_CUSTOMER" {
			resp.Top_customers = data[index]
		} else if data[index].Campaign_code == "NEW_CUSTOMER" {
			resp.New_customers = data[index]
		}
	}*/

	var dataToJson []map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadNonPagin(err, dataToJson)
}

/*func AddUser(UserMaster entity.UserEntity) map[string]string {

	err := config.DBsql().Create(&UserMaster).Error

	return utils.ResDataAdd(err)

}
func UpdateUser(UserMaster entity.UserEntity) map[string]string {

	err := config.DBsql().Model(&UserMaster).Where("user_id = ?", UserMaster.User_id).Update(&UserMaster).Error

	return utils.ResDataEdit(err)
}
func UpdateUserPin(PinReq model.PinReq) map[string]string {

	var UserMaster entity.UserEntity
	UserMaster.Pin = utils.HashPin(PinReq.Pin)
	err := config.DBsql().Model(&UserMaster).Where("user_id = ?", PinReq.User_id).Update("pin", utils.HashPin(PinReq.Pin)).Error

	return utils.ResDataEdit(err)
}*/
