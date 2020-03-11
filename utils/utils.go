package utils

func Substring(text string, indexs... int) string {
	runes := []rune(text)
	if len(indexs) > 0 {
		if indexs[0] < 0 || indexs[0] > len(runes) {
			return ""
		}
		if len(indexs) > 1 {
			if indexs[0] >= indexs[1] {
				return ""
			}
			if indexs[1] <= len(runes) {
				return string(runes[indexs[0]:indexs[1]])
			}
		}
		return string(runes[indexs[0]:])
	}else{
		return text
	}
}

func FileMove(src string, dst string, file... string)  {
	pCount := len(file)
	if pCount == 0 {
		movePath(src,dst)
	}else if pCount == 1 {
		moveFile(src,dst,file[0])
	}else{
		moveFileRename(src,dst,file[0],file[1])
	}
}