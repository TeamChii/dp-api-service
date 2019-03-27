package resolver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"th.co.droppoint/config"
)

func GetMerchantByIdResolver(params graphql.ResolveParams) (interface{}, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.URL_API+"/merchant/id/"+params.Args["id"].(string), nil)
	//req.Header.Add("Authorization", "Bearer "+params.Args["authorization"].(string))
	req.Header.Add("Authorization", "Bearer "+params.Args["authorization"].(string))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Cal Error")
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var userJson map[string]interface{}
	json.Unmarshal(data, &userJson)
	return userJson, nil
}

/*func GetCustomerSearchResolver(params graphql.ResolveParams) (interface{}, error) {

	var str = `{"searchString":"` + params.Args["searchString"].(string) + `",
		"paging":{
			"pageNo":` + params.Args["pageNo"].(string) + `,
			"pageSize":` + params.Args["pageSize"].(string) + `,
			"sortBy":"` + params.Args["sortBy"].(string) + `",
			"orderBy":"` + params.Args["orderBy"].(string) + `"
			}
		}`

	var jsonStr = []byte(str)
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.URL_API+"/customer-master/search-by-attr", bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", params.Args["authorization"].(string))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Cal Error")
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(data)
	var userJson map[string]interface{}
	json.Unmarshal(data, &userJson)
	return userJson, nil
}*/
