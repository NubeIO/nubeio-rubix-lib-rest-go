package rubixio

import (
	"reflect"
)

//GetInputValues get all the input values by passing in the IONum (pass in UI as an example)
func GetInputValues(strut interface{}, key string) (found bool, temp, voltage, current, raw float64, digital int, err error) {
	val := reflect.ValueOf(strut).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if typeField.Name == key {
			found = true
			thermistor10KType2 := val.Field(i).FieldByName("Thermistor10KType2")
			temp = float(thermistor10KType2)

			voltageDc := val.Field(i).FieldByName("VoltageDc")
			voltage = float(voltageDc)

			rawVal := val.Field(i).FieldByName("Raw")
			raw = float(rawVal)

			currentVal := val.Field(i).FieldByName("Current")
			current = float(currentVal)

			digitalVal := val.Field(i).FieldByName("Digital")
			digital = integer(digitalVal)
		}
	}
	return
}

func float(v reflect.Value) float64 {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("Recovered")
		}
	}()
	return v.Float()
}

func integer(v reflect.Value) int {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("Recovered")
		}
	}()
	return int(v.Int())
}
