package main
 
import (
    "flag"
    "fmt"
    "io/ioutil"
	"net/http"
	"encoding/json"
	"strings"
	"os"
	"os/exec"
)

func requestError(){
	red := "\033[31m"
	fmt.Println(string(red), "[✘] Error on pulling cpf data, exiting...")
}


func parseRawCpfData (rawCpfData string, raw bool) {

	cyan := "\033[36m"
	white := "\033[37m"
	green := "\033[32m"
	
	fmt.Println(string(green), "[✔] success!")

	cpfFields := make(map[string]string)
	cpfFields["@nopessoafisica"] = "Nome"
	cpfFields["@dtnascimento"] = "Data de Nascimento"
	cpfFields["@nologradouro"] = "Endereço"
	cpfFields["@nobairro"] = "Bairro"
	cpfFields["@nrcep"] = "Cep"
	cpfFields["@nomunicipio"] = "Municipio"
	cpfFields["@nomae"] = "Nome da mãe"
	cpfFields["@nrcpf"] = "Cpf"
	cpfFields["@dscomplemento"] = "Complemento"
	cpfFields["@sguf"] = "Estado (uf)"
	cpfFields["@nrlogradouro"] = "Numero"
	
	cpfData := map[string]interface{}{} 
	json.Unmarshal([]byte(rawCpfData), &cpfData)

	if raw {
		fmt.Println(string(white), cpfData)
		return;
	}

	fmt.Println(string(cyan), "[-] Parsing output... ")
	data := cpfData["PESSOAFISICA"].(map[string]interface{})
	

	for key, value := range data {
		realKeyName := cpfFields[strings.ToLower(key)]
		fmt.Println(string(white), realKeyName, string(green), " -> ",string(white), value)
	}
	


}

func getCpfData (cpf string, raw bool) {
	apiUrl := fmt.Sprintf("https://api.nesi.dev/cpf/%s", cpf)
	resp, err := http.Get(apiUrl)

	if err != nil {
		requestError()
		return;
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
		
	if err != nil {
		requestError()
		return;
	}
	parseRawCpfData(string(body), raw)
}
func clear(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func main() {
	clear()

	reset := "\033[0m"
	cyan := "\033[36m"

    var cpf string        
 	var help bool   
	var raw bool
    // flags declaration using flag package
    flag.StringVar(&cpf, "cpf", "", "Specify cpf to search")
  	flag.BoolVar(&help, "h", false, "Show help")
  	flag.BoolVar(&raw, "r", false, "Get raw value")

 	flag.Usage = func() {
        fmt.Printf("Usage: \n")
        fmt.Printf("doxx -cpf 453.178.287-91\n")
     	flag.PrintDefaults()
    }   
    flag.Parse()

	if cpf == "" || help {
		flag.Usage()
		return;
	}
	fmt.Println(string(cyan), "[-] selected:", cpf, string(reset))
	
	if raw {
		fmt.Println(string(cyan), "[-] desired output: raw ", string(reset))
	} else {
		fmt.Println(string(cyan), "[-] desired output: formatted ", string(reset))
	}
	
	
	fmt.Println(string(cyan), "[-] loading cpf data ...", string(reset))

	getCpfData(cpf, raw)

 	
}



