package builder


/**************************
Director
**************************/
type ManufacturingDirector struct {
	builder BuildProcess
}
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

/**************************
Process
**************************/
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Car builder
type CarBuilder struct {
	v VehicleProduct
}
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c }
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c }
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// Bike Builder
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b }
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b }
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

type VehicleProduct struct {
	Wheels    int
	Seats int
	Structure string
}

