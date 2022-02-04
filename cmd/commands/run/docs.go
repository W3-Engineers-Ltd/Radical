package run

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strings"

	radicalLogger "github.com/W3-Engineers-Ltd/Radical/logger"
)

var (
	swaggerVersion = "3"
	swaggerlink    = "https://github.com/W3-Engineers-Ltd/radiant-swagger/archive/master.zip"
)

func downloadFromURL(url, fileName string) {
	var down bool
	if fd, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
		down = true
	} else if fd.Size() == int64(0) {
		down = true
	} else {
		radicalLogger.Log.Infof("'%s' already exists", fileName)
		return
	}
	if down {
		radicalLogger.Log.Infof("Downloading '%s' to '%s'...", url, fileName)
		output, err := os.Create(fileName)
		if err != nil {
			radicalLogger.Log.Errorf("Error while creating '%s': %s", fileName, err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			radicalLogger.Log.Errorf("Error while downloading '%s': %s", url, err)
			return
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			radicalLogger.Log.Errorf("Error while downloading '%s': %s", url, err)
			return
		}
		radicalLogger.Log.Successf("%d bytes downloaded!", n)
	}
}

func unzipAndDelete(src string) error {
	radicalLogger.Log.Infof("Unzipping '%s'...", src)
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	rp := strings.NewReplacer("swagger", "swagger")
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fname := rp.Replace(f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fname, f.Mode())
		} else {
			f, err := os.OpenFile(
				fname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	radicalLogger.Log.Successf("Done! Deleting '%s'...", src)
	return os.RemoveAll(src)
}
