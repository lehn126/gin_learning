package conf

import (
	"fmt"
	"gin_app/app/util/file"

	"gopkg.in/yaml.v2"
)

// Fill yaml instances with the content from the yml config file
func LoadConfFromYml(filePath string, confInstancePointer any) any {
	ptr := file.ReadFile(filePath)
	e := yaml.Unmarshal([]byte(*ptr), confInstancePointer)
	if e != nil {
		panic(fmt.Sprintf("meet error when unmarshal app config file %s, error: %v", filePath, e))
	}
	return confInstancePointer
}
