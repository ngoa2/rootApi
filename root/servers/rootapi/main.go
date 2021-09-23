package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Brand data type for export
type Brand struct {
	ID               int
	BrandName        string
	Environment      int
	EthicalPractices int
	Transparency     int
	Average          float64
	Tags             string
	AltBrands        string
}

var db *sql.DB

func main() {

	// get the value of the ADDR environment variable
	addr := os.Getenv("ADDR")

	tlscert := os.Getenv("TLSCERT")
	tlskey := os.Getenv("TLSKEY")

	if len(tlscert) == 0 || len(tlskey) == 0 {
		// write error to standard out, and exit with non zero code
		err := errors.New("Missing TLSCERT or TLSKEY")
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	// if it's blank, default to ":80", which means
	// listen port 80 for requests addressed to any host
	if len(addr) == 0 {
		addr = ":443"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/brandsearch", searchHandler)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, mux))
}

// Search handler queries the database, and responds with a json encoded struct of the requested row
func searchHandler(w http.ResponseWriter, r *http.Request) {

	// - Add an HTTP header to the response with the name
	// `Access-Control-Allow-Origin` and a value of `*`. This will
	//  allow cross-origin AJAX requests to your server.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 	Get the `name` query string parameter value from the request.
	// 	If it is not supplied, respond with an http.StatusBadRequest error.
	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Bad request error", http.StatusBadRequest)
		return
	}

	// call extract brand, to get a brand struct of the requested row
	brand, err := extractSearchItem(name)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(brand)
}

func extractSearchItem(name string) (*Brand, error) {
	//create the data source name, which identifies the
	//user, password, server address, and default database
	//dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3307)/rootdb", "SSAJpass")
	dsn := fmt.Sprintf("root:%s@tcp(172.17.0.3:3306)/rootdb", "SSAJpass")

	//create a database object, which manages a pool of
	//network connections to the database server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error opening database: %v\n", err)
		os.Exit(1)
	}

	//ensure that the database gets closed when we are done
	defer db.Close()

	//for now, just ping the server to ensure we have
	//a live connection to it
	if err := db.Ping(); err != nil {
		fmt.Printf("error pinging database: %v\n", err)
	} else {
		fmt.Printf("successfully connected!\n")
	}

	insq := "insert into brands(id, BrandName, Environment, EthicalPractices, Transparency, Average, Tags, AltBrands) values (?,?,?,?,?,?,?,?)"
	db.Exec(insq, "1", "HM", 3, 3, 3, 3, "100% organic cotton, Child labour policies, Freedom of association, Nondiscrimination policies, Recycled materials, Reports available online, Supply chain transparent", "https://backbeat.co/, https://kotn.com/, https://www.thereformation.com/")
	db.Exec(insq, "2", "Uniqlo", 3, 1, 2, 2, "Eco-friendly materials, Living wage payment unclear, Low supply chain transparency, No COVID-19 worker protection policies, Reports available online", "https://synergyclothing.com/, https://wearpact.com/, https://www.bleed-clothing.com/english")
	db.Exec(insq, "3", "Zara", 1, 1, 2, 1.3, "100% renewable energy, Low supply chain transparency, No guarantee of living wage, Reports available online, Unfavorable working conditions, Uses wool and leather", "https://www.thereformation.com/, https://www.ganni.com/us/home, https://hopaal.com/")
	db.Exec(insq, "4", "Nike", 2, 2, 2, 2, "100% renewable energy, FLA Workplace Code of Conduct Certified, Moderate supply chain transparency, Reports available online, Sustainable Apparel Coalition Certified, Uses wool and leather", "https://www.adidas.com/us, https://www.patagonia.com/home/, https://www.bleed-clothing.com/english")
	db.Exec(insq, "5", "Adidas", 3, 3, 3, 3, "100% sustainable cotton, Chemical usage reduction, Child labour policies, Eco-friendly materials, FLA Workplace Code of Conduct Certified, Nondiscrimination policies, Recycled materials, Reports available online, Supply chain transparent, Water usage reduction", "https://www.patagonia.com/home/, https://www.bleed-clothing.com/english, https://superstainable.com/")
	db.Exec(insq, "6", "Patagonia", 3, 3, 3, 3, "Chemical usage reduction, Eco-friendly materials, FLA Workplace Code of Conduct Certified, Fair Trade Certified, Recycled materials, Reports available online, Supply chain transparent, Uses blend of recycled wool, Water usage reduction", "https://houdinisportswear.com/en-us, https://finisterre.com/, https://superstainable.com/")
	db.Exec(insq, "7", "Reformation", 3, 2, 3, 2.6, "100% Living Wage, Bluesign Certified, Child labour policies, Eco-friendly materials, No exotic animal use, OEKO TEX STANDARD 100 Certified, Recycled materials, Reports available online, Supply chain transparent, Uses wool and leather", "https://synergyclothing.com/, https://www.ganni.com/us/home, https://hopaal.com/")
	db.Exec(insq, "8", "Allbirds", 3, 2, 3, 2.6, "Eco-friendly materials, Low supply chain transparency, No exotic animal use, No guarantee of living wage, Recycled materials, ZQ Merino Certified Wool", "https://www.adidas.com/us, https://kotn.com/, https://www.bleed-clothing.com/english")
	db.Exec(insq, "9", "Everlane", 2, 2, 2, 2, "No COVID-19 worker protection policies, No chemical waste reduction initiatives, No guarantee of living wage, No textile waste reduction initiatives, Recycled materials", "https://hopaal.com/, https://kotn.com/, https://synergyclothing.com/")

	// if err != nil {
	// 	fmt.Printf("error inserting new row: %v\n", err)
	// } else {
	// 	//get the auto-assigned ID for the new row
	// 	id, err := res.LastInsertId()
	// 	if err != nil {
	// 		fmt.Printf("error getting new ID: %v\n", id)
	// 	} else {
	// 		fmt.Printf("ID for new row is %d\n", id)
	// 	}
	// }

	log.Print("extracting item...")
	log.Print("brand name is " + name)
	stmt := "SELECT * FROM brands WHERE BrandName = ?;"
	brand := Brand{}
	log.Print("Query mysql database for row")
	row := db.QueryRow(stmt, name)

	log.Print("Scan row into brand struct")
	err = row.Scan(&brand.ID, &brand.BrandName, &brand.Environment, &brand.EthicalPractices, &brand.Transparency, &brand.Average, &brand.Tags, &brand.AltBrands)

	log.Print(brand)

	if err != nil {
		log.Printf("error scanning row")
		return nil, err
	}

	log.Print("return brand")
	return &brand, nil
}
