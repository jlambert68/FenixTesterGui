package importFilesFromGitHub

import (
	sharedCode "FenixTesterGui/common_code"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"strings"
)

// Create Domain-Select-DropDown
func (importFilesFromGitHubObject *ImportFilesFromGitHubStruct) generateGitHubRepositorySelect(
	githubRepositoryUrls []string,
	templateRepositoryApiUrls []*fenixGuiTestCaseBuilderServerGrpcApi.RepositoryApiUrlResponseMessage) {

	importFilesFromGitHubObject.githubRepositorySelect = widget.NewSelect(
		githubRepositoryUrls,
		func(selected string) {

			// Loop RepositoryApiUrlResponseMessages to find correct 'GitHubApiKey'
			var indexForGithubRepositoryUrl int
			var foundGithubRepositoryUrl bool
			for selectIndex, githubRepositoryUrl := range githubRepositoryUrls {
				if selected == githubRepositoryUrl {
					foundGithubRepositoryUrl = true
					indexForGithubRepositoryUrl = selectIndex
					break
				}
			}

			// Shouldn't happen
			if foundGithubRepositoryUrl == false {
				sharedCode.Logger.WithFields(logrus.Fields{
					"ID":       "a4f01d73-80b4-400f-a45d-7c2e5026a909",
					"selected": selected,
				}).Error("Didn't find the correct URL in list with repository URLs")

				selected = "Didn't find the correct URL in list with repository URLs"
			} else {
				importFilesFromGitHubObject.currentGitHubApiKey = templateRepositoryApiUrls[indexForGithubRepositoryUrl].GetGitHubApiKey()
			}

			importFilesFromGitHubObject.currentPathShowedinGUI.Set(strings.Split(selected, "?")[0])
			importFilesFromGitHubObject.rootApiUrl = selected
			importFilesFromGitHubObject.getFileListFromGitHub(selected)
			importFilesFromGitHubObject.filterFileListFromGitHub()
			importFilesFromGitHubObject.filteredFileList.Refresh() // Refresh the list to update it with the new contents
		})

	importFilesFromGitHubObject.githubRepositorySelect.Refresh()
}
