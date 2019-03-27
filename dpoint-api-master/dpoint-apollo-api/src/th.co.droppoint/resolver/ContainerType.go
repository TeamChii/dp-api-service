package resolver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../th.co.droppoint/config"
	"github.com/graphql-go/graphql"
)

func GetContainerTypeResolver(params graphql.ResolveParams) (interface{}, error) {

	var str = `{"paging":{ "pageNo":` + params.Args["pageNo"].(string) + `, "pageSize":` + params.Args["pageSize"].(string) + `, "sortBy":"` + params.Args["sortBy"].(string) + `", "orderBy":"` + params.Args["orderBy"].(string) + `"} }`

	var jsonStr = []byte(str)
	client := &http.Client{}
	req, err := http.NewRequest("POST", config.URL_API+"/container-type/list", bytes.NewBuffer(jsonStr))
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
