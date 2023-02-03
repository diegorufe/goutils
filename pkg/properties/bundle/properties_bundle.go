package bundle

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/diegorufe/goutils/pkg/file/utils"
	"github.com/diegorufe/goutils/pkg/properties/structs"
	stringUtils "github.com/diegorufe/goutils/pkg/strings/utils"
)

type PropertiesBundle interface {
	LoadFiles(path string) error
	LoadFile(path string) error
	Get(value string) string
	Set(key string, value string)
}

type DefaultPropertiesBundle struct {
	Properties structs.Properties
}

func (ptBundle DefaultPropertiesBundle) setKeyLoadFile(key string, value string) {
	if key != "" {
		value = strings.TrimSpace(strings.ReplaceAll(value, "\\", ""))
		ptBundle.Set(key, value)
	}
}

func InstancePropertiesBundle() PropertiesBundle {
	ptBundle := DefaultPropertiesBundle{Properties: structs.Properties{}}
	ptBundle.Properties.Values = make(map[string]string)
	return ptBundle
}

func (ptBundle DefaultPropertiesBundle) LoadFiles(path string) error {
	var err error

	if utils.DirExists(path) {
		files, err := ioutil.ReadDir(path)
		if err == nil {
			for _, file := range files {
				if !file.IsDir() {
					err = ptBundle.LoadFile(path + "/" + file.Name())
					if err != nil {
						break
					}
				}
			}
		}

	} else {
		err = errors.New("Dir " + path + " dont exists")
	}

	return err
}

func (ptBundle DefaultPropertiesBundle) LoadFile(path string) error {
	var err error

	if utils.FileExist(path) {
		readFile, err := os.Open(path)

		if err == nil {
			fileScanner := bufio.NewScanner(readFile)

			fileScanner.Split(bufio.ScanLines)
			var key string = ""
			var value string = ""

			for fileScanner.Scan() {
				text := fileScanner.Text()

				// Ignore comments
				if !stringUtils.StartWith(text, "#") {

					// If firts value si blank, is a value
					if string(text[0]) == " " {
						value = value + " " + strings.TrimSpace(strings.ReplaceAll(text, "\\", ""))
					} else if stringUtils.Contains(text, "=") {
						ptBundle.setKeyLoadFile(key, value)
						indexEqual := strings.Index(text, "=")
						textLenth := len(text)
						key = text[0:indexEqual]
						value = text[indexEqual+1 : textLenth]
					}

				} else {
					ptBundle.setKeyLoadFile(key, value)
					key = ""
					value = ""
				}

			}

			ptBundle.setKeyLoadFile(key, value)

			// for key, element := range ptBundle.Properties.Values {
			// 	fmt.Println("Key:", key, "=>", "Element:", element)
			// }
		}

	} else {
		err = errors.New("File " + path + " dont exist")
	}

	return err
}

func (ptBundle DefaultPropertiesBundle) Get(key string) string {
	var result string = key
	val, ok := ptBundle.Properties.Values[key]

	if ok {
		result = val
	}

	return result
}

func (ptBundle DefaultPropertiesBundle) Set(key string, value string) {
	ptBundle.Properties.Values[key] = value
}
