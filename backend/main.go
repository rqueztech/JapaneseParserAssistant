package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type KanjiMaps struct {
	KanjiFrequencies map[string]string `json:"kanjiFrequencies"`
	KanjiReadings    map[string]string `json:"kanjiReadings"`
}

func main() {
	// Initialize KanjiMaps struct with empty maps
	kanjiMaps := KanjiMaps{
		KanjiFrequencies: make(map[string]string),
		KanjiReadings:    make(map[string]string),
	}

	// Open kanjifrequency.csv
	frequencyfile, err := os.Open("kanjifrequency.csv")
	if err != nil {
		log.Fatalf("Kanji Frequency Not Found: %v", err)
	}
	defer frequencyfile.Close()

	// Create CSV reader for kanjifrequency.csv
	frequencyreader := csv.NewReader(frequencyfile)

	// Read and populate KanjiFrequencies map
	records, err := frequencyreader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading kanjifrequency.csv: %v", err)
	}
	for _, record := range records {
		kanjiMaps.KanjiFrequencies[record[0]] = record[1]
	}

	// Open kanjimeanings.csv
	meaningfile, err := os.Open("kanjimeanings.csv")
	if err != nil {
		log.Fatalf("Kanji Meanings Not Found: %v", err)
	}
	defer meaningfile.Close()

	// Create CSV reader for kanjimeanings.csv
	meaningreader := csv.NewReader(meaningfile)

	// Read and populate KanjiReadings map
	records, err = meaningreader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading kanjimeanings.csv: %v", err)
	}
	for _, record := range records {
		kanjiMaps.KanjiReadings[record[0]] = record[1]
	}

	// Setup HTTP handlers
	http.HandleFunc("/kanjifrequencies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(kanjiMaps.KanjiFrequencies)
	})

	http.HandleFunc("/kanjireadings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(kanjiMaps.KanjiReadings)
	})

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

