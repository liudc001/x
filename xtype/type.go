package xtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const (
	imgprefix  = "https://image.xuebaox.com/"
	fileprefix = "https://static.xuebaox.com/"
)

// ImageURL ===============图片链接===============
type ImageURL string

// String 转换为string类型
func (f ImageURL) String() string {
	s := string(f)
	var url = s
	if !strings.HasPrefix(s, "http") {
		url = imgprefix + s
	}
	return url
}

// IsEmpty 是否为空
func (f ImageURL) IsEmpty() bool {
	s := string(f)
	if s == "" {
		return true
	}
	return false
}

// MarshalJSON 转换为json类型 加域名
func (f ImageURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

// UnmarshalJSON 不做处理
func (f *ImageURL) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	tmp = strings.TrimPrefix(tmp, imgprefix)
	*f = ImageURL(tmp)
	return nil
}

// Scan implements the Scanner interface.
func (f *ImageURL) Scan(src interface{}) error {
	if src == nil {
		*f = ""
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("Read file url data from DB failed")
	}
	*f = ImageURL(tmp)
	return nil
}

// Value implements the driver Valuer interface.
func (f ImageURL) Value() (driver.Value, error) {
	return string(f), nil
}

// FileURL ================文件链接================
type FileURL string

// String 转换为string类型
func (f FileURL) String() string {
	s := string(f)
	var url = s
	if !strings.HasPrefix(s, "http") {
		url = fileprefix + s
	}
	return url
}

// IsEmpty 是否为空
func (f FileURL) IsEmpty() bool {
	s := string(f)
	if s == "" {
		return true
	}
	return false
}

// MarshalJSON 转换为json类型 加域名
func (f FileURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

// UnmarshalJSON 不做处理
func (f *FileURL) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	tmp = strings.TrimPrefix(tmp, fileprefix)
	*f = FileURL(tmp)
	return nil
}

// Scan implements the Scanner interface.
func (f *FileURL) Scan(src interface{}) error {
	if src == nil {
		*f = ""
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("Read file url data from DB failed")
	}
	*f = FileURL(tmp)
	return nil
}

// Value implements the driver Valuer interface.
func (f FileURL) Value() (driver.Value, error) {
	return string(f), nil
}

// Strings ===========字符串列表===========
type Strings []string

// String 转换为string类型
func (t Strings) String() string {
	var s = []string(t)
	return strings.Join(s, ",")
}

// MarshalJSON 转换为json类型
func (t Strings) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(t))
}

// UnmarshalJSON 不做处理
func (t *Strings) UnmarshalJSON(data []byte) error {
	var tmp []string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = Strings(tmp)
	return nil
}

// MarshalYAML 转换为json类型
func (t Strings) MarshalYAML() ([]byte, error) {
	return yaml.Marshal([]string(t))
}

// UnmarshalYAML 不做处理
func (t *Strings) UnmarshalYAML(data []byte) error {
	var tmp []string
	if err := yaml.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = Strings(tmp)
	return nil
}

// Scan implements the Scanner interface.
func (t *Strings) Scan(src interface{}) error {
	*t = make([]string, 0)
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("Read tags from DB failed")
	}
	if len(tmp) == 0 {
		return nil
	}
	*t = strings.Split(string(tmp), ",")
	return nil
}

// Value implements the driver Valuer interface.
func (t Strings) Value() (driver.Value, error) {
	return t.String(), nil
}

// Numbers ===========数字列表===========
type Numbers []int

// String 转换为string类型
func (t Numbers) String() string {
	var s = []int(t)
	var ns []string
	for _, i := range s {
		ns = append(ns, strconv.Itoa(i))
	}
	return strings.Join(ns, ",")
}

// MarshalJSON 转换为json类型 加域名
func (t Numbers) MarshalJSON() ([]byte, error) {
	return json.Marshal([]int(t))
}

// UnmarshalJSON 不做处理
func (t *Numbers) UnmarshalJSON(data []byte) error {
	var tmp []int
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = Numbers(tmp)
	return nil
}

// Scan implements the Scanner interface.
func (t *Numbers) Scan(src interface{}) error {
	*t = make([]int, 0)
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("Read tags from DB failed")
	}
	if len(tmp) == 0 {
		return nil
	}
	ts := strings.Split(string(tmp), ",")
	for _, i := range ts {
		n, _ := strconv.Atoi(i)
		*t = append(*t, n)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (t Numbers) Value() (driver.Value, error) {
	return t.String(), nil
}