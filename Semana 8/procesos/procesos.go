package procesos

import (
	"fmt"
	"time"
)


type Proceso struct{
	Estado bool
	Id int
	Show bool
}

func (p *Proceso) start() {
	i := 0
	for {
		if p.Estado == false{
			return
		}
		if p.Show == true {
			fmt.Println("ID: ", p.Id, ": ", i)
		}
		time.Sleep(time.Millisecond * 500)
	}
	
}

func (p *Proceso) show() {
	p.Show = !p.Show
}


func (p *Proceso) stop() {
	p.Estado = false
}