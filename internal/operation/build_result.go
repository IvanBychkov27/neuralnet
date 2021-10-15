package operation

import (
	"fmt"
	"sort"
)

// buildResult - возвращает ответ от нейросети в виде строки
func (w *WorkNN) buildResult(output []float64) string {
	type tempData struct {
		plName string
		score  float64
	}

	resData := make([]tempData, 0)
	for pos, val := range output {
		if val < 0.00001 {
			continue
		}
		d := tempData{score: val * 100}
		switch pos {
		case 0:
			d.plName = "Android"
		case 1:
			d.plName = "iOS"
		case 2:
			d.plName = "Windows"
		case 3:
			d.plName = "macOS"
		case 4:
			d.plName = "iPadOS"
		case 5:
			d.plName = "Linux"
		case 6:
			d.plName = "LinuxChrome OS"
		case 7:
			d.plName = "PlayStation 4"
		case 8:
			d.plName = "Tizen"
		case 9:
			d.plName = "Darwin"
		case 10:
			d.plName = "NetCast"
		case 11:
			d.plName = "KAIOS"
		case 12:
			d.plName = "Windows Phone"
		case 13:
			d.plName = "SmartTV"
		case 14:
			d.plName = "FreeBSD"
		case 15:
			d.plName = "BlackBerry"
		case 16:
			d.plName = "Trident"
		default:
			d.plName = "Unknown"
		}

		resData = append(resData, d)
	}

	sort.SliceStable(resData, func(i, j int) bool {
		return resData[i].score > resData[j].score // сортировка по убыванию
	})

	res := ""
	for i, d := range resData {
		res += fmt.Sprintf("%d: %.3f  %s\n", i, d.score, d.plName)
	}

	return res
}

// buildResult - возвращает ответ от нейросети в виде строки
//func (w *WorkNN) buildResult1(output []float64) string {
//	max := float64(-99999)
//	pos := -1
//	// Ищем позицию нейрона с самым большим весом
//	for i, value := range output {
//		if value > max {
//			max = value
//			pos = i
//		}
//	}
//
//	res := ""
//
//	res += fmt.Sprintf(" 0: %f Android \n 1: %f iOS \n 2: %f Windows \n 3: %f macOS \n 4: %f iPadOS \n 5: %f Linux \n 6: %f LinuxChrome OS \n 7: %f PlayStation 4 \n 8: %f Tizen \n 9: %f Darwin \n10: %f NetCast\n11: %f KAIOS\n12: %f Windows Phone\n13: %f SmartTV\n14: %f FreeBSD\n15: %f BlackBerry\n16: %f Trident\n",
//		output[0], output[1], output[2], output[3], output[4], output[5], output[6],
//		output[7], output[8], output[9], output[10], output[11],
//		output[12], output[13], output[14], output[15], output[16],
//	)
//	res += "result: "
//
//	// Теперь, в зависимости от позиции, возвращаем решение
//	switch pos {
//	case 0:
//		res += "Android"
//	case 1:
//		res += "iOS"
//	case 2:
//		res += "Windows"
//	case 3:
//		res += "macOS"
//	case 4:
//		res += "iPadOS"
//	case 5:
//		res += "Linux"
//	case 6:
//		res += "LinuxChrome OS"
//	case 7:
//		res += "PlayStation 4"
//	case 8:
//		res += "Tizen"
//	case 9:
//		res += "Darwin"
//	case 10:
//		res += "NetCast"
//	case 11:
//		res += "KAIOS"
//	case 12:
//		res += "Windows Phone"
//	case 13:
//		res += "SmartTV"
//	case 14:
//		res += "FreeBSD"
//	case 15:
//		res += "BlackBerry"
//	case 16:
//		res += "Trident"
//	default:
//		res += "Unknown"
//	}
//	return res
//}
