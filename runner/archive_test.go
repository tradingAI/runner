package runner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArchive(t *testing.T) {
	testDir := "testdata/achive"
	err := Archive(testDir)
	destPath := fmt.Sprintf("%s.%s", testDir, ARCHIVE_EXT)
	defer DeleteFile(destPath)
	assert.Nil(t, err)
}


func TestUnarchive(t *testing.T) {
	testPath := "testdata/unachive.zip"
	err := Unarchive(testPath)
	defer DeleteDirectory("testdata/unachive")
	assert.Nil(t, err)
}
