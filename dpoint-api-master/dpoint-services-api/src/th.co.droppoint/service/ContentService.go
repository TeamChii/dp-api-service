package service

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	/*
		"th.co.droppoint/config"
		"th.co.droppoint/entity"
		"th.co.droppoint/utils"
	*/

	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/config"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/entity"
	"github.com/TeamChii/dpoint-api-master/dpoint-services-api/src/th.co.droppoint/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func FindContent(content_id string) map[string]interface{} {
	var data entity.ContentEntity
	check := config.DBsql().Where("content_id = ?", content_id).Find(&data).RecordNotFound()

	var dataToJson map[string]interface{}
	if data.Create_date != nil && data.Create_date.Year() != 1 {
		data.CreatedDateStr = data.Create_date.Format("02/01/2006 15:04:05")
	}
	if data.Update_date != nil && data.Update_date.Year() != 1 {
		data.UpdatedDateStr = data.Update_date.Format("02/01/2006 15:04:05")
	}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &dataToJson)

	return utils.ResDataLoadById(check, dataToJson)
}
func UploadContent(file *multipart.FileHeader, content_type string) (string, string, error) {
	file_names := strings.Split(file.Filename, ".")
	extension := file_names[len(file_names)-1]
	root_path := utils.CONTENT_ROOT_CONTENT_PATH
	if strings.Contains(content_type, "video") {
		root_path = utils.CONTENT_ROOT_VIDEO_PATH
	}
	current_path := utils.CheckContentPath(root_path) // year/month/day
	file_name_gen := utils.RandString(20)
	path_gen := current_path + file_name_gen + "." + extension

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(root_path + path_gen)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", "", err
	}
	return root_path, path_gen, nil
}

func SaveContent(content entity.ContentEntity, ctx iris.Context) map[string]string {
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	content.Create_by = utils.Decode(jwt.Raw).(string)
	now := time.Now()
	content.Create_date = &now
	err := config.DBsql().Create(&content).Error
	return utils.ResDataAddWithId(err, strconv.Itoa(content.Content_id), content.Content_path)

}
func UpdateContent(content entity.ContentEntity, ctx iris.Context) map[string]string {
	//,content_old_root string,content_old_path string
	jwt := ctx.Values().Get("jwt").(*jwt.Token)
	//file_root := content_old_root ;//result.ContentRoot
	//file_path := content_old_path;//result.ContentPath
	content.Update_by = utils.Decode(jwt.Raw).(string)
	now := time.Now()
	content.Update_date = &now
	err := config.DBsql().Model(&content).Where("content_id = ?", content.Content_id).Update(&content).Error

	// delete old file content
	//os.RemoveAll(file_root+file_path)

	//return utils.ResDataEdit(err)
	return utils.ResDataEditWithId(err, strconv.Itoa(content.Content_id), content.Content_path)
}
func RemoveContent(content_id string) map[string]string {

	var data entity.ContentEntity
	check := config.DBsql().Where("content_id = ?", content_id).Find(&data).RecordNotFound()
	if check == false {
		file_root := data.Content_root
		file_path := data.Content_path
		var count int
		var deleted = 0
		conn := config.DBsql()
		conn.Where("content_id = ?", content_id).
			Table("dp_tb_content").Count(&count).Delete(entity.ContentEntity{})
		if count != 0 {
			deleted = deleted + 1
		}
		os.RemoveAll(file_root + file_path)
	}
	return utils.ResDataDel(1, 1)
}
