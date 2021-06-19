package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

var testSet []Entidad
var TrainSet []Entidad
var k int

func main_knn() {
	k = 36
	fmt.Println("total ")
	fmt.Println(len(entidades))
	/*for i := range entidades {
		testSet = append(testSet, entidades[i])
	}*/

	var predictions []string
	fmt.Println("test lenght")
	fmt.Println(len(testSet))

	for x := 0; x < len(testSet); x++ {

	}
	/*for x := range testSet {
		result := testCase(testSet, testSet[x], k)
		predictions = append(predictions, result[0].key)
		fmt.Printf("Predicted: %s, Actual: %s\n", result[0].key, testSet[x].NombreEntidad)
	}*/

	accuracy := getAccuracy(testSet, predictions)
	fmt.Printf("Accuracy: %f%s\n", accuracy, "%")
}

/*func testCase(trainSetA []Entidad, testSetObject Entidad, k int) sortedClassVotes {
	fmt.Println(testSetObject)
	neighbors := getNeighbors(trainSetA, testSetObject, k)
	result := getResponse(neighbors)
	return result
}*/

func getAccuracy(testSet []Entidad, predictions []string) float64 {
	correct := 0

	for x := range testSet {
		if testSet[x].NombreEntidad == predictions[x] {
			correct += 1
		}
	}

	return (float64(correct) / float64(len(testSet))) * 100.00
}

type classVote struct {
	key   string
	value int
}

type sortedClassVotes []classVote

func (scv sortedClassVotes) Len() int           { return len(scv) }
func (scv sortedClassVotes) Less(i, j int) bool { return scv[i].value < scv[j].value }
func (scv sortedClassVotes) Swap(i, j int)      { scv[i], scv[j] = scv[j], scv[i] }

func getResponse(neighbors []Entidad) sortedClassVotes {
	classVotes := make(map[string]int)

	for x := range neighbors {
		response := neighbors[x].NombreEntidad
		if contains(classVotes, response) {
			classVotes[response] += 1
		} else {
			classVotes[response] = 1
		}
	}

	scv := make(sortedClassVotes, len(classVotes))
	i := 0
	for k, v := range classVotes {
		scv[i] = classVote{k, v}
		i++
	}

	sort.Sort(sort.Reverse(scv))
	return scv
}

type distancePair struct {
	record   Entidad
	distance float64
}

type distancePairs []distancePair

func (slice distancePairs) Len() int           { return len(slice) }
func (slice distancePairs) Less(i, j int) bool { return slice[i].distance < slice[j].distance }
func (slice distancePairs) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func getNeighbors(trainingSet []Entidad, testRecord Entidad, k int) []Entidad {
	var distances distancePairs
	for i := range trainingSet {

		dist := euclidianDistance(testRecord, trainingSet[i])
		distances = append(distances, distancePair{trainingSet[i], dist})
	}

	sort.Sort(distances)

	var neighbors []Entidad

	for x := 0; x < k; x++ {
		neighbors = append(neighbors, distances[x].record)
	}

	return neighbors
}

func euclidianDistance(instanceOne Entidad, instanceTwo Entidad) float64 {
	var distance float64

	distance += math.Pow(float64((instanceOne.RUC - instanceTwo.RUC)), 2)

	return math.Sqrt(distance)
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func contains(votesMap map[string]int, name string) bool {
	for s := range votesMap {
		if s == name {
			return true
		}
	}

	return false
}
