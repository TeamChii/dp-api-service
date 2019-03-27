package service

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
)

func GetCRMReport(crmReportReq model.CRMReportReq) map[string]interface{} {
	var data model.CRMReportRes

	type CustomerSummary struct {
		Cust_all_count         int     `json:"cust_all_count"`
		Cust_new_count         int     `json:"cust_new_count"`
		Cust_repeating_count   int     `json:"cust_repeating_count"`
		Cust_new_percent       float64 `json:"cust_new_percent"`
		Cust_repeating_percent float64 `json:"cust_repeating_percent"`
	}

	type CRMSummary struct {
		CRMReportRes    model.CRMReportRes `json:"customer_report"`
		CustomerSummary CustomerSummary    `json:"customer_summary"`
	}

	var crm_result CRMSummary
	//check := config.DBsql().Where("mpos_menu_id = ?", id).Find(&data).RecordNotFound()
	check := false
	mc_id := strconv.Itoa(crmReportReq.MC_Id)
	from_date := crmReportReq.From_date
	to_date := crmReportReq.To_date
	from_date_array := strings.Split(from_date, "/")
	to_date_array := strings.Split(to_date, "/")

	/*
		data.All_Customer_Visit = 1000.0
		data.New_Customer_Visit = 41
		data.Repeating_Customer_Visit = 959.0
		data.Percent_Customer_Visit = "0"
		if data.All_Customer_Visit > 0 {
			percent := 	(data.Repeating_Customer_Visit*100.00)/(data.All_Customer_Visit)
			data.Percent_Customer_Visit = fmt.Sprintf("%.2f", percent)
		}
	*/
	// not clear
	//Member_Card_Reward float64    `json:"payment_received_credits"`
	//Customer_Point_Reward float64    `json:"average_sales"`
	//Point_Card_Reward int    `json:"total_bills_receipts"`

	start_time := from_date_array[0] + "-" + from_date_array[1] + "-" + from_date_array[2] + " 00:00:00"
	finish_time := to_date_array[0] + "-" + to_date_array[1] + "-" + to_date_array[2] + " 23:59:50"
	sql := `
	select
	 coalesce(cust_give_member.member_cust_count,0) give_member_cust_count
  	,coalesce(give_member.member_card_count,0) give_member_card_count
  	,coalesce(give_member.member_amt,0) give_member_amt

  	,coalesce(cust_give_point.point_cust_count,0) give_point_cust_count
  	,coalesce(give_point.point_card_count,0) give_point_card_count
  	,coalesce(give_point.point_amt,0) give_point_amt

  	,coalesce(cust_give_coupon.coupon_cust_count,0) give_coupon_cust_count
  	,coalesce(give_coupon.coupon_card_count,0) give_coupon_card_count
  	,coalesce(give_coupon.coupon_amt,0) give_coupon_amt

  	,coalesce(cust_redeem_point.point_cust_count,0) redeem_point_cust_count
  	,coalesce(redeem_point.point_card_count,0) redeem_point_card_count
  	,coalesce(redeem_point.point_amt,0) redeem_point_amt

  	,coalesce(cust_redeem_coupon.coupon_cust_count,0) redeem_coupon_cust_count
  	,coalesce(redeem_coupon.coupon_card_count,0) redeem_coupon_card_count
  	,coalesce(redeem_coupon.coupon_amt,0) redeem_coupon_amt
    from
	-- give member
		(
			select
				count(1) as member_card_count
				,sum(x.point_amt) as member_amt
					from (
						select tp.container_id
							,sum(tp.point_amt) as point_amt
						from dp_tb_point tp join
  							dp_ms_container mc on tp.container_id=mc.container_id
							where tp.mc_id = ` + mc_id + `
							and mc.container_type_id in (3)
							and
							tp.create_date >= '` + start_time + `'
							and
							tp.create_date <= '` + finish_time + `'
						group by tp.container_id
					) as x
		) give_member,
        (
			select
				count(1) as member_cust_count
					from (
						select tp.cust_id
						from dp_tb_point tp join
  							dp_ms_container mc on tp.container_id=mc.container_id
							where tp.mc_id = ` + mc_id + `
							and mc.container_type_id in (3)
							and
							tp.create_date >= '` + start_time + `'
							and
							tp.create_date <= '` + finish_time + `'
						group by tp.cust_id
					) as x
		) cust_give_member,
		-- give point
		(
			select
				count(1) as point_card_count
				,sum(x.point_amt) as point_amt
					from (
						select tp.container_id
							,sum(tp.point_amt) as point_amt
						from dp_tb_point tp join
  							dp_ms_container mc on tp.container_id=mc.container_id
							where tp.mc_id = ` + mc_id + `
							and mc.container_type_id in (1,4,5,6)
							and
							tp.create_date >= '` + start_time + `'
							and
							tp.create_date <= '` + finish_time + `'
						group by tp.container_id
					) as x
		) give_point,
        (
			select
				count(1) as point_cust_count
					from (
						select tp.cust_id
						from dp_tb_point tp join
  							dp_ms_container mc on tp.container_id=mc.container_id
							where tp.mc_id = ` + mc_id + `
							and mc.container_type_id in (1,4,5,6)
							and
							tp.create_date >= '` + start_time + `'
							and
							tp.create_date <= '` + finish_time + `'
						group by tp.cust_id
					) as x
		) cust_give_point,
		(
			select
				count(1) as coupon_card_count
				,sum(x.coupon_amt) as coupon_amt
					from (
						select tp.container_id
						,sum(tp.point_amt) as coupon_amt
						from dp_tb_point tp join
  							dp_ms_container mc on tp.container_id=mc.container_id
							where tp.mc_id = ` + mc_id + `
							and mc.container_type_id in (2)
							and
							tp.create_date >= '` + start_time + `'
							and
							tp.create_date <= '` + finish_time + `'
						group by tp.container_id
					) as x
		) give_coupon,
		(
			select
				count(1) as coupon_cust_count
					from (
						select tp.cust_id
						from dp_tb_point tp join
  							dp_ms_container mc on tp.container_id=mc.container_id
							where tp.mc_id = ` + mc_id + `
							and mc.container_type_id in (2)
							and
							tp.create_date >= '` + start_time + `'
							and
							tp.create_date <= '` + finish_time + `'
							group by tp.cust_id
					) as x
		) cust_give_coupon,
        -- redeem
		(
			select
				count(1) as point_card_count
				,sum(x.point_amt) as point_amt
					from (
						select tr.container_id
							,sum(tr.point_amt) as point_amt
						from dp_tb_redeem tr join
  							dp_ms_container mc on tr.container_id=mc.container_id
							where tr.mc_id = ` + mc_id + `
							and mc.container_type_id in (1,4,5,6)
							and
							tr.create_date >= '` + start_time + `'
							and
							tr.create_date <= '` + finish_time + `'
						group by tr.container_id
					) as x
		) redeem_point,
        (
			select
				count(1) as point_cust_count
					from (
						select tr.cust_id
						from dp_tb_redeem tr join
  							dp_ms_container mc on tr.container_id=mc.container_id
							where tr.mc_id = ` + mc_id + `
							and mc.container_type_id in (1,4,5,6)
							and
							tr.create_date >= '` + start_time + `'
							and
							tr.create_date <= '` + finish_time + `'
						group by tr.cust_id
					) as x
		) cust_redeem_point,
		(
			select
				count(1) as coupon_card_count
				,sum(x.coupon_amt) as coupon_amt
					from (
						select tr.container_id
							,sum(tr.point_amt) as coupon_amt
						from dp_tb_redeem tr join
  							dp_ms_container mc on tr.container_id=mc.container_id
							where tr.mc_id = ` + mc_id + `
							and mc.container_type_id in (2)
							and
							tr.create_date >= '` + start_time + `'
							and
							tr.create_date <= '` + finish_time + `'
						group by tr.container_id
					) as x
		) redeem_coupon,
		(
			select
				count(1) as coupon_cust_count
					from (
						select tr.cust_id
						from dp_tb_redeem tr join
  							dp_ms_container mc on tr.container_id=mc.container_id
							where tr.mc_id = ` + mc_id + `
							and mc.container_type_id in (2)
							and
							tr.create_date >= '` + start_time + `'
							and
							tr.create_date <= '` + finish_time + `'
						group by tr.cust_id
					) as x
		) cust_redeem_coupon
     `
	err := config.DBsql().Raw(sql).
		Scan(&data).Error
	if err != nil {
		println(err)
	}

	var customerSummary CustomerSummary
	sql = `
			select
				(cust_new_count.count+cust_repeating_count.count) as cust_all_count
				, cust_new_count.count cust_new_count
				,cust_repeating_count.count cust_repeating_count
				,case
					when (cust_new_count.count+cust_repeating_count.count)=0 then
						0
					else
						round(100.0 * (cust_new_count.count::decimal/(cust_new_count.count+cust_repeating_count.count)),1)
				end as cust_new_percent
				,case
					when (cust_new_count.count+cust_repeating_count.count)=0 then
						0
					else
						round(100.0 * (cust_repeating_count.count::decimal/(cust_new_count.count+cust_repeating_count.count)),1)
				end as cust_repeating_percent
				-- ,round((2::decimal/6),2) as test
			from (
				select count(*) as count
				from
					(select cust_id, count(cust_id)
					from dp_mp_merchant_customer
					where mc_id = ` + mc_id + `
					AND
						create_date >= '` + start_time + `'
					AND
						create_date <= '` + finish_time + `'
					group by cust_id
					having count(cust_id) = 1
					) as x
				) as cust_new_count,
				(select count(*) as count
				from (
					(select cust_id, count(cust_id)
					from dp_mp_merchant_customer
					where mc_id = ` + mc_id + `
					AND
						create_date >= '` + start_time + `'
					AND
						create_date <= '` + finish_time + `'
					group by cust_id
					having count(cust_id) > 1
					)
					) as y
				) as cust_repeating_count
		`

	err = config.DBsql().Raw(sql).
		Scan(&customerSummary).Error
	if err != nil {
		println(err)
	}

	crm_result.CRMReportRes = data
	crm_result.CustomerSummary = customerSummary
	var dataToJson map[string]interface{}
	b, _ := json.Marshal(crm_result)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
