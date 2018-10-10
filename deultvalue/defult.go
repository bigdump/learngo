package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	str  string
	str2 string
	bo   bool
	in   int
	fi   interface{}
}

type Payload struct {
	code  string
	Param string
}

func main() {
	mp := make(map[string]string)
	mp["code"] = con2
	mp["params"] = ""
	data, _ := json.Marshal(mp)
	fmt.Println(string(data))
	param := make(map[string]string)
	json.Unmarshal(data, &param)
	code := string(param["code"])
	fmt.Println(code)
	fmt.Println()
	params := "init(" + string(param["params"]) + ")"
	fmt.Println(params)
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
