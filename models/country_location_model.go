package model

import "github.com/weiloon1234/gokit/models"

type CountryLocations struct {
	ID        uint64 `json:"id" gorm:"column:id"`
	CountryID uint64 `json:"country_id" gorm:"column:country_id"` // Country of the location
	ParentID  uint64 `json:"parent_id" gorm:"column:parent_id"`   // If present, it is area, this is state id
	Sorting   uint64 `json:"sorting" gorm:"column:sorting"`
	NameEn    string `json:"name_en" gorm:"column:name_en"` // Bank name in english
	NameZh    string `json:"name_zh" gorm:"column:name_zh"` // Bank name in chinese
	model.BaseModel
	model.SoftDeleteModel
	Country Country            `json:"country" gorm:"foreignKey:CountryID;references:ID"`
	State   *CountryLocations  `json:"state" gorm:"foreignKey:ParentID;references:ID"` // Self-referencing relationship
	Areas   []CountryLocations `json:"sub_locations" gorm:"foreignKey:ParentID"`       // Reverse relationship (children)
}

func (m *CountryLocations) TableName() string {
	return "country_locations"
}

func (m *Country) GetMultilanguageColumns() []string {
	return []string{"name"}
}
