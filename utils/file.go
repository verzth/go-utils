package utils

import (
	"io/ioutil"
	"os"
)

func movePath(src string, dst string)  {
	oFile, err := ioutil.ReadFile(src)
	if err == nil {
		os.MkdirAll(dst, os.ModeDir)
		nFile, err := os.Create(dst)
		if err != nil {
			nFile.Write(oFile)
		}
		os.Remove(src)
	}
}

func moveFile(src string, dst string, filename string)  {
	moveFileRename(src,dst,filename,filename)
}

func moveFileRename(src string, dst string, filename string, newname string)  {
	movePath(src+filename, dst+newname)
}