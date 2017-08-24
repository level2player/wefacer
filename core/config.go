package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	WefacerConfig Config
)

const middle = "========="

type Config struct {
	ConfigMap map[string]string
	strcet    string
}

func init() {
	config, err := GetiniConfig("config.ini")
	if err != nil {
		fmt.Println("config.ini initialization error")
		return
	}
	WefacerConfig = config
}

func GetiniConfig(path string) (Config, error) {
	c := Config{}
	c.ConfigMap = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		return c, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return c, err
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := frist
		c.ConfigMap[key] = strings.TrimSpace(second)
	}
	return c, nil
}

func (c Config) Read(node, key string) string {
	key = node + middle + key
	v, found := c.ConfigMap[key]
	if !found {
		return ""
	}
	return v
}
