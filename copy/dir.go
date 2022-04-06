package copy

import (
	"io/ioutil"
	"os"
	"path"
)

func Dir(src string, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}
	fds, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, fd := range fds {
		srcFp := path.Join(src, fd.Name())
		dstFp := path.Join(dst, fd.Name())
		if fd.IsDir() {
			if err = Dir(srcFp, dstFp); err != nil {
				return err
			}
		} else {
			if err = File(srcFp, dstFp); err != nil {
				return err
			}
		}
	}
	return nil
}
