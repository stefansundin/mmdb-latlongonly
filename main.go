package main

import (
	"fmt"
	"os"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/oschwald/maxminddb-golang"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input.mmdb> <output.mmdb>\n", os.Args[0])
		os.Exit(1)
	}

	stat, err := os.Stat(os.Args[1])
	if err != nil {
		panic(err)
	}
	inputFileSize := stat.Size()

	if _, err := os.Stat(os.Args[2]); err == nil {
		fmt.Fprintf(os.Stderr, "Error: output path already exists.\n")
		os.Exit(1)
	}

	db, err := maxminddb.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening mmdb file: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	writer, err := mmdbwriter.New(mmdbwriter.Options{
		BuildEpoch:   int64(db.Metadata.BuildEpoch),
		DatabaseType: fmt.Sprintf("%s-LatLongOnly", db.Metadata.DatabaseType),
		Description: map[string]string{
			"en": fmt.Sprintf("%s (LatLongOnly)", db.Metadata.Description["en"]),
		},
		Languages: []string{
			"en",
		},
	})
	if err != nil {
		panic(err)
	}

	i := 0
	messagePrinter := message.NewPrinter(language.English)
	messagePrinter.Fprintf(os.Stderr, "NodeCount: %v\n", db.Metadata.NodeCount)
	messagePrinter.Fprintf(os.Stderr, "Processed %10d records", i)

	networks := db.Networks(maxminddb.SkipAliasedNetworks)
	for networks.Next() {
		record := make(map[string]any)
		subnet, err := networks.Network(&record)
		if err != nil {
			panic(err)
		}

		location, ok := record["location"].(map[string]any)
		if !ok {
			continue
		}
		latitude, ok := location["latitude"].(float64)
		if !ok {
			continue
		}
		longitude, ok := location["longitude"].(float64)
		if !ok {
			continue
		}

		new_record := mmdbtype.Map{
			"location": mmdbtype.Map{
				"latitude":  mmdbtype.Float64(latitude),
				"longitude": mmdbtype.Float64(longitude),
			},
		}

		err = writer.Insert(subnet, new_record)
		if err != nil {
			panic(err)
		}

		i += 1
		if i%1_000 == 0 {
			messagePrinter.Fprintf(os.Stderr, "\rProcessed %10d records", i)
		}
	}
	messagePrinter.Fprintf(os.Stderr, "\rProcessed %10d records\n", i)

	output, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}

	outputFileSize, err := writer.WriteTo(output)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Reduced size of the database by %.2f%%\n", 100*(1-float64(outputFileSize)/float64(inputFileSize)))

	reader, err := maxminddb.Open(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	if err := reader.Verify(); err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Output file is valid")
}
