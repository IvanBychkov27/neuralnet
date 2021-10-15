package operation

import (
	"github.com/fxsjy/gonn/gonn"
	"go.uber.org/zap"
)

type WorkNN struct {
	logger *zap.Logger
	nn     *gonn.NeuralNetwork
}

func WorkNeuralNet(logger *zap.Logger) *WorkNN {
	return &WorkNN{
		logger: logger,
	}
}

// LoadNN - загрузка нейросети из файла
func (w *WorkNN) LoadNN(fileNameNN string) {
	w.nn = gonn.LoadNN(fileNameNN)
	//w.logger.Debug("neuralnet is loaded", zap.Int("input", len(w.nn.InputLayer)-1),zap.Int("output", len(w.nn.HiddenLayer)-1), zap.Int("output", len(w.nn.OutputLayer)))
}

// GetResultFromNN - получаем результат от нейросети
func (w *WorkNN) GetResultFromNN(data []float64) []float64 {
	inputNeurons := len(w.nn.InputLayer) - 1

	if len(data) < inputNeurons {
		w.logger.Debug("attention! not enough input data!")
		for len(data) < inputNeurons {
			data = append(data, 0)
		}
	}
	if len(data) > inputNeurons {
		w.logger.Debug("attention! lots of input data!")
		data = data[:inputNeurons]
	}

	return w.nn.Forward(data)
}

// ResultPlName - возвращает ответ от нейросети в виде строки
func (w *WorkNN) ResultPlName(data []float64, expected string) string {
	return w.buildResult(data, expected)
}
