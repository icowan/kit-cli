package foo

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TODO 以下都是参考代码，不需要可以直接删除

type SysSetting struct {
	Id          int64      `gorm:"column:id;rimary_key" json:"id"`
	Key         string     `gorm:"column:key;index;unique;size:128" json:"key"`            // 标识
	Value       string     `gorm:"column:value;notnull;size:5000" json:"value"`            // 值
	Description string     `gorm:"column:description;notnull;size:500" json:"description"` // 备注
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`                    // 创建时间
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`                    // 更新时间
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                    // 删除时间
}

// TableName set table
func (*SysSetting) TableName() string {
	return "sys_setting"
}

type SysSettingRepository interface {
	Add(key, val, desc string) (err error)
	Delete(key string) (err error)
	Update(data *SysSetting) (err error)
	Find(key string) (res SysSetting, err error)
}

type sysSettingRepository struct {
	db *gorm.DB
}

func (s *sysSettingRepository) Delete(key string) (err error) {
	return s.db.Model(&types.SysSetting{}).Where("key = ?", key).Delete(&types.SysSetting{}).Error
}

func (s *sysSettingRepository) Update(data *types.SysSetting) (err error) {
	return s.db.Model(data).Where("id = ?", data.Id).Update(data).Error
}

func (s *sysSettingRepository) Find(key string) (res types.SysSetting, err error) {
	err = s.db.Model(&types.SysSetting{}).Where("`key` = ?", key).First(&res).Error
	return
}

func (s *sysSettingRepository) Add(key, val, desc string) (err error) {
	return s.db.Model(&types.SysSetting{}).Save(&types.SysSetting{
		Key:         key,
		Value:       val,
		Description: desc,
	}).Error
}

func NewSysSettingRepository(db *gorm.DB) SysSettingRepository {
	return &sysSettingRepository{db: db}
}
