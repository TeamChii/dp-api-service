package controller

import (
	"net/http"
	"strconv"

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/service"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/kataras/iris"
)

func ContentById(ctx iris.Context) {
	id := ctx.Params().Get("content_id")
	var a = service.FindContent(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}
func SaveContent(ctx iris.Context) {
	var content entity.ContentEntity
	err := ctx.ReadJSON(&content)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		var a = service.SaveContent(content, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
	/*
		content := new(entity.ContentEntity)
		content.Create_by = ctx.PostValue("create_by")
		// Source
		file, info, err := ctx.FormFile("file")
		defer file.Close()

			content_type := info.Header.Get("Content-Type")
			root_path,content_path, err := service.UploadContent(info,content_type);
			if(err == nil ) {
				content.Content_root = root_path
				content.Content_path = content_path
				content.File_name = info.Filename;
				//content.ContentSize = info.Size
				content.Content_type = info.Header.Get("Content-Type")
				var a = service.SaveContent(*content, ctx)
				ctx.StatusCode(iris.StatusOK)
				ctx.JSON(a)
			}else{
				return
			}
	*/
}
func UpdateContent(ctx iris.Context) {
	id := ctx.Params().Get("content_id")
	contnt_id, _ := strconv.Atoi(id)
	var content entity.ContentEntity
	err := ctx.ReadJSON(&content)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(utils.ResponseError())
		return
	} else {
		//merchantMaster.UpdatedBy = username
		content.Content_id = contnt_id
		var a = service.UpdateContent(content, ctx)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(a)
	}
	/*
		var content entity.ContentEntity
		content_id ,err := strconv.Atoi(id)
		content.Content_id = content_id

		file, info, err := ctx.FormFile("file")
		defer file.Close()

		content_type := info.Header.Get("Content-Type")
		root_path, content_path, err := service.UploadContent(info, content_type)
		if err == nil {
			var content_old = service.FindContent(id)
			var dataJson = content_old["data"].([]map[string]interface{})

				old_root_root := dataJson[0]["content_root"].(string)
				old_root_path := dataJson[0]["content_path"].(string)
					content.Content_root = root_path
					content.Content_path = content_path
					content.File_name = info.Filename;
					//content.ContentSize = info.Size
					content.Content_type = info.Header.Get("Content-Type")
					var a = service.UpdateContent(content, ctx,old_root_root,old_root_path)
					ctx.StatusCode(iris.StatusOK)
					ctx.JSON(a)
			}else{
				return
			}
	*/
}
func RemoveContent(ctx iris.Context) {
	content_id := ctx.Params().Get("content_id")
	var a = service.RemoveContent(content_id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(a)
}

/*
func InlineContent(r *http.Request) {
	path := r.URL.Path
	paths := strings.Split(path, "/")
	content_id := paths[len(paths)-1]
	content := service.FindContent(content_id)
	var dataJson = content["data"].([]map[string]interface{})

	//if len(content.Data)>0 && content.Data[0].Id != "" {
	content_path := dataJson[0]["content_path"].(string)
	forward_path := utils.CONTENT_CONTEXT + utils.CONTENT_MAPP_URL + "/" + content_path
	r.Host = utils.CONTENT_HOST
	r.URL.Host = r.Host
	r.URL.Scheme = utils.CONTENT_SCHEMA
	r.URL.Path = forward_path
	//}
}
*/
func RedirectToContent(ctx iris.Context) {
	id := ctx.Params().Get("content_id")
	content := service.FindContent(id)
	var dataJson = content["data"].([]map[string]interface{})
	content_path := dataJson[0]["content_path"].(string)
	newUrl := utils.CONTENT_URL + "/" + content_path
	http.Redirect(ctx.ResponseWriter(), ctx.Request(), newUrl, http.StatusSeeOther)
	//ctx.Redirect(newUrl,
	//	200)
}
