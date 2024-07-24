package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//o t.run = usado para separar as coisas, ter duas funções para ser testada
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			//o fatal e error são parecidos, mas o fatal encerra assim que acha o erro, o error continua testando
			//t.Errorf("A area recebida %f é diferente da esperada %f",areaRecebida,areaEsperada)
			t.Fatalf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()
		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

	})
}
