package main

// import (
// 	"fmt"
// )

func main() {

}

func os10maioresEstadosDoBrasil() ([]string, error) {
	var data []string
	data = append(data,
		[]string{
			"Amazonas",
			"Pará",
			"Mato Grosso",
			"Minas Gerais",
			"Bahia",
			"Mato Grosso do Sul",
			"Goiás",
			"Maranhão",
			"Rio Grande do Sul",
			"Tocantins",
		}...,
	)
	return data, nil
}
