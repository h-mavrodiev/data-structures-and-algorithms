package topsort

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Graph structure as adjacency list graph
type Graph struct {
	size     int
	vertices []*Vertex
}

// Vertix structure
type Vertex struct {
	key      string
	adjacent []*Vertex
	indegree int
}

// DoTopSort performs a topological sort on a given DAG
func (g *Graph) DoTopSort() ([]string, error) {

	var sortedPackages []string
	var noDependencyPackagesSlice []string

	//checks for packages with no dependecies
	for _, vertex := range g.vertices {
		if vertex.indegree == 0 {
			noDependencyPackagesSlice = append(noDependencyPackagesSlice, vertex.key)
		}
	}

	for len(noDependencyPackagesSlice) > 0 {

		noDependencyPackage := noDependencyPackagesSlice[0]
		noDependencyPackagesSlice = noDependencyPackagesSlice[1:]

		sortedPackages = append(sortedPackages, noDependencyPackage)

		for _, vertex := range g.vertices {
			for _, adjVertex := range vertex.adjacent {
				if adjVertex.key == noDependencyPackage {
					g.RemoveEdge(noDependencyPackage, vertex.key)
					if vertex.indegree == 0 {
						noDependencyPackagesSlice = append(noDependencyPackagesSlice, vertex.key)
					}
				}
			}
		}
	}

	for _, v := range g.vertices {
		if v.indegree > 0 {
			return nil, errors.New("the packages that you want to install have at least one cyclic dependancy")
		}

	}
	return sortedPackages, nil
}

// ReadFile reads the input for the graph from a txt file
func (g *Graph) ReadFile(path string) {

	// open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// Close the file at the end of the program
	defer file.Close()

	//read the file line by line
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	packageSlice := g.AddVerticesFromFile(scanner)
	g.AddEdgesFromFile(packageSlice)
}

// AddVerticesFromFile loops on the input file and adds the vertices
func (g *Graph) AddVerticesFromFile(scanner *bufio.Scanner) [][]string {
	var packageSlice [][]string
	for scanner.Scan() {
		if _, err := strconv.Atoi(scanner.Text()); err == nil {
			size, err := strconv.Atoi(scanner.Text())
			g.size = size
			fmt.Printf("Number of packages to install : %v \n", size)
			if err != nil {
				log.Print(err)
			}
		} else {
			inputVertices := strings.Fields(scanner.Text())
			packageSlice = append(packageSlice, inputVertices)
			g.AddVertex(inputVertices[0])
		}
	}
	return (packageSlice)
}

// AddEdgesFromFile loops on the input file and adds the edges
func (g *Graph) AddEdgesFromFile(packageSlice [][]string) {

	for _, slice := range packageSlice {
		for num := range slice {
			if num > 0 {
				g.AddEdge(slice[0], slice[num])
			}
		}
	}
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k string) {
	if containsEdge(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it already exists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds edge to the graph
func (g *Graph) AddEdge(from, to string) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v to %v) ... the provided Vertices do not exist ", from, to)
		fmt.Println(err.Error())
	} else if containsEdge(fromVertex.adjacent, to) {
		err := fmt.Errorf("edge (%v to %v) already exists", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		fromVertex.indegree += 1
	}

}

// removes string from slice of strings
func removeIndex(fromVertAdj []*Vertex, toVert *Vertex) []*Vertex {
	for i, v := range fromVertAdj {
		if v == toVert {
			return append(fromVertAdj[:i], fromVertAdj[i+1:]...)
		}
	}
	return fromVertAdj
}

// AddEdge removes edge from the graph
func (g *Graph) RemoveEdge(from, to string) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v to %v) ... the provided Vertices do not exist ", from, to)
		fmt.Println(err.Error())
	} else {
		toVertex.adjacent = removeIndex(toVertex.adjacent, fromVertex)
		toVertex.indegree -= 1
	}

}

// getVertex returns a pointer to the Vertex witha key string
func (g *Graph) getVertex(k string) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// containsEdge checks if an edge already exists
func containsEdge(s []*Vertex, k string) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
		fmt.Printf("\nIndegree count: %v \n", v.indegree)
	}
}
