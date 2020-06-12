package main

import (
	"fmt"
	"hello/common"
	"math/big"
)

type Node struct {
	str  string
	str2 string
	bo   bool
	in   int
	fi   interface{}
}

// Account is the hyperchain consensus representation format.
type Account struct {
	Nonce             uint64           `json:"nonce"`
	Balance           *big.Int         `json:"balance"`
	Root              common.Hash      `json:"merkleRoot"`
	CodeHash          []byte           `json:"codeHash"`
	DeployedContracts []string         `json:"contracts"`
	Creator           common.Address   `json:"creator"`
	Status            int              `json:"status"`
	CreateTime        uint64           `json:"createTime"`
	ApproverList      []common.Address `json:"approverList,omitempty"`
	NeededApproval    int              `json:"neededApproval,omitempty"`
}

type Payload struct {
	code  string
	Param string
}

func main() {
	// new
	account := Account{
		Nonce:             123,
		Balance:           big.NewInt(123214),
		Root:              common.Hash{12, 14, 4, 41, 34},
		CodeHash:          []byte{123, 23, 123},
		DeployedContracts: []string{"123", "abc"},
		Creator:           common.Address{31, 31, 3, 1},
		Status:            0,
		CreateTime:        0,
		ApproverList:      nil,
		NeededApproval:    0,
	}
	// call
	fmt.Printf("account address before is %p\n", &account)
	fmt.Printf("before DeployedContracts : %v\n", account.DeployedContracts[0])
	changeAccount(account)
	fmt.Printf("Nonce : %v\n", account.Nonce)
	fmt.Printf("Balance : %v\n", account.Balance)
	fmt.Printf("DeployedContracts : %v\n", account.DeployedContracts[0])
	fmt.Printf("codeHash : %v\n", account.CodeHash)

	// print

	//mp := make(map[string]string)
	//mp["code"] = con2
	//mp["params"] = ""
	//data, _ := json.Marshal(mp)
	//fmt.Println(string(data))
	//param := make(map[string]string)
	//json.Unmarshal(data, &param)
	//code := string(param["code"])
	//fmt.Println(code)
	//fmt.Println()
	//params := "init(" + string(param["params"]) + ")"
	//fmt.Println(params)
}

func changeAccount(account Account) {
	fmt.Printf("account address is %p\n", &account)
	account.Nonce = 1000
	account.Balance.Add(account.Balance, account.Balance)
	account.DeployedContracts[0] = "666"
	account.CodeHash[0] = 9

}

var con2 string = "4124144"

var con string = `
'use strict';

var BankStore = function() {
};

BankStore.prototype = {
    store1: 100,
    store2: "store"
};

var BankContract = function () {
    LocalContractStorage.defineProperty(this, "bankVault", {});
    LocalContractStorage.defineMapProperty(this, "bankMap", {});
};

// save value to contract, only after height of block, users can takeout
BankContract.prototype = {
    init: function (){

    },
    save: function () {
        this.bankVault = 100;
        return 100;
    },
    transfer: function(){
        var aa = this.bankVault.get("aaa");
        return aa;
    },
    show: function () {
        return this.bankVault;
    },

    mapAdd: function(){
        this.bankMap.put("balance", 998);
        return true;
    },

    mapGet: function(){
        return this.bankMap.get("balance");
    },
    mapAddObj: function(){
        var obj = new BankStore;
        obj.store1 = 102;
        obj.store2 = "store";
        this.bankMap.put("obj", obj);
    },
    mapGetObj: function(){
        var obj = this.bankMap.get("obj");
        return obj;
    }
};
`
