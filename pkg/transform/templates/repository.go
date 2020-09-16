/**
 * @Time : 2020/9/16 4:26 PM
 * @Author : solacowa@gmail.com
 * @File : repository
 * @Software: GoLand
 */

package foo

import "github.com/jinzhu/gorm"

type Repository interface {
	SysSetting() SysSettingRepository
}

type repository struct {
	sysSetting SysSettingRepository
}

func (r *repository) SysSetting() SysSettingRepository {
	return r.sysSetting
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		sysSetting: NewSysSettingRepository(db),
	}
}
