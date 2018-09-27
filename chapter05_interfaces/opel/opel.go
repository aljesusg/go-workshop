package opel

type Opel struct {
	name     string
	price    float64
	discount float64
	color    string
}

// No name Method
/*
func (t Opel) Name() string {
	return t.name
}
*/
func (t Opel) Price() float64 {
	return t.price * t.discount
}

func (t Opel) Discount() float64 {
	return t.discount
}

func (t Opel) Color() string {
	return t.color
}

// NewCar contructure
func NewCar(name string, price float64, discount float64, color string) *Opel {
	return &Opel{
		name:     name,
		price:    price,
		discount: discount,
		color:    color,
	}
}
