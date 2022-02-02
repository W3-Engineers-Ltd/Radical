package dev

import (
	"os"

	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
)

var preCommit = `
goimports -w -format-only ./ \
ineffassign . \
staticcheck -show-ignored -checks "-ST1017,-U1000,-ST1005,-S1034,-S1012,-SA4006,-SA6005,-SA1019,-SA1024" ./ \
`

// for now, we simply override pre-commit file
func initGitHook() {
	// pcf => pre-commit file
	pcfPath := "./.git/hooks/pre-commit"
	pcf, err := os.OpenFile(pcfPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		radicalLogger.Log.Errorf("try to create or open file failed: %s, cause: %s", pcfPath, err.Error())
		return
	}

	defer pcf.Close()
	_, err = pcf.Write(([]byte)(preCommit))

	if err != nil {
		radicalLogger.Log.Errorf("could not init githooks: %s", err.Error())
	} else {
		radicalLogger.Log.Successf("The githooks has radicaln added, the content is:\n %s ", preCommit)
	}
}
