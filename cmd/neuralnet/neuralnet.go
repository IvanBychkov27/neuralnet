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
	inputCount := 7   // кол-во входных нейронов
	hiddenCount := 32 // кол-во внутренних (скрытых) нейронов
	outputCount := 17 // кол-во выходных нейронов
	rate1 := 0.25     // скорость обучения 1
	rate2 := 0.1      // скорость обучения 2

	countPlNameForTrain := 10 // кол-во записей каждой платформы в обучающих данных
	iteration := 100000       // кол-во итераций при обучении

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

	fileNameNeuralNet := "data/nn/gonn_7_32_17_170"
	net.LoadNN(fileNameNeuralNet)

	data := []float64{0.65535, 0.64, 0.9, 0.2, 0.53, 0.7327042439008, 0.1460} // iOS

	out := net.ResultFromNN(data)
	res := net.ResultPlName(out)

	fmt.Println(res)
}
