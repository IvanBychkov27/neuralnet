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

	run(logger)

	logger.Debug("Done...")
}

func run(logger *zap.Logger) {
	// входные параметры нейронной сети
	inputCount := 7   // кол-во входных нейронов
	hiddenCount := 32 // кол-во внутренних (скрытых) нейронов
	outputCount := 17 // кол-во выходных нейронов
	rate1 := 0.25     // скорость обучения 1
	rate2 := 0.1      // скорость обучения 2

	net := neuralnet.NewNeuralNet(logger, inputCount, hiddenCount, outputCount, false, rate1, rate2)

	net.CreateNeuralNet()

}
