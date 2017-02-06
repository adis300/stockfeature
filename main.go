package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	_ = os.Mkdir("computed", os.ModePerm)

	computeFile(djiFileName)
	for _, stockName := range djiNames {
		computeFile(stockName)
	}
}

func computeFile(filename string) {
	sourceFile, err := os.Open("data/" + filename + ".csv")
	if err != nil {
		fmt.Println("Error: reading data"+filename, err)
		return
	}

	defer sourceFile.Close()

	computedFile, err := os.Create("computed/" + filename + ".csv")
	if err != nil {
		fmt.Println("Error: writing data:"+filename, err)
		return
	}
	defer computedFile.Close()

	reader := csv.NewReader(sourceFile)
	writer := csv.NewWriter(computedFile)

	err = writer.Write(headers)

	defer writer.Flush()

	filter := [][]float64{}
	dateFilter := []string{}

	lineCount := 0
	for {

		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// record is an array of string so is directly printable
		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		// and we can iterate on top of that
		fmt.Println(record)
		lineCount++

		open, _ := strconv.ParseFloat(record[1], 64)
		high, _ := strconv.ParseFloat(record[2], 64)
		low, _ := strconv.ParseFloat(record[3], 64)
		close, _ := strconv.ParseFloat(record[4], 64)

		if len(filter) < FILTER_SIZE_LG {
			filter = append(filter, []float64{open, high, low, close})
			dateFilter = append(dateFilter, record[0])
		} else {
			filter = append(filter[1:], []float64{open, high, low, close})
			dateFilter = append(dateFilter[1:], record[0])

			computedFeatures := extractFeature(filter, dateFilter)

			err = writer.Write(computedFeatures)

			if err != nil {
				panic(err)
			}
			// checkError("Cannot write to file", err)
		}
	}
}

func extractFeature(data [][]float64, dates []string) []string {

	closeValue := data[0][3]
	// Initialize features with date and close value
	features := []string{dates[0], strconv.FormatFloat(closeValue, 'f', 12, 64)}
	// Diff0: Same day movement
	features = append(features, strconv.FormatFloat((closeValue-data[0][0])/closeValue, 'f', -1, 64))
	// DiffHighLow: A measurement of stability
	features = append(features, strconv.FormatFloat((data[0][1]-data[0][2])/closeValue, 'f', -1, 64))

	// DiffSM: Small interval movement
	features = append(features, strconv.FormatFloat((closeValue-data[FILTER_SIZE_SM-1][3])/closeValue, 'f', -1, 64))
	// DiffMD: Medium interval movement
	features = append(features, strconv.FormatFloat((closeValue-data[FILTER_SIZE_MD-1][3])/closeValue, 'f', -1, 64))
	// DiffLG: Large interval movement
	features = append(features, strconv.FormatFloat((closeValue-data[FILTER_SIZE_LG-1][3])/closeValue, 'f', -1, 64))

	// Avg0: Sameday average relative to close
	features = append(features, strconv.FormatFloat((closeValue+data[0][0])/closeValue, 'f', -1, 64))
	// AvgSM: Small interval average relative to close
	sum := 0.0
	for i := 0; i < FILTER_SIZE_SM; i++ {
		sum += data[i][3]
	}
	features = append(features, strconv.FormatFloat((sum/FILTER_SIZE_SM_FLOAT)/closeValue, 'f', -1, 64))
	// AvgMD: Mid interval average relative to close
	for i := FILTER_SIZE_SM; i < FILTER_SIZE_MD; i++ {
		sum += data[i][3]
	}
	features = append(features, strconv.FormatFloat((sum/FILTER_SIZE_MD_FLOAT)/closeValue, 'f', -1, 64))
	// AvgLG: Large interval average relative to close
	for i := FILTER_SIZE_MD; i < FILTER_SIZE_LG; i++ {
		sum += data[i][3]
	}
	features = append(features, strconv.FormatFloat((sum/FILTER_SIZE_LG_FLOAT)/closeValue, 'f', -1, 64))

	return features
}
