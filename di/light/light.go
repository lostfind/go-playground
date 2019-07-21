package light

type Incandescent struct{}

func (*Incandescent) LightUp() string {
	return "フィラメントが光るよ!"
}

type LedLight struct{}

func (*LedLight) LightUp() string {
	return "LEDが光るよ!"
}
