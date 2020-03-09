package model

import (
	"github.com/oelmekki/matrix"
)

type AHPMatrix struct {
	m matrix.Matrix
}

func NewAHPMatrix(values [][]float64) AHPMatrix {
	ret := AHPMatrix{
		m: matrix.GenerateMatrix(len(values), len(values[0])),
	}

	for i, row := range values {
		for j, f := range row {
			ret.m.SetAt(i, j, f)
		}
	}

	return ret
}

func (ahp AHPMatrix) summarizeColumn(j int) float64 {
	sum := 0.0
	for i := 0; i < ahp.m.Rows(); i++ {
		sum += ahp.m.At(i, j)
	}
	return sum
}

func (ahp *AHPMatrix) Normalize() {
	for j := 0; j < ahp.m.Cols(); j++ {
		sum := ahp.summarizeColumn(j)
		for i := 0; i < ahp.m.Rows(); i++ {
			ahp.m.SetAt(i, j, ahp.m.At(i, j) / sum)
		}
	}
}

func (ahp *AHPMatrix) GetPriorityVector() []float64 {
	vector := []float64{}
	for i := 0; i < ahp.m.Rows(); i++ {
		s := 0.0
		for j := 0; j < ahp.m.Cols(); j++ {
			s += ahp.m.At(i, j)
		}
		vector = append(vector, s / float64(ahp.m.Cols()))
	}

	return vector
}
