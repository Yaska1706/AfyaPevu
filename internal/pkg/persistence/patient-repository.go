package persistence

import (
	"strconv"

	"github.com/yaska1706/AfyaPevu/internal/pkg/db"
	models "github.com/yaska1706/AfyaPevu/internal/pkg/models/patients"
)

type PatientRepository struct{}

var patientRepository *PatientRepository

func GetPatientRepository() *PatientRepository {
	if patientRepository == nil {
		patientRepository = &PatientRepository{}
	}
	return patientRepository
}
func (r *PatientRepository) Get(id string) (*models.Patient, error) {
	var Patient models.Patient
	where := models.Patient{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Patient, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &Patient, err
}

func (r *PatientRepository) All() (*[]models.Patient, error) {
	var Patients []models.Patient
	err := Find(&models.Patient{}, &Patients, []string{"Role"}, "id asc")
	return &Patients, err
}

func (r *PatientRepository) Query(q *models.Patient) (*[]models.Patient, error) {
	var Patients []models.Patient
	err := Find(&q, &Patients, nil)
	return &Patients, err
}

func (r *PatientRepository) Add(Patient *models.Patient) error {
	err := Create(&Patient)
	err = Save(&Patient)
	return err
}

func (r *PatientRepository) Update(Patient *models.Patient) error {

	err := db.GetDB().Save(&Patient).Error

	return err
}

func (r *PatientRepository) Delete(Patient *models.Patient) error {
	err := db.GetDB().Unscoped().Delete(&Patient).Error
	return err
}
