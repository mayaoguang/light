package mysql

import (
	"testing"
	"time"
)

func TestGetDB(t *testing.T) {
	if err := Init(Conf{}, Conf{}); err != nil {
		return
	}
	type Light struct {
		Id         uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
		Name       string `gorm:"column:name;default:'';comment:'姓名'" json:"name"`
		CreateTime int64  `gorm:"column:create_time;default:0" json:"createTime"`
		UpdateTime int64  `gorm:"column:update_time;default:0" json:"updateTime"`
	}
	tx := NewWriteDB()
	if err := tx.AutoMigrate(Light{}); err != nil {
		t.Error(err)
	}
	lightD := &Light{Name: "light"}
	if err := tx.Create(lightD).Error; err != nil {
		t.Error(err)
	}
	if err := tx.Take(lightD).Error; err != nil {
		t.Error(err)
	}
	time.Sleep(2 * time.Second)
	lightD.Name = "light man"
	if err := tx.Updates(lightD).Error; err != nil {
		t.Error(err)
	}
	// 批量插入钩子函数只对第一条数据生效
	if err := tx.Create(&[]Light{{Name: "1"}, {Name: "2"}}).Error; err != nil {
		t.Error(err)
	}
}
