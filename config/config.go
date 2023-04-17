package config

var (
	Payload = "x"
	PayloadUrl string

)

func SetPayloadUrl(p string)  {
	PayloadUrl = p
}
func GetPayloadUrl()string  {
	return PayloadUrl
}