package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArchive(t *testing.T) {
	testDir := "testdata/achive"
	destPath, err := Archive(testDir)
	assert.Nil(t, err)
	assert.Equal(t, "testdata/achive.zip", destPath)
	err = DeleteFile(destPath)
	assert.Nil(t, err)
}


func TestUnarchive(t *testing.T) {
	srcPath := "testdata/unachive.zip"
	destDir := "testdata/unachive"
	err := Unarchive(srcPath, destDir)
	assert.Nil(t, err)
	err = DeleteDirectory(destDir)
	assert.Nil(t, err)
}
