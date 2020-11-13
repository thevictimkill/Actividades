package main

import (
  "fmt"
  "./figuras"
)

func sumArea(figs ...figuras.Figura) float64 {
	area_total := 0.0
	for _, f := range figs {
		area_total += f.Area()
	}
	return area_total
}

func main() {
  c01 := figuras.Circulo{6}
  fmt.Println(c01.Area())
  
  r01 := figuras.Rectangulo{10, 2}
  fmt.Println(r01.Area())

  fmt.Println(sumArea(&c01, &r01))

  fm := figuras.FiguraMultiple{
	  Figuras: []figuras.Figura{
		  &c01,
		  &r01,
		  &figuras.Circulo{2},
		},
	}
	

	fmt.Println(fm.Area())

	fm.Add(&c01)
	fmt.Println(fm.Area())

}