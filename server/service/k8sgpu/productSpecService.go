package k8sgpu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu"
	k8sgpuReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sgpu/request"
	"github.com/pkg/errors"
)

type ProductSpecService struct{}

var ProductSpecServiceApp = new(ProductSpecService)

// CreateProductSpec 创建产品规格
func (pss *ProductSpecService) CreateProductSpec(req k8sgpuReq.ProductSpecAdd) (err error) {
	var spec k8sgpu.ProductSpec
	spec.Name = req.Name
	spec.GPUModel = req.GPUModel
	spec.GPUCount = req.GPUCount
	spec.CPUCores = req.CPUCores
	spec.Memory = req.Memory
	spec.SystemDisk = req.SystemDisk
	spec.DataDisk = req.DataDisk
	spec.PricePerHour = req.PricePerHour
	spec.IsPublished = req.IsPublished
	spec.Remark = req.Remark
	err = global.GVA_DB.Create(&spec).Error
	return err
}

// DeleteProductSpec 删除产品规格
func (pss *ProductSpecService) DeleteProductSpec(id uint) (err error) {
	// 检查是否有关联的实例
	var count int64
	global.GVA_DB.Model(&k8sgpu.Instance{}).Where("spec_id = ?", id).Count(&count)
	if count > 0 {
		return errors.New("该规格下还有实例，无法删除")
	}
	err = global.GVA_DB.Delete(&k8sgpu.ProductSpec{}, id).Error
	return err
}

// UpdateProductSpec 更新产品规格
func (pss *ProductSpecService) UpdateProductSpec(req k8sgpuReq.ProductSpecUpdate) (err error) {
	var spec k8sgpu.ProductSpec
	spec.ID = req.ID
	spec.Name = req.Name
	spec.GPUModel = req.GPUModel
	spec.GPUCount = req.GPUCount
	spec.CPUCores = req.CPUCores
	spec.Memory = req.Memory
	spec.SystemDisk = req.SystemDisk
	spec.DataDisk = req.DataDisk
	spec.PricePerHour = req.PricePerHour
	if req.IsPublished != nil {
		spec.IsPublished = *req.IsPublished
	}
	spec.Remark = req.Remark
	err = global.GVA_DB.Save(&spec).Error
	return err
}

// GetProductSpec 获取产品规格信息
func (pss *ProductSpecService) GetProductSpec(id uint) (spec k8sgpu.ProductSpec, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&spec).Error
	return
}

// GetProductSpecList 分页获取产品规格列表
func (pss *ProductSpecService) GetProductSpecList(info k8sgpuReq.ProductSpecSearch) (list []k8sgpu.ProductSpec, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&k8sgpu.ProductSpec{})

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.GPUModel != "" {
		db = db.Where("gpu_model = ?", info.GPUModel)
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

// GetPublishedProductSpecList 获取已上架的产品规格列表(用于实例创建选择)
func (pss *ProductSpecService) GetPublishedProductSpecList() (list []k8sgpu.ProductSpec, err error) {
	err = global.GVA_DB.Where("is_published = ?", true).Order("id DESC").Find(&list).Error
	return
}
