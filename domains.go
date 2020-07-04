package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// Domains struct which contains
// an array of domains
type DomainsIndex struct {
    Domains []Domain `json:"domains"`
}

// Domain struct which contains a name
// a type and a list of social links
type Domain struct {
    Name   		string 	`json:"name"`
    Registrar   string 	`json:"registrar"`
    Price    	int    	`json:"price"`
    Renewal 	string	`json:"renewal"`
}

func (d Domain) String() string {
	return fmt.Sprintf(d.Name)
}

func main() {

/*------PARSER------*/
    jsonFile, err := os.Open("domains.json")
    	if err != nil {
        	fmt.Println(err)
   		}
		fmt.Println("Successfully Opened domains.json")
		fmt.Println(" ")
	defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var s DomainsIndex
	json.Unmarshal(byteValue, &s)
/*------------------*/

/*------PRINTER------*/
	a := " | "
	ptag := "$CAD"
	for i := 0; i < len(s.Domains); i++ {
		fmt.Println(i + 1, a, s.Domains[i].Name, a, s.Domains[i].Registrar, a, s.Domains[i].Price, ptag, a, s.Domains[i].Renewal)
		fmt.Println(" ")
	}
	Total := 0
	for t := 0; t < len(s.Domains); t++ {
		Total = Total + s.Domains[t].Price
	}
	fmt.Println("Total price: ", Total, ptag)
/*---------------------------*/

}
