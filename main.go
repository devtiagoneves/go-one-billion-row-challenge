package main

import (
	"fmt"
	"io"
	"net/http"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

// func main() {
// 	// Abrindo o arquivo measurements.txt
// 	measurements, err := os.Open("measurements.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Fechando o arquivo measurements.txt
// 	defer measurements.Close()

// 	dados := make(map[string]Measurement)

// 	// Lendo o arquivo measurements.txt
// 	scanner := bufio.NewScanner(measurements)
// 	for scanner.Scan() {
// 		rawData := scanner.Text()
// 		semicolon := strings.Index(rawData, ";")
// 		location := rawData[:semicolon]
// 		rawTemp := rawData[semicolon+1:]

// 		temp, _ := strconv.ParseFloat(rawTemp, 64)

// 		measurements, ok := dados[location]
// 		if !ok {
// 			measurements = Measurement{
// 				Min:   temp,
// 				Max:   temp,
// 				Sum:   temp,
// 				Count: 1,
// 			}
// 		} else {
// 			measurements.Count++
// 			measurements.Sum += temp
// 			measurements.Max = max(measurements.Max, temp)
// 			measurements.Min = min(measurements.Min, temp)
// 		}

// 		dados[location] = measurements
// 	}

// 	locations := make([]string, 0, len(dados))
// 	for location := range dados {
// 		locations = append(locations, location)
// 	}

// 	sort.Strings(locations)

// 	for _, location := range locations {
// 		measurements := dados[location]
// 		fmt.Printf("%s=%.1f/%.1f/%.1f, \n", location, measurements.Min, measurements.Max, measurements.Sum/float64(measurements.Count))
// 	}

// }

func main() {
	resp, err := http.Get("https://itarunning.com.br")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
