package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

//toda vez que for creado um "objeto" de Base os campos obrigatorios vão ser validados
func init() {
	govalidator.SetFieldsRequireByDefault(true)
}

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"update_at" valid:"-"`
}
