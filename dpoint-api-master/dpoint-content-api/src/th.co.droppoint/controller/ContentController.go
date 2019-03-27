package controller

import (
	"github.com/kataras/iris"
	"th.co.droppoint/service"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	 "th.co.droppoint/utils"
)

func SaveContent(ctx iris.Context)  {
	// Source
	file , info , err  :=  ctx.FormFile("file")
	defer file.Close()
	var width = ctx.FormValue("width")
	var height = ctx.FormValue("height")
	//println(width+","+height)
	content_type := info.Header.Get("Content-Type")
	root_path,content_path, err := service.UploadContent(info,content_type,width,height);
	if(err == nil ) {
		url := utils.SERVICE_API_HOST+"/upload"

		var jsonStr = []byte(`{
					"file_name":"`+info.Filename+`",
                    "content_root":"`+root_path+`",
					"content_path":"`+content_path+`",
					"content_type":"`+info.Header.Get("Content-Type")+`"
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

		//fmt.Println("response Status:", resp.Status)
		//fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println("response Body:", string(body))
		ctx.StatusCode(iris.StatusOK)
		var result map[string]interface{}
		//json.Unmarshal([]byte(body), &result)
		//datas := result["data"].(map[string]interface{})
		json.Unmarshal(body, &result)

		ctx.JSON(result)
	}else{
		return
	}


}
func UpdateContent(ctx iris.Context) {
	id := ctx.Params().Get("content_id")
		file , info , err  :=  ctx.FormFile("file")
		defer file.Close()
		var width = ctx.FormValue("width")
		var height = ctx.FormValue("height")
		content_type := info.Header.Get("Content-Type")
		root_path,content_path, err := service.UploadContent(info,content_type,width,height);
		if(err == nil ) {
			//var content_old = service.FindContent(id)
			url := utils.SERVICE_API_HOST+"/get/"+id
		//	fmt.Println("URL:>", url)

			req, err := http.NewRequest("GET", url, nil)
			req.Header.Set("Authorization", ctx.GetHeader("Authorization"))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			//fmt.Println("response Status:", resp.Status)
			//fmt.Println("response Headers:", resp.Header)
			body, _ := ioutil.ReadAll(resp.Body)
			//fmt.Println("response Body:", string(body))
			var result map[string]interface{}
			//json.Unmarshal([]byte(body), &result)
			//datas := result["data"].(map[string]interface{})
			json.Unmarshal(body, &result)
			//var dataJson = string(body)["data"].([]map[string]interface{})
			//dataJson := result["data"].([]map[string]interface{})
			dataJson := result["data"].([]interface{})


			old_root_root :=dataJson[0].(map[string]interface{})["content_root"].(string)
			old_root_path := dataJson[0].(map[string]interface{})["content_path"].(string)

				var a = service.UpdateContent(id,info.Filename,root_path,content_path,info.Header.Get("Content-Type"),
					ctx,old_root_root,old_root_path)
				ctx.StatusCode(iris.StatusOK)
				ctx.JSON(a)
		}else{
			return
		}
}
func RemoveContent(ctx iris.Context) {
	content_id := ctx.Params().Get("content_id")
	var a = service.RemoveContent(ctx.GetHeader("Authorization"),content_id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
