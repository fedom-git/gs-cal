package util

import (
	"github.com/fedom-git/gs-cal/util/log"
	"gopkg.in/yaml.v3"
	"os"
)

func ParseFromYaml(path string, ptr interface{}) interface{} {
	//ptr := obj.(*GsObject)
	file, err := os.ReadFile(path)
	if err != nil {
		log.Log.Println("parse error: ", err.Error())
		return nil
	}
	err = yaml.Unmarshal(file, ptr)
	if err != nil {
		log.Log.Println("unmarshal error: ", err.Error())
		return nil
	}
	log.Log.Printf("parsed %s:\n%+v", path, ptr)
	return ptr
}
