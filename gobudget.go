package main

import ("fmt"
        "time"
        "encoding/json"
        "io/ioutil"
)

type Budget struct {
    User    	string          `json:"user"`
	Wallet  	map[string]int64	`json:"wallet"`
	Commodities	map[string]int64		`json:"commodities"`
	Annual  	map[string]int64	`json:"annual"`
	Savings  	map[string]int64	`json:"savings"`
    Id      	int64           `json:"id"`
    Created 	time.Time       `json:"created"`
    private 	string //unexported CapLetter
}

func main() {

	wmap := map[string]int64{
				"index":0,
				"actual":1000,
				"pocket":50,
				"gaz":50,
				"food":50,
	}
	cmap := map[string]int64{
				"index":0,
				"actual":1,
				"house": 600,
				"car": 450,
				"phone": 85,
				"creditcard": 50,
	}
	amap := map[string]int64{
				"index":0,
				"actual":1,
				"clothes": 600,
				"glasses": 300,
				"dentist": 300,
				"licences": 100,
				"tools": 600,
				"gifts": 300,
	}
	smap := map[string]int64{
				"index":0,
				"actual":1,
	}                
	
	
    myBudget := Budget{
        User:       	"Yann",
		Wallet:      	wmap,
		Commodities:	cmap,
		Annual:			amap,
		Savings:		smap,
        Id:         	99,
        Created:    	time.Now(),
        private:    	"Second-rate",
    }

	m := &myBudget

	m.Wallet["index"] = m.Wallet["pocket"] + m.Wallet["gaz"] + m.Wallet["food"]
	m.Commodities["index"] = (m.Commodities["house"] + m.Commodities["car"] + m.Commodities["phone"] + m.Commodities["creditcard"]) / 4
	m.Annual["index"] = (m.Annual["clothes"] + m.Annual["glasses"] + m.Annual["dentist"] + m.Annual["licences"] + m.Annual["tools"] + m.Annual["gifts"]) / 52


	m.Savings["actual"] = m.Wallet["actual"] - m.Wallet["index"] - m.Commodities["index"] - m.Annual["index"]
		m.Commodities["actual"] = m.Commodities["index"]
		m.Annual["actual"] = m.Annual["index"]

	fmt.Println(myBudget)


/*------jsonWRITER------*/
    jsonData, _ := json.MarshalIndent(myBudget, "", "    ")
    _ = ioutil.WriteFile("budget.json", jsonData, 0644)
    fmt.Println(string(jsonData))  
}
