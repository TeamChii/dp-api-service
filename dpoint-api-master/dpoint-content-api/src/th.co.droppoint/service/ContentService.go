package service

import (
	"th.co.droppoint/utils"
	"os"
	"mime/multipart"
	"strings"
	"io"
	"github.com/kataras/iris"
	log "github.com/sirupsen/logrus"
	"net/http"

	"io/ioutil"
	"bytes"
	"encoding/json"
	"github.com/nfnt/resize"
	"strconv"

	//"image/gif"
	"image/jpeg"
)


func UploadContent(file *multipart.FileHeader ,content_type string,width string ,height string) ( string,string ,error){
	file_names := strings.Split(file.Filename, ".")
	extension := file_names[len(file_names)-1]
	root_path := utils.CONTENT_ROOT_CONTENT_PATH;
	if strings.Contains(content_type, "video") {
		root_path = utils.CONTENT_ROOT_VIDEO_PATH;
	}
	current_path := utils.CheckContentPath(root_path) // year/month/day
	file_name_gen := utils.RandString(20)
	path_gen := current_path+file_name_gen+"."+extension


	src, err := file.Open()
	if err != nil {
		return "","",err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(root_path+path_gen)
	if err != nil {
		return "","",err
	}
	defer dst.Close()

	if width != "" && height != "" {
		img, err := jpeg.Decode(src)
		//img, err := png.Decode(src)
		//img, err := gif.Decode(src)
		if err != nil {
			return "","",err
		}
		widthInt, err := strconv.Atoi(width)
		if err != nil {
			return "","",err;
		}
		heightInt, err := strconv.Atoi(height)
		if err != nil {
			return "","",err;
		}
		m := resize.Resize(uint(widthInt),uint(heightInt), img, resize.Lanczos3)
		// write new image to file
		jpeg.Encode(dst, m, nil)
		//png.Encode(dst, m)
		//gif.Encode(dst, m,nil)
	}else{
		// Copy
		//if _, err = io.Copy(dst, src); err != nil {
		if _, err = io.Copy(dst, src); err != nil {
			return "","",err
		}
	}







	return root_path, path_gen, nil;
}

func UpdateContent(content_id string,file_name string,content_root string, content_path string , content_type string, ctx iris.Context,
	content_old_root string,content_old_path string) map[string]interface{} {
	file_root := content_old_root ;
	file_path := content_old_path;

	url := utils.SERVICE_API_HOST+"/update/"+content_id
	log.Info("URL:>", url)

	var jsonStr = []byte(`{
					"file_name":"`+file_name+`",
                    "content_root":"`+content_root+`",
					"content_path":"`+content_path+`",
					"content_type":"`+content_type+`"
					}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", ctx.GetHeader("Authorization"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	//json.Unmarshal([]byte(body), &result)
	//datas := result["data"].(map[string]interface{})
	json.Unmarshal(body, &result)

	// delete old file content
	os.RemoveAll(file_root+file_path)

	 return result
}
func RemoveContent(authorization string ,content_id string) map[string]interface{}{
	url := utils.SERVICE_API_HOST+"/remove/"+content_id
	log.Info("URL:>", url)

	req, err := http.NewRequest("GET", url,nil)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return result
}