func main() {

    /*------jsonPARSER------*/
    jsonFile, err := os.Open("budget.json")
    	if err != nil {
        	fmt.Println(err)
   		}
		fmt.Println("Successfully Opened budget.json\n")
		defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
    var b Budget
	json.Unmarshal(byteValue, &b)
    /*------------------*/
    

    /*------PRINTER------*/
        fmt.Println("\nUser:",b.User)
		fmt.Println("\nDeposit:",b.Wallet["actual"])
		fmt.Println("\nWallet:",b.Wallet["index"])
		fmt.Println("\nCommodities:",b.Commodities["actual"])
		fmt.Println("\nAnnual:",b.Annual["actual"])
		fmt.Println("\nSavings:",b.Savings["actual"])
        fmt.Println("\nId:",b.Id)
        fmt.Println("\nCreated:",b.Created)
        fmt.Printf("\n------------------------------------------------------\n")
        fmt.Println("private: ->",b.private, "<-")
    /*-------------------*/

/*------Algorithm------*/
	wi, wa := b.Wallet["index"], b.Wallet["actual"]
	ci, ca  := b.Commodities["index"], b.Commodities["actual"]
	ai, aa  := b.Annual["index"], b.Annual["actual"]
	si, sa := b.Savings["index"], b.Savings["actual"]


	fmt.Println(wi, wa,
				ci, ca,
				ai, aa,
				si, sa)
	/*	saved := 150, 150, 50, 0
	ii := pPay
	ei := ci + wi + ai
	ii -= ei
	saved = ii;
	accSavings += ii;
/*------------------*/
}
