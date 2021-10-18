package main

import (
	"fmt"
	"github.com/IvanBychkov27/neuralnet/internal/neuralnet"
	"github.com/IvanBychkov27/neuralnet/internal/operation"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error init zap logger, %v \n", err)
		return
	}
	// создаем и тренируем нейросеть
	//createAndTrainNeuralNet(logger)

	// загружаем нейросеть из файла и получаем результат
	loadNeuralNet(logger)
}

// createAndTrainNeuralNet - создает и обучает нейронную сеть
func createAndTrainNeuralNet(logger *zap.Logger) {
	// входные параметры нейронной сети
	inputCount := 26  // кол-во входных нейронов
	hiddenCount := 78 // кол-во внутренних (скрытых) нейронов
	outputCount := 17 // кол-во выходных нейронов
	rate1 := 0.25     // скорость обучения 1
	rate2 := 0.1      // скорость обучения 2

	countPlNameForTrain := 100 // кол-во записей каждой платформы в обучающих данных - ДАННЫЕ ДЛЯ ОБУЧЕНИЯ (примерно n*12) (78 - 996)
	iteration := 100000        // кол-во итераций при обучении

	net := neuralnet.NewNeuralNet(logger, inputCount, hiddenCount, outputCount, false, rate1, rate2)

	net.CreateNeuralNet()

	// файл с большими входными данными
	fileNameStampBigDate := "data/bigdata/posfdata25ok.txt" // 289938 lines
	input, target := net.BuildTrainDataForNeuralNet(fileNameStampBigDate, countPlNameForTrain)

	logger.Debug("start training", zap.Int("iteration", iteration))
	net.TrainNeuralNet(input, target, iteration)

	// сохраняем файл с нейронной сетью
	dataFileName := fmt.Sprintf("%d_%d_%d_%d", inputCount, hiddenCount, outputCount, len(input))
	fileNameNeuralNet := "data/nn/gonn_" + dataFileName
	net.SaveNeuralNet(fileNameNeuralNet)

	logger.Debug("neuralnet saved", zap.String("file name", fileNameNeuralNet))
}

// loadNeuralNet - загружаем нейронную сеть и получаем результат работы нейросети
func loadNeuralNet(logger *zap.Logger) {
	net := operation.WorkNeuralNet(logger)

	//fileNameNeuralNet := "data/nn/gonn_26_78_17_170"
	fileNameNeuralNet := "data/nn/gonn_26_78_17_1257"
	net.LoadNN(fileNameNeuralNet)

	data, expected := []float64{0.65535, 0.64, 0.9, 0.2, 0.53, 0.280, 0.78, 0.141, 0.78, 0.78, 0.83, 0.69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.1460}, "iOS"
	//data, expected := []float64{0.65535,0.60,0.9,0.2,0.54,0.274,0.83,0.78,0.144,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1400}, "Android"
	//data, expected := []float64{0.64240,0.52,0.9,0.2,0.119,0.280,0.78,0.143,0.78,0.78,0.83,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "Windows"
	//data, expected := []float64{0.65535,0.64,0.9,0.2,0.51,0.280,0.78,0.141,0.78,0.78,0.83,0.69,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "macOS"
	//data, expected := []float64{0.65535,0.64,0.9,0.2,0.50,0.277,0.78,0.142,0.78,0.78,0.83,0.69,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1412}, "iPadOS"
	//data, expected := []float64{0.64240,0.60,0.9,0.2,0.49,0.280,0.83,0.78,0.142,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "Linux"
	//data, expected := []float64{0.64240, 0.60, 0.9, 0.2, 0.57, 0.278, 0.83, 0.78, 0.142, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.1440}, "Linux"
	//data, expected := []float64{0.29200,0.60,0.9,0.2,0.44,0.280,0.83,0.78,0.142,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "LinuxChrome OS"
	//data, expected := []float64{0.65535,0.60,0.9,0.2,0.47,0.280,0.78,0.142,0.83,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "PlayStation 4"
	//data, expected := []float64{0.29200,0.60,0.9,0.2,0.54,0.279,0.83,0.78,0.142,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1360}, "Tizen"
	//data, expected := []float64{0.29200,0.60,0.9,0.2,0.58,0.280,0.83,0.78,0.143,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "Darwin"
	//data, expected := []float64{0.29200,0.60,0.9,0.2,0.53,0.280,0.83,0.78,0.142,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "NetCast"
	//data, expected := []float64{0.27760,0.60,0.9,0.2,0.46,0.285,0.83,0.78,0.143,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1339}, "KAIOS"
	//data, expected := []float64{0.65535,0.52,0.9,0.2,0.113,0.281,0.78,0.143,0.78,0.78,0.83,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1452}, "Windows Phone"
	//data, expected := []float64{0.14600,0.60,0.9,0.2,0.59,0.280,0.83,0.78,0.142,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1460}, "SmartTV"
	//data, expected := []float64{0.14600,0.60,0.9,0.2,0.45,0.274,0.83,0.78,0.141,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1400}, "FreeBSD"
	//data, expected := []float64{0.65535,0.60,0.9,0.2,0.45,0.289,0.83,0.78,0.144,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1388}, "BlackBerry"
	//data, expected := []float64{0.65535,0.52,0.9,0.2,0.115,0.279,0.78,0.143,0.78,0.78,0.83,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0.1360}, "Trident"

	out := net.GetResultFromNN(data)
	fmt.Println(net.ResultPlName(out, expected))
}
