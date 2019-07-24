package mockDemo

type IPeople interface {
	GetName() string
	SetName(string) string
}

func GetPeopleName(ip IPeople) string {
	return ip.GetName()
}

func SetPeopleName(ip IPeople, name string) string {
	return ip.SetName(name)
}
