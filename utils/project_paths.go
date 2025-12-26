package utils

import "github.com/u00io/gazerstudio/localstorage"

func ProjectDirPath(projectId string) string {
	gazerStudioPath := localstorage.Path()
	return gazerStudioPath + "/projects/" + projectId
}

func ProjectDataItemsDirPath(projectId string) string {
	gazerStudioPath := localstorage.Path()
	return gazerStudioPath + "/projects/" + projectId + "/data_items"
}

func ProjectDataItemDirPath(projectId string, dataItemId string) string {
	gazerStudioPath := localstorage.Path()
	return gazerStudioPath + "/projects/" + projectId + "/data_items/" + dataItemId
}
