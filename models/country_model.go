package model

import (
	"github.com/shopspring/decimal"
	"github.com/weiloon1234/gokit/models"
)

type Country struct {
	ID             uint            `json:"id" gorm:"column:id;primaryKey"`
	Iso2           string          `json:"iso2" gorm:"column:iso2"`                       // ISO 3166-1 alpha-2 country code
	Iso3           string          `json:"iso3" gorm:"column:iso3"`                       // ISO 3166-1 alpha-3 country code
	Name           string          `json:"name" gorm:"column:name"`                       // Country name
	OfficialName   string          `json:"official_name" gorm:"column:official_name"`     // Official country name
	NumericCode    string          `json:"numeric_code" gorm:"column:numeric_code"`       // ISO 3166-1 numeric code
	PhoneCode      string          `json:"phone_code" gorm:"column:phone_code"`           // Country calling code, e.g. +1 for US
	Capital        string          `json:"capital" gorm:"column:capital"`                 // Capital city of the country
	CurrencyName   string          `json:"currency_name" gorm:"column:currency_name"`     // Currency name
	CurrencyCode   string          `json:"currency_code" gorm:"column:currency_code"`     // ISO 4217 currency code
	CurrencySymbol string          `json:"currency_symbol" gorm:"column:currency_symbol"` // Currency symbol
	ConversionRate decimal.Decimal `gorm:"column:conversion_rate;type:decimal(18,5);not null;default:0.00000" json:"conversion_rate"`
	Status         uint8           `json:"status" gorm:"column:status"`
	model.BaseModel
	States []CountryLocations `json:"locations" gorm:"foreignKey:CountryID;references:ID"` // HasMany relationship
}

func (m *Country) TableName() string {
	return "countries"
}
