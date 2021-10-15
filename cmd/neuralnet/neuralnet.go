package main

import (
	"fmt"
	"github.com/IvanBychkov27/neuralnet/internal/neuralnet"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error init zap logger, %v \n", err)
		return
	}
	logger.Debug("Start neuralnet")

	createAndTrainNeuralNet(logger)

	logger.Debug("Done...")
}

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
