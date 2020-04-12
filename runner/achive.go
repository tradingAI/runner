package runner

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/mholt/archiver/v3"
)

const ARCHIVE_EXT = "zip"

// NOTE(wen): 打包使用zip格式压缩
func Archive(srcDir string) (err error) {
	destPath := fmt.Sprintf("%s.%s", srcDir, ARCHIVE_EXT)
	err = archiver.Archive([]string{srcDir}, destPath)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

// NOTE(wen): 解压时按照文件名的扩展名的格式来判断文件格式
func Unarchive(srcPath string) (err error) {
	fileDir, fileName := path.Split(srcPath)
	dirName := strings.Split(fileName, ".")[0]
	destDir := path.Join(fileDir, dirName)
	err = archiver.Unarchive(srcPath, destDir)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

func DeleteFile(filePath string) (err error) {
	err = os.Remove(filePath)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteDirectory(dir string)(err error){
	err = os.RemoveAll(dir)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}
