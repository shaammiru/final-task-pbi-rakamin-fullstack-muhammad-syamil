package controllers

import (
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/database"
	"github.com/shaammiru/task-5-pbi-fullstack-developer-muhammadsyamil/models"
)

func CreatePhoto(photoData models.Photo) (models.Photo, error) {
	newPhoto := photoData

	result := database.DB.Create(&newPhoto)
	if result.Error != nil {
		return newPhoto, result.Error
	}

	return newPhoto, nil
}

func ListPhoto() ([]models.Photo, error) {
	var photos []models.Photo

	result := database.DB.Find(&photos)
	if result.Error != nil {
		return photos, result.Error
	}

	return photos, nil
}

func GetPhotoByID(photoID string) (models.Photo, error) {
	var photo models.Photo

	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		return photo, result.Error
	}

	return photo, nil
}

func UpdatePhotoByID(photoID string, photoData models.Photo) (models.Photo, error) {
	photo, err := GetPhotoByID(photoID)
	if err != nil {
		return photo, err
	}

	if photoData.Title != "" {
		photo.Title = photoData.Title
	}

	if photoData.Caption != "" {
		photo.Caption = photoData.Caption
	}

	if photoData.PhotoURL != "" {
		photo.PhotoURL = photoData.PhotoURL
	}

	if photoData.UserID != 0 {
		photo.UserID = photoData.UserID
	}

	result := database.DB.Save(&photo)
	if result.Error != nil {
		return photo, result.Error
	}

	return photo, nil
}

func DeletePhotoByID(photoID string) error {
	photo, err := GetPhotoByID(photoID)
	if err != nil {
		return err
	}

	result := database.DB.Delete(&photo)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
