package goapi

// import "fmt" 

func GetClientName(id int) string {
	if id == 1 {
		return "Hugo Ramos"
	} else if id == 2 {
		return "Janaina"
	} else {
		return "Outro nome"
	}

	return "nada..."
}