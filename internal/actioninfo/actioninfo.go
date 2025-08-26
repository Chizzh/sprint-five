package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("parsing error %s %v", v, err)
			continue
		}

		actioninfo, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error forming the string %v", err)
			continue
		}
		
		fmt.Println(actioninfo)
	}
}
