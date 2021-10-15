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
}

// ResultFromNN - получаем результат от нейросети
func (w *WorkNN) ResultFromNN(data []float64) []float64 {
	return w.nn.Forward(data)
}

// ResultPlName - возвращает ответ от нейросети в виде строки
func (w *WorkNN) ResultPlName(data []float64, expected string) string {
	return w.buildResult(data, expected)
}
