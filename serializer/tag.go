package serializer

import "record-server/model"

type Tag struct {
	ID        uint   `json:"id"`
	TagName   string `json:"tag_name"`
	CreatedAt int64  `json:"created_at"`
}

func BuildTag(tag model.Tag) Tag {
	return Tag{
		ID:        tag.ID,
		TagName:   tag.TagName,
		CreatedAt: tag.CreatedAt.Unix(),
	}
}

func BuildTagResponse(tag model.Tag) Response {
	return Response{
		Data: BuildTag(tag),
	}
}

func BuildTagListResponse(tags []model.Tag) Response {
	var respTags []Tag
	for _, tag := range tags {
		respTags = append(respTags, BuildTag(tag))
	}
	return Response{
		Data: respTags,
	}
}
