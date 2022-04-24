package conn

import (
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	"os"
)

type Conf_result struct {
	Secret      string
	Cdn_ip_port string
	Sni_name    string
	Host_name   string
}

func Get_sni_name() string {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Conf_result
	json.Unmarshal([]byte(byteValue), &result)

	return result.Sni_name
}

func Get_host_name() string {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Conf_result
	json.Unmarshal([]byte(byteValue), &result)

	return result.Host_name
}
