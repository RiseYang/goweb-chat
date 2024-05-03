package img_upload_connector

import (
	"goweb-blog/services"
	"goweb-blog/services/img_freeimage"
)

// 定义serve 的映射关系
var serveMap = map[string]services.ImgUploadInterface{
	"fi": &img_freeimage.ImgFreeImageService{},
}

func ImgCreate() services.ImgUploadInterface {
	return serveMap["fi"]
}
