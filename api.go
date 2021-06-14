package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var fallecidos []Fallecido

type Fallecido struct {
	RUC           int    `json:"ruc"`
	NombreEntidad string `json:"nombreentidad"`
	Departamento  string `json:"departamento"`
	Provincia     string `json:"provincia"`
	Distrito      string `json:"distrito"`
}

func lineStruct(datas [][]string) {
	for _, line := range datas {
		RUC, _ := strconv.Atoi(strings.TrimSpace(line[0]))

		fallecidos = append(fallecidos, Fallecido{
			RUC:           RUC,
			NombreEntidad: strings.TrimSpace(line[1]),
			Departamento:  strings.TrimSpace(line[2]),
			Provincia:     strings.TrimSpace(line[3]),
			Distrito:      strings.TrimSpace(line[4]),
		})
	}

}

func ReadCSVFile(url string) ([][]string, error) {
	// Abrir archivo CSV
	file, err := http.Get(url)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Body.Close()

	// Leer archivo CSV
	datas, err := csv.NewReader(file.Body).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return datas, nil
}

func main() {
	url := "https://raw.githubusercontent.com/Nico-j-j/TA2_Programacion_Concurrente_Y_Distribuida/master/dataset/Entidades_Contratantes_-20201001.csv"
	datas, err := ReadCSVFile(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Leyo archivos")
	lineStruct(datas)
	fmt.Println("Parseo Archivos")
}
