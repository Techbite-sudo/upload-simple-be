package models

import (
	"Galaxy/graph/model"
)

type Video struct {
	Base
	Title        string `json:"title"`
	Description  string `json:"description"`
	Duration     int    `json:"duration"`
	ThumbnailURL string `json:"thumbnailUrl"`
	VideoURL     string `json:"videoUrl"`
}

func (v *Video) ToGraphData() *model.Video {
	return &model.Video{
		ID:           v.ID.String(),
		Title:        v.Title,
		Description:  v.Description,
		Duration:     v.Duration,
		ThumbnailURL: v.ThumbnailURL,
		VideoURL:     v.VideoURL,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}
}
