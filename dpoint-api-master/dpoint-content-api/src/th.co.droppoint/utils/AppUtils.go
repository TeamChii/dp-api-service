package utils

import (
	"math/rand"
	"time"
	"os"
)

const CHARSET_RANDOM_WITH_STRING = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func RandString(n int) string {
	return stringWithCharset(n, CHARSET_RANDOM_WITH_STRING)
}

func CheckContentPath(root_path string) string{
	t := time.Now();
	year_dir := t.Format("2006")
	month_dir := t.Format("01")
	day_dir := t.Format("02")
	if  _, err := os.Stat(root_path+year_dir);
		os.IsNotExist(err) {
		os.MkdirAll(root_path+year_dir,os.ModePerm)
	}
	if _, err := os.Stat(root_path+year_dir+"/"+month_dir);
		os.IsNotExist(err) {
		os.MkdirAll(root_path+year_dir+"/"+month_dir,os.ModePerm)
	}
	if _, err := os.Stat(root_path+year_dir+"/"+month_dir+"/"+day_dir);
		os.IsNotExist(err) {
		os.MkdirAll(root_path+year_dir+"/"+month_dir+"/"+day_dir,os.ModePerm)
	}
	//current_path := root_path+year_dir+"/"+month_dir+"/"+day_dir+"/";
	current_path := year_dir+"/"+month_dir+"/"+day_dir+"/";
	return current_path;
}




//const CONTENT_CONTEXT  = "/dev-mchls-api/";
//const CONTENT_HOST  = "35.198.221.87";
//const CONTENT_CONTEXT  = "/content";
//const CONTENT_HOST  = "172.17.0.2";
//const CONTENT_SCHEMA = "http";
/* */
//const CONTENT_URL = CONTENT_SCHEMA+"://"+CONTENT_HOST+CONTENT_CONTEXT;
//const CONTENT_MAPP_URL  = "content"

const CONTENT_ROOT_VIDEO_PATH  = "/dpoint/video/";
const CONTENT_ROOT_CONTENT_PATH  = "/dpoint/content/";

//const CONTENT_ROOT_VIDEO_PATH  = "/Users/imake/Desktop/video/";
//const CONTENT_ROOT_CONTENT_PATH  = "/Users/imake/Desktop/content/";

const SERVICE_API_HOST  = "http://172.17.0.4:8002/api/content";
//const SERVICE_API_HOST  = "http://172.17.0.5:8002/api/content";
//const SERVICE_API_HOST  = "http://localhost:8002/api/content";