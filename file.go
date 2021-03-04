package toolkit

import (
	"os"
	"path/filepath"
)

// CreateFilePath create all dir and file for given path
func CreateFilePath(path string) (err error) {
	p, f := filepath.Split(path)
	p = filepath.Clean(p)
	_, err = os.Stat(p)
	if os.IsNotExist(err) {
		err = os.MkdirAll(p, 0755)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	if f != "" {
		fp := filepath.Join(p, f)
		_, err = os.Create(fp)
		if err != nil {
			return err
		}
	}
	return nil
}
