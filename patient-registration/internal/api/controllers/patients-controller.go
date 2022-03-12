package controllers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	models "github.com/yaska1706/AfyaPevu/patient-registration/internal/pkg/models/patients"
	"github.com/yaska1706/AfyaPevu/patient-registration/internal/pkg/persistence"
	http_err "github.com/yaska1706/AfyaPevu/patient-registration/pkg/http-err"
)

type PatientRequest struct {
	Name         string `json:"name" binding:"required"`
	IDNumber     int    `json:"id_number"`
	Location     string `json:"location"`
	Gender       GenderType
	ProfileImage string `json:"profile_image"`
	IDImage      string `json:"id_image"`
}

type GenderType struct {
	Type string `json:"type"`
}

func CreatePatient(c *gin.Context) {
	s := persistence.GetPatientRepository()
	var patientRequest PatientRequest
	_ = c.BindJSON(&patientRequest)
	patientRequest.ProfileImage, _ = SaveImage(c.Writer, *c.Request)
	patient := models.Patient{
		Name:         patientRequest.Name,
		IDNumber:     patientRequest.IDNumber,
		Location:     patientRequest.Location,
		Gender:       patientRequest.Gender.Type,
		ProfileImage: patientRequest.ProfileImage,
	}
	if err := s.Add(&patient); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, patient)
	}
}

func GetPatients(c *gin.Context) {
	s := persistence.GetPatientRepository()
	var q models.Patient
	_ = c.Bind(&q)
	if patients, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("patients not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, patients)
	}
}

func UpdatePatient(c *gin.Context) {
	s := persistence.GetPatientRepository()
	id := c.Params.ByName("id")
	var patientInput PatientRequest
	_ = c.BindJSON(&patientInput)
	if patient, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("patient not found"))
		log.Println(err)
	} else {
		patient.Name = patientInput.Name
		patient.Location = patientInput.Location
		patient.Gender = patientInput.Gender.Type
		if err := s.Update(patient); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, patient)
		}
	}
}

func DeletePatient(c *gin.Context) {
	s := persistence.GetPatientRepository()
	id := c.Params.ByName("id")
	var patientInput PatientRequest
	_ = c.BindJSON(&patientInput)
	if patient, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("patient not found"))
		log.Println(err)
	} else {
		if err := s.Delete(patient); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}

func SaveImage(w gin.ResponseWriter, r http.Request) (string, error) {
	//this function returns the filename(to save in database) of the saved file or an error if it occurs
	r.ParseMultipartForm(32 << 20)
	//ParseMultipartForm parses a request body as multipart/form-data
	file, handler, err := r.FormFile("file") //retrieve the file from form data
	if err != nil {
		return "", err
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {

		http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
		return "", err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	//this is path which  we want to store the file
	f, err := os.OpenFile("../../../data/image/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	defer f.Close()
	io.Copy(f, file)
	//here we save our file to our path
	return handler.Filename, nil

}

func UploadImage(c gin.Context) {

}
