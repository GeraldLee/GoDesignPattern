package simple_factory

import (
	"gotest.tools/assert"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	factory := SimpleFactory{}
	BMWCar := factory.NewCar("BMW")
	assert.Equal(t,BMWCar.GetName(),"BMW")
	BenzCar := factory.NewCar("Benz")
	assert.Equal(t,BenzCar.GetName(),"Benz")
	AudiCar := factory.NewCar("Audi")
	assert.Equal(t,AudiCar.GetName(),"Audi")
	AutoCar := factory.NewCar("Auto")
	assert.Equal(t,AutoCar,nil)
}