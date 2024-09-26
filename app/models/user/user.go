package user

import "github.com/VENI-VIDIVICI/plus/app/models"

type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Password string `json:"-"`
	Phone    string `json:"-"`
	models.CommonTimestampsField
}
