package simple_factory

type Car interface {
	GetName() string
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

type SimpleFactory struct {}

func (f SimpleFactory) NewCar(name string) Car {
	switch name {
	case "Benz":
		return &BenzCar{}
	case "BMW":
		return &BMWCar{}
	case "Audi":
		return &AudiCar{}
	}
	return nil
}
