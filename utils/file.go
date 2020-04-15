package utils

import (
	"io/ioutil"
	"os"
)

func movePath(src string, dst string, perm os.FileMode)  {
	oFile, err := ioutil.ReadFile(src)
	if err == nil {
		os.MkdirAll(dst, perm)
		nFile, err := os.Create(dst)
		if err != nil {
			nFile.Write(oFile)
		}
		os.Remove(src)
	}
}

func moveFile(src string, dst string, filename string, perm os.FileMode)  {
	moveFileRename(src,dst,filename,filename, perm)
}

func moveFileRename(src string, dst string, filename string, newname string, perm os.FileMode)  {
	movePath(src+filename, dst+newname, perm)
}