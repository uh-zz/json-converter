package main

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
	"sync"

	// "github.com/uh-zz/tagscanner/examples/jsondecoder/runtimescan"
	"gitlab.com/osaki-lab/tagscanner/runtimescan"
)

type bodyType int

const (
	bodyUnread bodyType = iota
	bodyUrlEncoding
	bodyMultipart
	bodyJson
	bodyUnknown
)

type decoder struct {
	src      string
	once     *sync.Once
	bodyType bodyType
}

func (m *decoder) ParseTag(name, tagKey, tagStr, pathStr string, elemType reflect.Type) (tag interface{}, err error) {
	return ParseJsonTag(name, tagStr, pathStr, elemType)
}

func (m *decoder) ExtractValue(tagInstance interface{}) (value interface{}, err error) {
	t := tagInstance.(*JsonTag)

	switch t.Type {
	case JsonField:
		tValue, _ := m.existKey(t.Name)

		return tValue, nil
	case IgnoreField:
		return nil, runtimescan.Skip
	}
	return nil, runtimescan.Skip
}

// Decode json文字列から構造体へアサインする
func Decode(dest interface{}, r string) error {
	decoder := &decoder{
		src:      r,
		once:     &sync.Once{},
		bodyType: bodyUnread,
	}
	return runtimescan.Decode(dest, []string{"json-converter"}, decoder)
}

// existKey json文字列にタグフィールドで指定されたキーがあれば、値を返す
func (m *decoder) existKey(key string) (value interface{}, err error) {

	str := strings.Trim(m.src, "{}")

	spslice := strings.Split(str, ",")

	for _, v := range spslice {
		rep := regexp.MustCompile(`\s*:\s*`)
		result := rep.Split(v, -1)
		if strings.Trim(result[0], "\"") == key {
			return strings.Trim(result[1], "\""), nil
		}
	}
	return nil, errors.New("key does not exist")
}
