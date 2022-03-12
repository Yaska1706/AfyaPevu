package patients

import (
	"time"

	"github.com/yaska1706/AfyaPevu/patient-registration/internal/pkg/models"
)

type Patient struct {
	models.Model
	Name         string `gorm:"column:name;not null;unique_index:name" json:"name" form:"name"`
	IDNumber     int    `gorm:"column:id_number;not null;unique_index:id_number" json:"id_number" form:"id_number"`
	Location     string `gorm:"column:location;not null;unique_index:location" json:"location" form:"location"`
	Gender       string `gorm:"column:gender;not null;unique_index:gender" json:"gender" form:"gender"`
	ProfileImage string `gorm:"column:profile_image;not null;unique_index:profile_image" json:"profile_image" form:"profile_image"`
}

func (m *Patient) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Patient) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
