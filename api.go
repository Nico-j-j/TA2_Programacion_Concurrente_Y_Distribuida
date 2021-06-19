//https://www.datosabiertos.gob.pe/dataset/entidades-contratantes-organismo-supervisor-de-las-contrataciones-del-estado-osce-0

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var entidades []Entidad

type Entidad struct {
	RUC                 int    `json:"ruc"`
	NombreEntidad       string `json:"nombreentidad"`
	Departamento        string `json:"departamento"`
	Provincia           string `json:"provincia"`
	Distrito            string `json:"distrito"`
	CodigoSIAF          int    `json:"codigo_siaf"`
	CodConSuCode        int    `json:"codconsucode"`
	Estado              string `json:"estado"`
	UltimaActualizacion string `json:"ultima_actualizacion"`
}

func lineStruct(datas [][]string) {
	for _, line := range datas {
		RUC, _ := strconv.Atoi(strings.TrimSpace(line[0]))
		CodigoSIAF, _ := strconv.Atoi(strings.TrimSpace(line[5]))
		CodConSuCode, _ := strconv.Atoi(strings.TrimSpace(line[6]))

		entidades = append(entidades, Entidad{
			RUC:                 RUC,
			NombreEntidad:       strings.TrimSpace(line[1]),
			Departamento:        strings.TrimSpace(line[2]),
			Provincia:           strings.TrimSpace(line[3]),
			Distrito:            strings.TrimSpace(line[4]),
			CodigoSIAF:          CodigoSIAF,
			CodConSuCode:        CodConSuCode,
			Estado:              strings.TrimSpace(line[7]),
			UltimaActualizacion: strings.TrimSpace(line[8]),
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

func GetEntities(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(entidades)
}

func GetEntity(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request) // Obtiene los parámetros
	for _, item := range entidades {
		ruc, _ := strconv.Atoi(parameters["RUC"])
		if item.RUC == ruc {
			json.NewEncoder(write).Encode(item)
			return
		}
	}
	json.NewEncoder(write).Encode(&Entidad{})
}

func GetCategoriaEntity(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Access-Control-Allow-Origin", "*")

	write.Header().Set("Content-Type", "application/json")

	var entity Entidad
	_ = json.NewDecoder(request.Body).Decode(&entity)

	k := 20 + rand.Intn(20)
	fmt.Println(k)
	/*result := testCase(entidades, entity, k)
	fmt.Printf("Predicted: %s, Actual: %s\n", result[0].key, entity.NombreEntidad)

	json := simplejson.New()
	json.Set("knn", result[0].key)
	json.Set("actual", entity.NombreEntidad)
	json.Set("predicted", result[0].key == entity.NombreEntidad)*/

	entidades = append(entidades, entity)

	//payload, _ := json.MarshalJSON()
	//write.Write(payload)
}

func CreateEntity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var entity Entidad
	_ = json.NewDecoder(r.Body).Decode(&entity)
	entidades = append(entidades, entity)
	json.NewEncoder(w).Encode(entity)
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

	request := mux.NewRouter()
	request.HandleFunc("/entidades", GetEntities).Methods("GET")
	request.HandleFunc("/entidades/{RUC}", GetEntity).Methods("GET")
	request.HandleFunc("/entidades", CreateEntity).Methods("POST")
	request.HandleFunc("/knn", GetCategoriaEntity).Methods("POST")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Start server
	port := ":7000"
	fmt.Println("Port: " + port)
	//main_knn()
	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(request)))

}
