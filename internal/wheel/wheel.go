package wheel

type Wheel struct {
	Num_of_cells int8
	Cell         Cell
}

func (p Wheel) Rotation() {

}

type Cell struct {
	Name   string
	Image  string
	Chance int8
}
