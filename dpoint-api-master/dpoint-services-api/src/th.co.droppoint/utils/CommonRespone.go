package utils

import (
	"fmt"
	"math"
	"strconv"

	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/model"
		"th.co.droppoint/utils"
		"th.co.droppoint/service/authentication"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/model"
	"github.com/go-sql-driver/mysql"
)

func ResponseError() map[string]interface{} {
	return map[string]interface{}{
		"statusCode":  STATUS_CODE_ERROR,
		"messageCode": "400",
		"messageAbbr": MESSAGE_ABBR_ERROR,
		"messageDesc": "Post Data Err"}
}

func ResDataLoad(err error, dataToJson []map[string]interface{}, paging model.PagingModel, totalRecord int) map[string]interface{} {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		case 1054:
			msg = string(err.(*mysql.MySQLError).Message)
			code = "E003"
		default:
			msg = string(err.(*mysql.MySQLError).Message)
			code = "E000"
		}
		return map[string]interface{}{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}
	} else {
		totalPage := int(math.Ceil(float64(paging.TotalRecord) / float64(paging.PageSize)))
		return map[string]interface{}{
			"data":        dataToJson,
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_LOAD_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_LOAD_SUCCESS,
			"messageDesc": MESSAGE_DESC_LOAD_SUCCESS,
			"paging": map[string]interface{}{
				"pageNo":      paging.PageNo,
				"pageSize":    paging.PageSize,
				"totalRecord": totalRecord,
				"totalPage":   totalPage,
			},
		}
	}
}
func ResDataLoadNonPagin(err error, dataToJson []map[string]interface{}) map[string]interface{} {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		case 1054:
			msg = string(err.(*mysql.MySQLError).Message)
			code = "E003"
		default:
			msg = string(err.(*mysql.MySQLError).Message)
			code = "E000"
		}
		return map[string]interface{}{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}
	} else {
		return map[string]interface{}{
			"data":        dataToJson,
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_LOAD_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_LOAD_SUCCESS,
			"messageDesc": MESSAGE_DESC_LOAD_SUCCESS,
		}
	}
}
func ResDataLoadNonPaginNonArray(err error, dataToJson map[string]interface{}) map[string]interface{} {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		case 1054:
			msg = string(err.(*mysql.MySQLError).Message)
			code = "E003"
		default:
			msg = string(err.(*mysql.MySQLError).Message)
			code = "E000"
		}
		return map[string]interface{}{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}
	} else {
		return map[string]interface{}{
			"data":        dataToJson,
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_LOAD_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_LOAD_SUCCESS,
			"messageDesc": MESSAGE_DESC_LOAD_SUCCESS,
		}
	}
}
func ResDataAdd(err error) map[string]string {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		default:
			msg = string(err.Error())
			code = "E000"
		}
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}
	} else {
		return map[string]string{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_ADD_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_ADD_SUCCESS,
			"messageDesc": MESSAGE_DESC_ADD_SUCCESS}
	}
}
func ResDataAdd2(err error, cust_id int) map[string]interface{} {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		default:
			msg = string(err.Error())
			code = "E000"
		}
		return map[string]interface{}{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg,
			"cust_id":     nil}
	} else {
		return map[string]interface{}{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_ADD_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_ADD_SUCCESS,
			"messageDesc": MESSAGE_DESC_ADD_SUCCESS,
			"cust_id":     cust_id}
	}
}
func ResDataAddWithId(err error, id string, id2 string) map[string]string {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		default:
			msg = string(err.Error())
			code = "E000"
		}
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg,
			"refId":       id,
		}
	} else {
		return map[string]string{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_ADD_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_ADD_SUCCESS,
			"messageDesc": MESSAGE_DESC_ADD_SUCCESS,
			"refId":       id,
			"refPath":     id2,
		}
	}
}
func ResDataEdit(err error) map[string]string {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		default:
			msg = string(err.Error())
			code = "E000"
		}
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}
	} else {
		return map[string]string{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_UPDATE_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_UPDATE_SUCCESS,
			"messageDesc": MESSAGE_DESC_UPDATE_SUCCESS}
	}
}
func ResDataEditWithId(err error, id string, id2 string) map[string]string {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		default:
			msg = string(err.Error())
			code = "E000"
		}
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}
	} else {
		return map[string]string{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_UPDATE_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_UPDATE_SUCCESS,
			"messageDesc": MESSAGE_DESC_UPDATE_SUCCESS,
			"refId":       id,
			"refPath":     id2,
		}
	}
}

func ResDataDel(dataAll int, deleted int) map[string]string {

	if dataAll != deleted && deleted != 0 {
		return map[string]string{
			"statusCode":  "W",
			"messageCode": "W001",
			"messageAbbr": "Warning",
			"messageDesc": "data " + strconv.Itoa(dataAll) + " deleted " + strconv.Itoa(deleted)}
	} else if dataAll != deleted && deleted == 0 {
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": "E005",
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": "data " + strconv.Itoa(dataAll) + " deleted " + strconv.Itoa(deleted)}

	} else {
		return map[string]string{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_DELETE_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_DELETE_SUCCESS,
			"messageDesc": MESSAGE_DESC_DELETE_SUCCESS}
	}
}

func ResDataLoadById(check bool, dataToJson map[string]interface{}) map[string]interface{} {

	if !check {
		return map[string]interface{}{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": "E004",
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": "data not found"}
	} else {
		return map[string]interface{}{
			"data":        []map[string]interface{}{dataToJson},
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_GET_BY_ID_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_GET_BY_ID_SUCCESS,
			"messageDesc": MESSAGE_DESC_GET_BY_ID_SUCCESS,
		}
	}
}
func ResDataAddMerchant(err error, data []map[string]interface{}, tokenString string) map[string]interface{} {
	msg := ""
	code := ""
	fmt.Println(err)
	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		default:
			msg = string(err.Error())
			code = "E000"
		}
		return map[string]interface{}{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}

	} else {

		return map[string]interface{}{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_ADD_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_ADD_SUCCESS,
			"messageDesc": MESSAGE_DESC_ADD_SUCCESS,
			"data":        data,
			"token":       tokenString,
		}
	}
}

func ResDataMapping(dataAll int, deleted int) map[string]string {

	if dataAll != deleted && deleted != 0 {
		return map[string]string{
			"statusCode":  "W",
			"messageCode": "W001",
			"messageAbbr": "Warning",
			"messageDesc": "data " + strconv.Itoa(dataAll) + " mapped " + strconv.Itoa(deleted)}
	} else if dataAll != deleted && deleted == 0 {
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": "E005",
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": "data " + strconv.Itoa(dataAll) + " mapped " + strconv.Itoa(deleted)}

	} else {
		return map[string]string{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_MAPPING_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_MAPPING_SUCCESS,
			"messageDesc": MESSAGE_DESC_MAPPING_SUCCESS}
	}
}

func ResDataRemvoveMapping(dataAll int, deleted int) map[string]string {

	if dataAll != deleted && deleted != 0 {
		return map[string]string{
			"statusCode":  "W",
			"messageCode": "W001",
			"messageAbbr": "Warning",
			"messageDesc": "data " + strconv.Itoa(dataAll) + " removed " + strconv.Itoa(deleted)}
	} else if dataAll != deleted && deleted == 0 {
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": "E005",
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": "data " + strconv.Itoa(dataAll) + " removed " + strconv.Itoa(deleted)}

	} else {
		return map[string]string{
			"statusCode":  STATUS_CODE_SUCCESS,
			"messageCode": MESSAGE_CODE_REMOVE_MAPPING_SUCCESS,
			"messageAbbr": MESSAGE_ABBR_REMOVE_MAPPING_SUCCESS,
			"messageDesc": MESSAGE_DESC_REMVOE_MAPPING_SUCCESS}
	}
}

func ResDataReceipt(err error, id string, member_point string) map[string]string {
	msg := ""
	code := ""

	if err != nil {
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			msg = "ข้อมูลซ้ำ"
			code = "E001"
		case 1452:
			msg = "ไม่พบ Foreign Key"
			code = "E002"
		default:
			msg = string(err.Error())
			code = "E000"
		}
		return map[string]string{
			"statusCode":  STATUS_CODE_ERROR,
			"messageCode": code,
			"messageAbbr": MESSAGE_ABBR_ERROR,
			"messageDesc": msg}
	} else {
		return map[string]string{
			"statusCode":   STATUS_CODE_SUCCESS,
			"messageCode":  MESSAGE_CODE_UPDATE_SUCCESS,
			"messageAbbr":  MESSAGE_ABBR_UPDATE_SUCCESS,
			"messageDesc":  MESSAGE_DESC_UPDATE_SUCCESS,
			"receipt_id":   id,
			"member_point": member_point,
		}
	}
}
