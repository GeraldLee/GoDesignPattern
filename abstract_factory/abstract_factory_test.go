package abstract_factory

import (
	"gotest.tools/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	BMWCar := GetCar("BMW")
	assert.Equal(t,BMWCar.GetName(),"BMW")
	BenzCar := GetCar("Benz")
	assert.Equal(t,BenzCar.GetName(),"Benz")
	AudiCar := GetCar("Audi")
	assert.Equal(t,AudiCar.GetName(),"Audi")
}

