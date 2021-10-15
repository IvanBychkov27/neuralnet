package neuralnet

import (
	"github.com/fxsjy/gonn/gonn"
	"go.uber.org/zap"
)

type NN struct {
	logger      *zap.Logger
	nn          *gonn.NeuralNetwork
	inputCount  int
	hiddenCount int
	outputCount int
	regression  bool
	rate1       float64
	rate2       float64
}

func NewNeuralNet(logger *zap.Logger, inputCount, hiddenCount, outputCount int, regression bool, rate1, rate2 float64) *NN {
	return &NN{
		logger:      logger,
		inputCount:  inputCount,
		hiddenCount: hiddenCount,
		outputCount: outputCount,
		regression:  regression,
		rate1:       rate1,
		rate2:       rate2,
	}
}

// CreateNeuralNet - Создаем нейросеть с Параметрами нейросети: кол-во нейронов (входных, внутренних, выходных, регрессия, rate1, rate2)
func (net *NN) CreateNeuralNet() {
	net.nn = gonn.NewNetwork(net.inputCount, net.hiddenCount, net.outputCount, net.regression, net.rate1, net.rate2)
}

// BuildTrainDataForNeuralNet - строит входные и выходные данные для нейронной сети
func (net *NN) BuildTrainDataForNeuralNet(fileName string, countPlNameForTrain int) (input, target [][]float64) {
	return net.processingBigData(fileName, countPlNameForTrain)
}

// TrainNeuralNet - Обучаем нейросеть
func (net *NN) TrainNeuralNet(input, target [][]float64, iteration int) {
	net.nn.Train(input, target, iteration)
}

// SaveNeuralNet - Сохраняем нейросеть
func (net *NN) SaveNeuralNet(fileName string) {
	gonn.DumpNN(fileName, net.nn)
}
