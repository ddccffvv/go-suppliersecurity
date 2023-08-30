# go-suppliersecurity

## About suppliersecurity.info

Suppliersecurity.info is a service which monitors and catalogs compliance, privacy and security data of the top SaaS companies. More information on our website: https://suppliersecurity.info.

## Prerequisites

Generate an API key at https://suppliersecurity.info/api_tokens (Authentication required)

## Example usage

```go
package main

import (
        "fmt"
        "github.com/ddccffvv/go-suppliersecurity" // get the library
        "log"
        "os"
)

func main() {
        key := os.Getenv("SUPPLIERSECURITY_API_KEY") // retrieve the API key in some way
        if key == "" {
                log.Fatal("Please set the environment variable SUPPLIERSECURITY_API_KEY. Find or generate your key at https://suppliersecurity.info/api_tokens")
        }
        res, err := suppliersecurity.RetrieveProduct(key, "1") // call the library
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(res) // get the result
        fmt.Println(res.Name)
}
```

## Available functionality

There are 2 methods available:

`RetrieveProduct(api_key string, product_id string) (Product, error)` can be used when the product id is known. You'll receive a single `Product`.

`Search(api_key string, search_term string) (Productlist, error)` is used to retrieve one or more products using (part of) the URL as search term. You'll receive a `Productlist`, which contains multiple `Product`s.

These are the available fields:

```go
type Product struct {
	Id                            int //internal id of the product
	Name                          string // name of the product
	Url                           string // url of the product
	DataProtectionOfficerContact  string // contact (email) of the data protection officer
	PrivacyPages                  []string // list of discovered pages containing information about the privacy policy and program
	SecurityPages                 []string // list of discovered pages containing information about the security policy and program
	ProductPage                   string // product page on suppliersecurity.info
	HipaaCertified                bool // is this product hipaa certified (false or nil means "not detected")
	Iso27001Certified             bool // is this product iso 27001 certified (false or nil means "not detected")
	Soc2Type1Certified            bool // is this product soc 2 type 1 certified (false or nil means "not detected")
	Soc2Type2Certified            bool // is this product soc 2 type 2 certified (false or nil means "not detected")
	CustomerLightPersonalData     bool // this product typically contains "non-sensitive" customer information (such as names, email)
	CustomerSensitivePersonalData bool // this product typically contains "sensitive" customer information (such as financial data, religion,...)
	EmployeeLightPersonalData     bool // this product typically contains "non-sensitive" customer information (such as names, email)
	EmployeeSensitivePersonalData bool // this product typically contains "sensitive" customer information (such as financial data, religion,...)
	CorporateSensitiveInformation bool // this product typically contains corporate sensitive information (such as intellectual property, code,...)
	LastChecked                   string // date and time of the last time we checked the above information
}

type Productlist struct {
	Products []Product
}
```
