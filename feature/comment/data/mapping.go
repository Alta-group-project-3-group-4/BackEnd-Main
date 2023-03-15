package data

import (
	"airbnb/feature/comment"

	"gorm.io/gorm"
)

func CoreToData(core comment.Core) Comment {
	return Comment{
		Model:       gorm.Model{ID: core.ID},
		UserId:      core.UserId,
		HomeId:      core.HomeId,
		Rate:        core.Rate,
		Description: core.Description,
	}
}

func DataToCore(data Comment) comment.Core {
	return comment.Core{
		ID:          data.ID,
		UserId:      data.UserId,
		HomeId:      data.HomeId,
		Rate:        data.Rate,
		Description: data.Description,
	}
}

func (dataModel *HomeComment) HomeCommentToCore() comment.Core {
	return comment.Core{
		ID:          dataModel.ID,
		UserId:      dataModel.UserId,
		Rate:        dataModel.Rate,
		Description: dataModel.Description,
	}
}

func ListHomeCommentToCore(dataModels []HomeComment) []comment.Core {
	var dataCore []comment.Core
	for _, value := range dataModels {
		dataCore = append(dataCore, value.HomeCommentToCore())
	}
	return dataCore
}
