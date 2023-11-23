package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BaseTime struct {
	CreatedAt uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:创建时间"` // 使用时间戳秒数填充创建时间
	UpdatedAt uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:修改时间"` // 使用时间戳秒数填充修改时间
}

func (bt *BaseTime) BeforeCreate(tx *gorm.DB) (err error) {
	timestamp := uint32(time.Now().Unix())
	bt.CreatedAt = timestamp
	bt.UpdatedAt = timestamp
	return
}

func (bt *BaseTime) BeforeUpdate(tx *gorm.DB) (err error) {
	bt.UpdatedAt = uint32(time.Now().Unix())
	return
}

func Init(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&ProductCategory{},

		&ProductCategoryAttributeRelation{},

		&Brand{},

		&ProductAttributeCategory{},
		&ProductAttribute{},
		&ProductAttributeValue{},

		&Product{},
		&SkuStock{},
		&ProductLadder{},
		&ProductFullReduction{},
		&MemberPrice{},

		&Comment{},
		&CommentReplay{},

		&ProductVertifyRecord{},
		&ProductOperateLog{},
	); err != nil {
		return err
	}

	return nil
}
