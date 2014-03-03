package ethereum-contract-query

import (
    "encoding/json"
    "github.com/ethereum/eth-go"
    )

type QueryInfo struct {
    ContractAddr string `json:"contract,omitempty"`
    Index string `json:"index,omitempty"`
}

func (e *Ethereum) QueryContract(addr string, index string) ([]byte,error) {
    rawcontract, err := e.db.Get(arr)
    if err != nil {
        return []byte(`{"response":"error","data":`+err.Error()+`}`), err
    }
    data := ethutil.Encode(rawcontract).Get(index)
    resp := make(map[string]interface{})
    resp["response"] = "success"
    resp["data"] = data
    return json.Marshal(resp), nil
}
