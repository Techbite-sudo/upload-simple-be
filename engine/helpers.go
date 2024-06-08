package engine

import (
	"Galaxy/models"
	"Galaxy/utils"
	"errors"

	"gorm.io/gorm/clause"
)

func FetchVideoById(id string)(*models.Video, error) {
	var video models.Video
	err := utils.DB.Preload(clause.Associations).Where("id = ?", id).First(&video).Error
	if err!= nil {
		return nil, errors.New("invalid id video not found")
	}
	return &video, nil
}

func FetchAllVideos()([]models.Video, error) {
	var videos []models.Video
    err := utils.DB.Preload(clause.Associations).Find(&videos).Error
    if err != nil {
        return nil, errors.New("Videos could not be found")
    }
    return videos, nil
}
