package mtools

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/axgle/mahonia"
)

//Convert2uft 中文转换为utf8..
func Convert2uft(srcCode string, tagCode string, tmp string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(tmp)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//DecideType 类型断言..
func DecideType(src interface{}) (string, error) {
	tmp := ""
	switch t := src.(type) {
	case nil:
		tmp = "null"
	case bool:
		if t {
			tmp = "True"
		} else {
			tmp = "False"
		}
	case []byte:
		tmp = string(t)
	case time.Time:
		tmp = t.Format("2006-01-02 15:04:05.999")
	case int:
		tmp = strconv.Itoa(t)
	case int32:
		tmp = strconv.Itoa(int(t))
	case int64:
		tmp = strconv.FormatInt(t, 10)
	case float32:
		tmp = strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		tmp = strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		tmp = string(t)
	default:
		err := errors.New("no this type")
		return tmp, err
	}
	return tmp, nil
}

//CheckErr 检查错误
func CheckErr(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
	}
}

//PathExists 目录判断
func PathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
