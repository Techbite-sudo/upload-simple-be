package videos

import (
	"Galaxy/engine"
	"Galaxy/graph/model"
	"Galaxy/models"
	"Galaxy/utils"
	"errors"
)

func UploadVideo(input model.UploadVideoInput) (*model.Video, error) {
	video := models.Video{
		Title:        input.Title,
		Description:  input.Description,
		Duration:     input.Duration,
		ThumbnailURL: input.ThumbnailURL,
		VideoURL:     input.File,
	}
	err := utils.DB.Create(&video).Error
	if err != nil {
		return nil, err
	}
	return video.ToGraphData(), nil
}

func Videos() ([]*model.Video, error) {
	videos, err := engine.FetchAllVideos()
	if err != nil {
		return nil, errors.New("videos could not be fetched")
	}
	var videoGraphData []*model.Video
	for _, video := range videos {
		videoGraphData = append(videoGraphData, video.ToGraphData())
	}
	return videoGraphData, nil
}

func Video(id string) (*model.Video, error) {
	video, err := engine.FetchVideoById(id)
	if err != nil {
		return nil, errors.New("video could not be fetched")
	}
	return video.ToGraphData(), nil
}
