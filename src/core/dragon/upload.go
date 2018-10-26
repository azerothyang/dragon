package dragon

import (
	"io"
	"log"
	"net/http"
	"os"
)

//suggest OSS (object storage service). file upload
func Upload(r *http.Request, file string, saveTo string) error {
	fi, header, err := r.FormFile(file)
	if err != nil {
		log.Println(err)
		return err
	}
	srcFile, err := header.Open()
	defer srcFile.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	dstFile, err := os.Create(saveTo)
	defer dstFile.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}