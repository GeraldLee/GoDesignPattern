package factory_method

type Car interface {
	GetName() string
}

type Factory interface {
	GetCar() Car
}

type BenzCarFactory struct {
}

func (f BenzCarFactory) GetCar() Car{
	return &BenzCar{}
}


type BMWCarFactory struct {
}

func (f BMWCarFactory) GetCar() Car{
	return &BMWCar{}
}


type AudiCarFactory struct {
}

func (f AudiCarFactory) GetCar() Car{
	return &AudiCar{}
}

type BenzCar struct {
}

func (c *BenzCar) GetName() string {
	return "Benz"
}

type BMWCar struct {
}

func (c *BMWCar) GetName() string {
	return "BMW"
}

type AudiCar struct {
}

func (c *AudiCar) GetName() string {
	return "Audi"
}
