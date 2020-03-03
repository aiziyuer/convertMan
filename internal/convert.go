package internal

import (
	"errors"
	"github.com/gogf/gf/encoding/gparser"
	"github.com/gogf/gf/os/gfile"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

func ConvertFiles(iFormat string, oFormat string, paths []string) (string, error) {

	switch iFormat {
	case "yaml":

		var sb strings.Builder

		// if all file is yaml, the yaml can be combine

		for _, path := range paths {

			if iFormat == gfile.Ext(path) {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					return "", err
				}
				sb.WriteString(string(content))
			}
		}

		if sb.Len() == 0 {
			return "", errors.New("no content detected, empty files? ")
		}

		p, err := gparser.LoadContent(sb.String())
		if err != nil {
			return "", err
		}

		return reduce(p, oFormat)

	default:
		p, err := gparser.Load(paths[0])
		if err != nil {
			return "", err
		}
		return reduce(p, oFormat)
	}

}

func ConvertContent(iFormat string, oFormat string, content string) (string, error) {

	var p *gparser.Parser
	var err error

	switch iFormat {
	case "yaml":
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
	case "yaml":
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
