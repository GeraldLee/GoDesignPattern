package factory_method

import (
	"gotest.tools/assert"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	BMWFactory := BMWCarFactory{}
	BMWCar := BMWFactory.GetCar()
	assert.Equal(t,BMWCar.GetName(),"BMW")
	BenzFactory := BenzCarFactory{}
	BenzCar := BenzFactory.GetCar()
	assert.Equal(t,BenzCar.GetName(),"Benz")
	AudiFactory := AudiCarFactory{}
	AudiCar := AudiFactory.GetCar()
	assert.Equal(t,AudiCar.GetName(),"Audi")
}
