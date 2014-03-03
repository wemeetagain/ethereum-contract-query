package ethereum-contract-query

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"    
    "net/http"
    "github.com/ethereum/eth-go"
    )

const (
    portDefaut = "7311"
    portDesc = "port to run on"
    )

var (
    port string
    ethereum *Ethereum
    )

func init() {
    flag.StringVar(&port,"port",portDefault,portDesc)
    flag.StringVar(&port,"p",portDefault,portDesc+" (shorthand)")
    flag.Parse()
}

func queryHandler(w http.ResponseWriter,r *http.request) {
    var query *QueryInfo
    // parse query string
    v := r.URL.Query()
    query.ContractAddr = v.get("contract")
    query.Index = v.get("index")
    //parse request body
    rawbody, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    json.Unmarshal(rawbody,&query)
    
    // query ethereum
    resp, err := ethereum.QueryContract(query.ContractAddr,query.Index)
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, resp)
    return
}

func main() {
    // initialize and start ethereum service
    ethereum, err := eth.New(eth.CapDefault, false)
    ethereum.Start()
    if err != nil {
        log.Println("eth start err:", err)
        return
    }
    
    // start http server
    http.HandleFunc("/query",queryHandler)
    log.Fatal(http.ListenAndServe(":"+port,nil))
}
