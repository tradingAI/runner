package runner

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/mholt/archiver/v3"
)

const ARCHIVE_EXT = "zip"

// NOTE(wen): 打包使用zip格式压缩
func Archive(srcDir string) (destPath string, err error) {
	destPath = fmt.Sprintf("%s.%s", srcDir, ARCHIVE_EXT)
	err = archiver.Archive([]string{srcDir}, destPath)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

// NOTE(wen): 解压时按照文件名的扩展名的格式来判断文件格式
func Unarchive(srcPath, destDir string) (err error) {
	err = archiver.Unarchive(srcPath, destDir)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

func DeleteFile(filePath string) (err error) {
	glog.Infof("Runner DeleteFile: %s", filePath)
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		glog.Infof("Runner DeleteFile: %s is not exist!", filePath)
		return nil
	}
	err = os.Remove(filePath)
	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func DeleteDirectory(dir string) (err error) {
	glog.Infof("Runner DeleteDirectory: %s", dir)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		glog.Infof("Runner DeleteDirectory: %s is not exist!", dir)
		return nil
	}
	err = os.RemoveAll(dir)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}
