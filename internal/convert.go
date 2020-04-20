package internal

import (
	"github.com/gogf/gf/encoding/gparser"
	"github.com/imdario/mergo"
	_ "github.com/imdario/mergo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

const FileFormatYaml = "yaml"

// ConvertFiles convert from files
func ConvertFiles(iFormat string, oFormat string, paths []string) (string, error) {

	switch iFormat {
	case FileFormatYaml:

		// if all file is yaml, the yaml can be combine
		if len(paths) > 1 {

			defaultObject, err := file2Map(paths[0])
			if err != nil {
				return "", err
			}

			for i := 1; i < len(paths); i++ {
				newObject, err := file2Map(paths[i])
				if err != nil {
					return "", err
				}
				err = mergo.Merge(&defaultObject, newObject,
					mergo.WithOverrideEmptySlice,
					mergo.WithOverride,
					mergo.WithOverwriteWithEmptyValue,
				)
				if err != nil {
					return "", err
				}
			}

			return gparser.VarToYamlString(defaultObject)

		} else {

			content, err := ioutil.ReadFile(paths[0])
			if err != nil {
				return "", err
			}

			if len(content) == 0 {
				return "", nil
			}

			p, err := gparser.LoadContent(content)
			if err != nil {
				return "", err
			}

			return reduce(p, oFormat)
		}

	default:
		p, err := gparser.Load(paths[0])
		if err != nil {
			return "", err
		}
		return reduce(p, oFormat)
	}

}

func file2Map(path string) (map[string]interface{}, error) {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	p, err := gparser.LoadContent(content)
	if err != nil {
		return nil, err
	}

	return p.ToMap(), nil
}

// ConvertContent convert content directly
func ConvertContent(iFormat string, oFormat string, content string) (string, error) {

	var p *gparser.Parser
	var err error

	switch iFormat {
	case FileFormatYaml:
		p, err = gparser.LoadYaml(content)
	case "json":
		p, err = gparser.LoadJson(content)
	case "ini":
		p, err = gparser.LoadIni(content)
	case "toml":
		p, err = gparser.LoadToml(content)
	case "xml":
		p, err = gparser.LoadXml(content)
	default:
		p, err = gparser.LoadContent(content)
	}

	if err != nil {
		logrus.Errorf("parse content failed: %s", err)
		return "", err
	}

	return reduce(p, oFormat)
}

func reduce(p *gparser.Parser, oFormat string) (string, error) {

	var ret string
	var err error

	switch oFormat {
	case FileFormatYaml:
		ret, err = p.ToYamlString()
	case "json":
		ret, err = p.ToJsonString()
	case "ini":
		ret, err = p.ToIniString()
	case "toml":
		ret, err = p.ToTomlString()
	case "xml":
		ret, err = p.ToXmlString()
	default:
		ret, err = p.ToJsonString()
	}

	return ret, err

}
