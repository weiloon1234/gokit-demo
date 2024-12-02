package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/weiloon1234/gokit/database"
	model "github.com/weiloon1234/gokit/models"
	"github.com/weiloon1234/gokit/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type User struct {
	ID                uint64    `json:"id" gorm:"column:id"`
	Username          string    `json:"username" gorm:"column:username"`
	Name              string    `json:"name" gorm:"column:name"`
	Email             string    `json:"email" gorm:"column:email"`
	EmailVerifiedAt   time.Time `json:"email_verified_at" gorm:"column:email_verified_at"`
	Password          string    `json:"password" gorm:"column:password" json:"-"`
	Password2         string    `json:"password2" gorm:"column:password2" json:"-"`
	CountryID         uint64    `json:"country_id" gorm:"column:country_id"`
	ContactCountryID  uint64    `json:"contact_country_id" gorm:"column:contact_country_id"`
	ContactNumber     string    `json:"contact_number" gorm:"column:contact_number"`
	FullContactNumber string    `json:"full_contact_number" gorm:"column:full_contact_number"`
	IntroducerUserID  uint64    `json:"introducer_user_id" gorm:"column:introducer_user_id"`
	Lang              string    `json:"lang" gorm:"column:lang"`
	Avatar            string    `json:"avatar" gorm:"column:avatar"`
	Credit1           float64   `json:"credit_1" gorm:"column:credit_1"`
	Credit2           float64   `json:"credit_2" gorm:"column:credit_2"`
	Credit3           float64   `json:"credit_3" gorm:"column:credit_3"`
	Credit4           float64   `json:"credit_4" gorm:"column:credit_4"`
	Credit5           float64   `json:"credit_5" gorm:"column:credit_5"`
	BankID            uint64    `json:"bank_id" gorm:"column:bank_id"`
	BankAccountName   string    `json:"bank_account_name" gorm:"column:bank_account_name"`
	BankAccountNumber string    `json:"bank_account_number" gorm:"column:bank_account_number"`
	NationalID        string    `json:"national_id" gorm:"column:national_id"`
	FirstLogin        uint8     `json:"first_login" gorm:"column:first_login"`
	BanUntil          time.Time `json:"ban_until" gorm:"column:ban_until"`
	NewLoginAt        time.Time `json:"new_login_at" gorm:"column:new_login_at"`
	LastLoginAt       time.Time `json:"last_login_at" gorm:"column:last_login_at"`
	Unilevel          uint64    `json:"unilevel" gorm:"column:unilevel"`
	model.BaseModel
	model.SoftDeleteModel
	Introducer     *User   `json:"introducer" gorm:"foreignKey:IntroducerUserID;references:ID"`
	Country        Country `json:"country" gorm:"foreignKey:CountryID;references:ID"`
	ContactCountry Country `json:"contact_country" gorm:"foreignKey:ContactCountryID;references:ID"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	//Before create lifecycle
	return
}

func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	//Before save lifecycle
	return
}

// SetPassword setter
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// SetPassword2 setter
func (u *User) SetPassword2(password2 string) error {
	hashedPassword2, err := bcrypt.GenerateFromPassword([]byte(password2), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password2 = string(hashedPassword2)
	return nil
}

// GetIdentity getter
func (u *User) GetIdentity() string {
	return u.Username
}

// SetAvatar setter, the default GORM setter don't have gin.Context, so this method SetAvatar must be call exactly
func (u *User) SetAvatar(c *gin.Context, fileHeader *multipart.FileHeader) error {
	filePath, err := utils.UploadImage(c, fileHeader)
	if err != nil {
		return err
	}

	u.Avatar = filePath

	return nil
}

func (u *User) GetAvatar() string {
	if u.Avatar == "" {
		return ""
	}

	return utils.GetFile(u.Avatar)
}

func GetUserByID(userID uint64) (*User, error) {
	var user User
	err := database.GetDB().First(&user, "id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &user, nil
}
