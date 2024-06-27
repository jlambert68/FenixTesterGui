package importFilesFromGitHub

import (
	"encoding/base64"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

// Generate the button that imports the selected files from Github
func (importFilesFromGitHubObject ImportFilesFromGitHubStruct) generateImportSelectedFilesFromGithubButton(parentWindow fyne.Window) {

	importFilesFromGitHubObject.importSelectedFilesFromGithubButton = widget.NewButton("Import Files", func() {
		for fileIndex, file := range importFilesFromGitHubObject.selectedFiles {
			content, err := importFilesFromGitHubObject.loadFileContent(file)
			if err != nil {
				dialog.ShowError(err, parentWindow)
				continue
			}
			// Do something with the content, e.g., display it, process it, etc.
			importFilesFromGitHubObject.selectedFiles[fileIndex].Content = content

			extractedContent, err := importFilesFromGitHubObject.extractContentFromJson(string(content))
			if err != nil {
				log.Fatalf("Error parsing JSON: %s", err)
			}

			contentAsString, err := importFilesFromGitHubObject.decodeBase64Content(string(extractedContent))
			if err != nil {
				log.Fatalf("Failed to decode content: %s", err)
			}
			// 'content' now contains the decoded file content as a string
			//fmt.Println(contentAsString)

			// Save the file content
			importFilesFromGitHubObject.selectedFiles[fileIndex].FileContentAsString = contentAsString
		}

		fenixMainWindow.Show()
		parentWindow.Close()
	})
}

// Extra the file content from the json
func (importFilesFromGitHubObject ImportFilesFromGitHubStruct) extractContentFromJson(jsonData string) (string, error) {
	var fileDetail GitHubFileDetail
	err := json.Unmarshal([]byte(jsonData), &fileDetail)
	if err != nil {
		return "", err
	}

	return fileDetail.Content, nil
}

// Decode the file content from base64 to string
func (importFilesFromGitHubObject ImportFilesFromGitHubStruct) decodeBase64Content(encodedContent string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
