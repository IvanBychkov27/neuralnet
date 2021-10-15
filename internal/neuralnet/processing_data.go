package neuralnet

import (
	"encoding/csv"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
)

func (net *NN) processingBigData(fileName string, countPlNameForTrain int) (input, target [][]float64) {
	stamps := net.getStampsDataFromFile(fileName)
	net.logger.Debug("getStampsDataFromFile", zap.Int("stamps", len(stamps)))

	dataTrain, countLines := net.buildTrainData(stamps, countPlNameForTrain)

	dataFileName := fmt.Sprintf("%d_%d_%d_%d", net.inputCount, net.hiddenCount, net.outputCount, countLines)
	fileNameSave := "data/traindata/train_" + dataFileName + ".csv"
	net.saveTrainDataFile(fileNameSave, dataTrain)

	input, target = net.getTrainDataForNeuralNet(fileNameSave)

	return
}

func (net *NN) buildTrainData(stamps []stampDataFromFile, countPlNameForTrain int) ([]byte, int) {
	var count, countAnd, countIOS, countWin, countMacOS, countIPadOS,
		countLinux, countLinuxChromeOS, countPlayStation4, countTizen,
		countDarwin, countNetCast, countKAIOS, countWindowsPhone,
		countSmartTV, countFreeBSD, countBlackBerry, countTrident int

	var res string

	// 26 входных и 17 выходных нейронов для нейронной сети
	res = "TCPWindowSize,TCPHeaderLength,IPFlags,TCPFlags,IPTTL,TCPOptions1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,MSS," +
		"Android,iOS,Windows,macOS,iPadOS,Linux,LinuxChromeOS,PlayStation4,Tizen,Darwin,NetCast,KAIOS," +
		"WindowsPhone,SmartTV,FreeBSD,BlackBerry,Trident\n"

	countPlName := countPlNameForTrain // количество каждой plName для тренировки нейронной сети
	for _, st := range stamps {
		platform := ""
		switch st.plName {
		case "Android":
			platform = "1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0" // 1
			if countAnd > countPlName {
				continue
			}
			countAnd++
		case "iOS":
			platform = "0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0" // 2
			if countIOS > countPlName {
				continue
			}
			countIOS++
		case "Windows":
			platform = "0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0" // 3
			if countWin > countPlName {
				continue
			}
			countWin++
		case "macOS":
			platform = "0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0" // 4
			if countMacOS > countPlName {
				continue
			}
			countMacOS++
		case "iPadOS":
			platform = "0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0" // 5
			if countIPadOS > countPlName {
				continue
			}
			countIPadOS++
		case "Linux":
			platform = "0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0" // 6
			if countLinux > countPlName {
				continue
			}
			countLinux++

		case "LinuxChrome OS":
			platform = "0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0" // 7
			if countLinuxChromeOS > countPlName {
				continue
			}
			countLinuxChromeOS++
		case "PlayStation 4":
			platform = "0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0" // 8
			if countPlayStation4 > countPlName {
				continue
			}
			countPlayStation4++
		case "Tizen":
			platform = "0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0" // 9
			if countTizen > countPlName {
				continue
			}
			countTizen++
		case "Darwin":
			platform = "0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0" // 10
			if countDarwin > countPlName {
				continue
			}
			countDarwin++
		case "NetCast":
			platform = "0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0" // 11
			if countNetCast > countPlName {
				continue
			}
			countNetCast++
		case "KAIOS":
			platform = "0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0" // 12
			if countKAIOS > countPlName {
				continue
			}
			countKAIOS++
		case "Windows Phone":
			platform = "0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0" // 13
			if countWindowsPhone > countPlName {
				continue
			}
			countWindowsPhone++
		case "SmartTV":
			platform = "0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0" // 14
			if countSmartTV > countPlName {
				continue
			}
			countSmartTV++
		case "FreeBSD":
			platform = "0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0" // 15
			if countFreeBSD > countPlName {
				continue
			}
			countFreeBSD++
		case "BlackBerry":
			platform = "0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0" // 16
			if countBlackBerry > countPlName {
				continue
			}
			countBlackBerry++
		case "Trident":
			platform = "0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1" // 17
			if countTrident > countPlName {
				continue
			}
			countTrident++

		default:
			continue
		}

		if countAnd > countPlName && countIOS > countPlName && countWin > countPlName && countMacOS > countPlName &&
			countIPadOS > countPlName && countLinux > countPlName && countLinuxChromeOS > countPlName && countPlayStation4 > countPlName &&
			countTizen > countPlName && countDarwin > countPlName && countNetCast > countPlName && countKAIOS > countPlName &&
			countWindowsPhone > countPlName && countSmartTV > countPlName && countFreeBSD > countPlName && countBlackBerry > countPlName &&
			countTrident > countPlName {
			break
		}

		stamp := parseStamp(st.stamp)

		// 7 входных и 17 выходных данных
		res += fmt.Sprintf("0.%s,0.%s,0.%s,0.%s,0.%s,%s,0.%s,%s\n",
			stamp.TCPWindowSize,   // 1
			stamp.TCPHeaderLength, // 2
			stamp.IPFlags,         // 3
			stamp.TCPFlags,        // 4
			stamp.IPTTL,           // 5
			stamp.TCPOptions,      // 6
			stamp.MSS,             // 7
			platform,
		)
		count++
	}
	return []byte(res), count
}

// saveTrainDataFile - сохраняет файл с тренировочными данными для нейронной сети
func (net *NN) saveTrainDataFile(fileName string, data []byte) {
	df, err := os.Create(fileName)
	if err != nil {
		net.logger.Error("error create train data file", zap.Error(err))
		return
	}
	defer df.Close()

	_, err = df.Write(data)
	if err != nil {
		net.logger.Error("error write data train file", zap.Error(err))
		return
	}
	net.logger.Debug("train data file saved", zap.String("file name", fileName))
}

// getTrainDataForNeuralNet - получаем входные и выходные данные для нейронной сети
func (net *NN) getTrainDataForNeuralNet(fileName string) (inputs, targets [][]float64) {
	file, err := os.Open(fileName)
	if err != nil {
		net.logger.Error("error open train data file", zap.Error(err))
		return
	}
	defer file.Close()

	inp := net.inputCount  // входных нейронов
	out := net.outputCount // выходных нейронов

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = inp + out // входных + выходных нейронов
	reader.Comment = '#'

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		net.logger.Error("error read data train file", zap.Error(err))
		return
	}

	for idx, record := range rawCSVData {

		inputsData := make([]float64, 0, inp) // входных нейронов
		labelsData := make([]float64, 0, out) //  выходных нейронов

		if idx == 0 || strings.Contains(record[0], "#") {
			continue
		}

		for i, val := range record {
			d, errPars := strconv.ParseFloat(val, 64)
			if errPars != nil {
				net.logger.Error("error parse float", zap.Error(err))
				return
			}
			if i < inp {
				inputsData = append(inputsData, d)
			} else {
				labelsData = append(labelsData, d)
			}
		}
		inputs = append(inputs, inputsData)
		targets = append(targets, labelsData)
	}

	return
}
