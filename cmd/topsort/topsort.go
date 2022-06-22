package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	top "github.com/h-mavrodiev/data-structures-and-algorithms/pkg/topsort"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	path := filepath.Join(exPath, "../../pkg/topsort/test-data.txt")

	fmt.Println("Please enter path to a packages text file ... ")
	if _, err := fmt.Scanln(&path); err != nil {
		fmt.Printf("No path was provided.\nDefault packages would be installed from file %v\n", path)
	}

	graph := &top.Graph{}
	graph.ReadFile(path)
	soretedPackages, err := graph.DoTopSort()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Packages will be installed in the following order : %v", soretedPackages)
	}

}
