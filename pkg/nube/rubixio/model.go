package rubixio

type Write struct {
	IoNum string `json:"io_num"`
	Value int    `json:"value"`
}

type Ping struct {
	Ok bool `json:"ok"`
}

type Inputs struct {
	UI1 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI1"`
	UI2 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI2"`
	UI3 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI3"`
	UI4 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI4"`
	UI5 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI5"`
	UI6 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI6"`
	UI7 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI7"`
	UI8 struct {
		Raw     int     `json:"raw"`
		Temp10K float64 `json:"temp_10_k"`
		Volt    float64 `json:"volt"`
		Current float64 `json:"current"`
		Bool    int     `json:"bool"`
	} `json:"UI8"`
}
