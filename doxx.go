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


func parseRawCpfData (rawCpfData string, raw bool) ([]string) {

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

	parsedData := []string{}

	if raw {
		fmt.Println(string(white), cpfData)
		return parsedData
	}

	fmt.Println(string(cyan), "[-] Parsing output... ")
	data := cpfData["PESSOAFISICA"].(map[string]interface{})
	
	
	for key, value := range data {
		realKeyName := cpfFields[strings.ToLower(key)]
		fullText := fmt.Sprint(string(white), realKeyName, string(green), " -> ",string(white), value)
	 	parsedData = append(parsedData, fullText)
	}
	
	return parsedData
}

func getCpfData (cpf string, raw bool) ([]string){
	apiUrl := fmt.Sprintf("https://api.nesi.dev/cpf/%s", cpf)
	resp, err := http.Get(apiUrl)

	noData := []string{}

	if err != nil {
		requestError()
		return noData;
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
		
	if err != nil {
		requestError()
		return noData;
	}
	parsedData := parseRawCpfData(string(body), raw)
	return parsedData
}
func clear(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func getParams() (
	bool, 
	string,
	bool, 
	//string
) {
	
	var cpf string        
 	var help bool   
	var raw bool
	//var output string
	
	white := "\033[37m"
	
    flag.StringVar(&cpf, "cpf", "", "Specify cpf to search")
  	flag.BoolVar(&help, "h", false, "Show help")
  	flag.BoolVar(&raw, "r", false, "Get raw value")
	//flag.StringVar(&output, "o", "", "Output values in file")

 	flag.Usage = func() {
        fmt.Printf("Usage: \n")
        fmt.Println(string(white), "doxx -cpf 453.178.287-91")
        fmt.Println(string(white), "Raw output: doxx -cpf 453.178.287-91 -r")
        //fmt.Println(string(white), "File output: doxx -cpf 453.178.287-91 -o filename.txt")
     	flag.PrintDefaults()
    }   
    flag.Parse()

	if cpf == "" || help {
		flag.Usage()
		return true, "", false//, ""
	}
	return false, 
	cpf, 
	raw
	//output
}

func outputGatheredData (data []string) {
	fmt.Println("");
	for _, value := range data{
		fmt.Println(value)
	}
	fmt.Println("");
}

func showConfig(raw bool, output string) {
	reset := "\033[0m"
	cyan := "\033[36m"
	green := "\033[32m"

	hasOutput := output == ""

	switch raw {
		case true:
			fmt.Println(string(cyan), "[-] desired format:", string(green), "raw", string(reset))
		break
		case false:
			fmt.Println(string(cyan), "[-] desired format:", string(green), "formatted", string(reset))
		break
	}
	switch hasOutput {
		case true:
			fmt.Println(string(cyan), "[-] desired output:", string(green), "print", string(reset))
		break
		case false:
			fmt.Println(string(cyan), "[-] desired output:", string(green), "file", string(reset))
		break
	}
}
func main() {
	clear()

	reset := "\033[0m"
	cyan := "\033[36m"
	green := "\033[32m"

	cancelExec, 
	cpf, 
	raw := getParams()
	
	if cancelExec {
		return;
	}

	fmt.Println("");

	fmt.Println(string(cyan), "[-] selected:", string(green), cpf, string(reset))
	
	
	showConfig(raw, "")
	
	fmt.Println(string(cyan), "[-] loading cpf data ...", string(reset))

	fmt.Println("");

	parsedData := getCpfData(cpf, raw)

 	outputGatheredData(parsedData)
}



