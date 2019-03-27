package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const CHARSET_RANDOM_WITH_STRING = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const CHARSET_RANDOM = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func RandSeq(n int) string {
	return stringWithCharset(n, CHARSET_RANDOM)
}
func RandString(n int) string {
	return stringWithCharset(n, CHARSET_RANDOM_WITH_STRING)
}
func StringToDate(str string) time.Time {
	timeFormat := "02/01/2006"
	create_time, _ := time.Parse(timeFormat, str)
	return create_time
}
func GenIssueExpireDate(date string, mode string) map[string]interface{} {
	//now := time.Now()
	var now = time.Now()

	var issue_date *time.Time
	var expire_date *time.Time

	if mode == "DD" {
		dateStr, _ := strconv.Atoi(date)
		issue_date = &now
		date := (now.AddDate(0, 0, dateStr))
		expire_date = &date
	} else if mode == "MM" {
		dateStr, _ := strconv.Atoi(date)
		issue_date = &now
		date := now.AddDate(0, dateStr, 0)
		expire_date = &date
	} else if mode == "YY" {
		dateStr, _ := strconv.Atoi(date)
		now := time.Now()
		issue_date = &now
		date := now.AddDate(dateStr, 0, 0)
		expire_date = &date
	} else if mode == "EM" {
		/*currentYear, currentMonth, _ := now.Date()
		currentLocation := now.Location()

		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		issue_date = &now
		expire_date = &lastOfMonth*/

		dateStr, _ := strconv.Atoi(date)
		issue_date = &now
		date := now.AddDate(0, dateStr, 0)
		expire_date = &date
	} else if mode == "EY" {
		dateStr, _ := strconv.Atoi(date)
		issue_date = &now
		date := now.AddDate(dateStr, 0, 0)
		expire_date = &date
	} else if mode == "FX" {
		timeFormat := "02/01/2006"
		create_time, _ := time.Parse(timeFormat, date)

		issue_date = &now
		expire_date = &create_time
	}

	return map[string]interface{}{
		"issue_date":  issue_date,
		"expire_date": expire_date,
	}
}
func CommonSendOTPCode(mobile string) string {
	code := RandSeq(6)
	/* */
	resp, err := http.PostForm("http://www.thaibulksms.com/sms_api.php",
		url.Values{
			"username": {"0818904560"},
			"password": {"029262"},
			"msisdn":   {mobile},
			"message":  {"SMS OTP for transaction approval on DropPoint \r<OTP: " + code + "> will be expried in 5 minutes"},
			"sender":   {"NOTICE"},
			"force":    {"standard"},
		})
	/* */
	/*
		    <SMS>
		    <QUEUE>
		        <Msisdn>0966408458</Msisdn>
		        <Status>1</Status>
		        <Transaction>23a7e0c3bef949be927f01b70852727c</Transaction>
		        <UsedCredit>1</UsedCredit>
		        <RemainCredit>1532</RemainCredit>
		    </QUEUE>
		    <!-- 0.015279054641724 -->
			</SMS>
	*/
	/* */
	if nil != err {
		fmt.Println("errorination happened getting the response", err)
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		fmt.Println("errorination happened reading the body", err)
	}
	/* */
	return code
}

/*
func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
*/
func CheckContentPath(root_path string) string {
	t := time.Now()
	year_dir := t.Format("2006")
	month_dir := t.Format("01")
	day_dir := t.Format("02")
	if _, err := os.Stat(root_path + year_dir); os.IsNotExist(err) {
		os.MkdirAll(root_path+year_dir, os.ModePerm)
	}
	if _, err := os.Stat(root_path + year_dir + "/" + month_dir); os.IsNotExist(err) {
		os.MkdirAll(root_path+year_dir+"/"+month_dir, os.ModePerm)
	}
	if _, err := os.Stat(root_path + year_dir + "/" + month_dir + "/" + day_dir); os.IsNotExist(err) {
		os.MkdirAll(root_path+year_dir+"/"+month_dir+"/"+day_dir, os.ModePerm)
	}
	//current_path := root_path+year_dir+"/"+month_dir+"/"+day_dir+"/";
	current_path := year_dir + "/" + month_dir + "/" + day_dir + "/"
	return current_path
}

const TOKEN_SECRET_KEY = "GOAPISECRETs"
const USE_SCORE_FLAG = 1
const NOT_USE_SCORE_FLAG = 0
const ACTIVE_FLAG = 1
const IN_ACTIVE_FLAG = 0

const STATUS_CODE_SUCCESS = "S"
const STATUS_CODE_ERROR = "E"

const MESSAGE_ABBR_ERROR = "ERROR"

// for add
const MESSAGE_CODE_ADD_SUCCESS = "S001"
const MESSAGE_ABBR_ADD_SUCCESS = "Success"
const MESSAGE_DESC_ADD_SUCCESS = "Add Data Success"

// for update
const MESSAGE_CODE_UPDATE_SUCCESS = "S002"
const MESSAGE_ABBR_UPDATE_SUCCESS = "Success"
const MESSAGE_DESC_UPDATE_SUCCESS = "Update Data Success"

// for delete
const MESSAGE_CODE_DELETE_SUCCESS = "S003"
const MESSAGE_ABBR_DELETE_SUCCESS = "Success"
const MESSAGE_DESC_DELETE_SUCCESS = "Delete Data Success"

// for load
const MESSAGE_CODE_LOAD_SUCCESS = "S004"
const MESSAGE_ABBR_LOAD_SUCCESS = "Success"
const MESSAGE_DESC_LOAD_SUCCESS = "Load Data Success"

// for get by id
const MESSAGE_CODE_GET_BY_ID_SUCCESS = "S005"
const MESSAGE_ABBR_GET_BY_ID_SUCCESS = "Success"
const MESSAGE_DESC_GET_BY_ID_SUCCESS = "Get by Id  Success"

// for add mapping
const MESSAGE_CODE_MAPPING_SUCCESS = "S006"
const MESSAGE_ABBR_MAPPING_SUCCESS = "Success"
const MESSAGE_DESC_MAPPING_SUCCESS = "Mapping Data Success"

// for remove mapping
const MESSAGE_CODE_REMOVE_MAPPING_SUCCESS = "S007"
const MESSAGE_ABBR_REMOVE_MAPPING_SUCCESS = "Success"
const MESSAGE_DESC_REMVOE_MAPPING_SUCCESS = "Remove Mapping Data Success"

//const CONTENT_CONTEXT  = "/dev-mchls-api/";
//const CONTENT_HOST  = "35.198.221.87";
//const CONTENT_CONTEXT  = "/content";
//const CONTENT_HOST  = "172.17.0.2";
//const CONTENT_SCHEMA = "http";
/* */
//const CONTENT_URL = CONTENT_SCHEMA+"://"+CONTENT_HOST+CONTENT_CONTEXT;
const CONTENT_URL = "https://droppointz.com/content"
const CONTENT_MAPP_URL = "content"

const CONTENT_ROOT_VIDEO_PATH = "/dpoint/video/"
const CONTENT_ROOT_CONTENT_PATH = "/dpoint/content/"

//const CONTENT_ROOT_VIDEO_PATH  = "/Users/imake/Desktop/video/";
//const CONTENT_ROOT_CONTENT_PATH  = "/Users/imake/Desktop/content/";
