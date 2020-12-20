package utils

//判断元素是否存在slice中
func IsExistsElementInt64(s []int64, e int64) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func IsExistsElementInt(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
