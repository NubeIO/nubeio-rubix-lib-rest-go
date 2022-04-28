package rubixio

type Write struct {
	IoNum string `json:"io_num"`
	Value int    `json:"value"`
}

type BulkWrite struct {
	IONum string `json:"io_num"`
	Value int    `json:"value"`
}

type BulkResponse struct {
	Ok bool `json:"ok"`
}

type Ping struct {
	Ok     bool `json:"ok"`
	PigPIO bool `json:"pigio_is_running"`
}

type Inputs struct {
	UI1 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI1"`
	UI2 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI2"`
	UI3 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI3"`
	UI4 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI4"`
	UI5 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI5"`
	UI6 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI6"`
	UI7 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI7"`
	UI8 struct {
		IoNum              string  `json:"io_num"`
		Raw                int     `json:"raw"`
		Thermistor10KType2 float64 `json:"thermistor_10k_type_2"`
		VoltageDc          float64 `json:"voltage_dc"`
		Current            float64 `json:"current"`
		Digital            int     `json:"digital"`
	} `json:"UI8"`
}
