package rubixio

import "reflect"

//GetInputValues get all the input values by passing in the IONum (pass in UI as an example)
func GetInputValues(strut interface{}, key string) (found bool, temp, voltage, current, raw float64, digital int, err error) {
	val := reflect.ValueOf(strut).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if typeField.Name == key {
			found = true
			Thermistor10KType2 := val.Field(i).FieldByName("Thermistor10KType2")
			if Thermistor10KType2.CanFloat() {
				temp = Thermistor10KType2.Float()
			}
			VoltageDc := val.Field(i).FieldByName("VoltageDc")
			if VoltageDc.CanFloat() {
				voltage = VoltageDc.Float()
			}
			Raw := val.Field(i).FieldByName("Raw")
			if Raw.CanFloat() {
				raw = Raw.Float()
			}
			Current := val.Field(i).FieldByName("Current")
			if Current.CanFloat() {
				current = Current.Float()
			}
			Digital := val.Field(i).FieldByName("Digital")
			if Digital.CanInt() {
				digital = int(Digital.Int())
			}
		}
	}
	return
}
