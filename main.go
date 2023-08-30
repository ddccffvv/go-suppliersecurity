package suppliersecurity

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"bytes"
)

const BaseURL = "https://suppliersecurity.info/api/"

type Product struct {
	Id                            int `json:"id"`
	Name                          string
	Url                           string
	DataProtectionOfficerContact  string
	PrivacyPages                  []string `json:"privacy_pages"`
	SecurityPages                 []string `json:"security_pages"`
	ProductPage                   string
	HipaaCertified                bool
	Iso27001Certified             bool
	Soc2Type1Certified            bool
	Soc2Type2Certified            bool
	CustomerLightPersonalData     bool
	CustomerSensitivePersonalData bool
	EmployeeLightPersonalData     bool
	EmployeeSensitivePersonalData bool
	CorporateSensitiveInformation bool
	LastChecked                   string
}

type Productlist struct {
	Products []Product
}

type productresult struct {
  Product Product `json:"product"`
}

func Search(api_key string, search_term string) (Productlist, error){
	URL := fmt.Sprintf("%sv1/search/", BaseURL)
	fmt.Println(search_term)

	jsonBody := []byte(fmt.Sprintf("{\"product\": {\"url\": \"%s\"}}", search_term))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, URL, bodyReader)

	if err != nil {
		return Productlist{}, fmt.Errorf("Could not create request: %s\n", err)
	}


	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", api_key))

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return Productlist{}, fmt.Errorf("Error making http request: %s\n", err)
	}


	if res.StatusCode == 401 {
		return Productlist{}, fmt.Errorf("Received http error code 401 (unauthorized). Is your API key correct?")
	}
	if res.StatusCode == 400 {
		return Productlist{}, fmt.Errorf("Received http error code 400 (bad request). Did you send data the server can't understand?")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return Productlist{}, fmt.Errorf("Error, could not read response body: %s\n", err)
	}

	jsondata := bytes.NewReader(resBody)

	variable := Productlist{}
	err = json.NewDecoder(jsondata).Decode(&variable)
	if err != nil {
		return Productlist{}, fmt.Errorf("Error, something went wrong: %s\n", err)
	}

	return variable, nil
}

func RetrieveProduct(api_key string, product_id string) (Product, error){
	URL := fmt.Sprintf("%sv1/products/%s", BaseURL, product_id)

	req, err := http.NewRequest(http.MethodGet, URL, nil)

	if err != nil {
		return Product{}, fmt.Errorf("Could not create request: %s\n", err)
	}


	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", api_key))

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return Product{}, fmt.Errorf("Error making http request: %s\n", err)
	}


	if res.StatusCode == 401 {
		return Product{}, fmt.Errorf("Received http error code 401 (unauthorized). Is your API key correct?")
	}
	if res.StatusCode == 400 {
		return Product{}, fmt.Errorf("Received http error code 400 (bad request). Did you send data the server can't understand?")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return Product{}, fmt.Errorf("Error, could not read response body: %s\n", err)
	}

	jsondata := bytes.NewReader(resBody)

	variable := productresult{}
	err = json.NewDecoder(jsondata).Decode(&variable)
	if err != nil {
		return Product{}, fmt.Errorf("Error, something went wrong: %s\n", err)
	}

	return variable.Product, nil

}
