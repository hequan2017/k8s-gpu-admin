package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu"
	k8sgpuReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu/request"
)

type ImageRegistryService struct{}

var ImageRegistryServiceApp = new(ImageRegistryService)

// CreateImageRegistry 创建镜像
func (irs *ImageRegistryService) CreateImageRegistry(req k8sgpuReq.ImageRegistryAdd) (err error) {
	var image k8sgpu.ImageRegistry
	image.Name = req.Name
	image.Address = req.Address
	image.Description = req.Description
	image.Source = req.Source
	image.IsPublished = req.IsPublished
	image.Remark = req.Remark
	err = global.GVA_DB.Create(&image).Error
	return err
}

// DeleteImageRegistry 删除镜像
func (irs *ImageRegistryService) DeleteImageRegistry(id uint) (err error) {
	err = global.GVA_DB.Delete(&k8sgpu.ImageRegistry{}, id).Error
	return err
}

// UpdateImageRegistry 更新镜像
func (irs *ImageRegistryService) UpdateImageRegistry(req k8sgpuReq.ImageRegistryUpdate) (err error) {
	var image k8sgpu.ImageRegistry
	image.ID = req.ID
	image.Name = req.Name
	image.Address = req.Address
	image.Description = req.Description
	image.Source = req.Source
	if req.IsPublished != nil {
		image.IsPublished = *req.IsPublished
	}
	image.Remark = req.Remark
	err = global.GVA_DB.Save(&image).Error
	return err
}

// GetImageRegistry 获取镜像信息
func (irs *ImageRegistryService) GetImageRegistry(id uint) (image k8sgpu.ImageRegistry, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&image).Error
	return
}

// GetImageRegistryList 分页获取镜像列表
func (irs *ImageRegistryService) GetImageRegistryList(info k8sgpuReq.ImageRegistrySearch) (list []k8sgpu.ImageRegistry, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&k8sgpu.ImageRegistry{})

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Source != "" {
		db = db.Where("source = ?", info.Source)
	}
	if info.IsPublished != nil {
		db = db.Where("is_published = ?", *info.IsPublished)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id DESC").Find(&list).Error
	return list, total, err
}

// GetPublishedImageRegistryList 获取已上架的镜像列表(用于实例创建选择)
func (irs *ImageRegistryService) GetPublishedImageRegistryList() (list []k8sgpu.ImageRegistry, err error) {
	err = global.GVA_DB.Where("is_published = ?", true).Order("id DESC").Find(&list).Error
	return
}
