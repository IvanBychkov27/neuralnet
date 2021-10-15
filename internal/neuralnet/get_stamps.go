package neuralnet

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type stampDataFromFile struct {
	plName string
	stamp  string
}

func (net *NN) getStampsDataFromFile(fileName string) []stampDataFromFile {
	file, errOpen := ioutil.ReadFile(fileName)
	if errOpen != nil {
		fmt.Println("error open file stamp data", errOpen.Error())
		return nil
	}

	dataFile := strings.Split(string(file), "\n")
	fmt.Println("dataFile", len(dataFile))

	dStamp := make([]stampDataFromFile, 0, 150000)
	for _, lineFile := range dataFile {
		lf := strings.Split(lineFile, `";"`)
		if len(lf) != 9 {
			continue
		}
		platformName := strings.Trim(lf[3], " \"\n\t")
		stamp := strings.Trim(lf[8], " \"\n\t")

		if platformName == "" || stamp == "" {
			continue
		}

		d := stampDataFromFile{platformName, stamp}
		dStamp = append(dStamp, d)
	}
	return dStamp

}
