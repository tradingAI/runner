package runner

import (
	"os"
	"path"

	"github.com/golang/glog"
)

func mustCreatePath(filePath string) (err error) {
	dir := path.Dir(filePath)
	err = mustCreateDir(dir)
	if err != nil {
		glog.Error(err)
        return
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		glog.Error(err)
        return
	}
	f.Close()
	return
}

func mustCreateDir(dir string)(err error){
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			glog.Error(err)
			return err
		}
	}
	return err
}
