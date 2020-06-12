package main

import (
	"fmt"
	"github.com/hyperchain/gosdk/account"
	"github.com/hyperchain/gosdk/common"
	"github.com/hyperchain/gosdk/rpc"
	"github.com/hyperchain/gosdk/utils/java"
)

func main() {
	// 获取logger
	logger := common.GetLogger("main")

	// 随机创建账户，获取account json
	accountJSON, sysErr := account.NewAccount("TomKK")
	if sysErr != nil {
		logger.Error(sysErr)
		return
	}

	logger.Debugf(accountJSON)

	// 根据account json得到*ecdsa.Key结构体
	key, sysErr := account.NewAccountFromAccountJSON(accountJSON, "TomKK")
	if sysErr != nil {
		logger.Error(sysErr)
		return
	}

	// 设置配置文件根目录，构建rpc结构体
	hrpc := rpc.NewRPCWithPath("gosdktest/conf")

	// 部署合约
	// 读取java合约，因为构造函数无参，故没有填写参数
	bin, sysErr := java.ReadJavaContract("/home/june/hyperchain/javasdk/src/test/resources/javaContractExample/contract01")
	if sysErr != nil {
		logger.Error(sysErr)
		return
	}

	// 构建部署交易结构体
	tx := rpc.NewTransaction(key.GetAddress()).
		Deploy(bin).
		VMType(rpc.JVM)

	// 交易签名
	tx.Sign(key)

	// 向hyperchain发起部署交易
	txReceipt, stdErr := hrpc.DeployContract(tx)
	if stdErr != nil {
		logger.Error(stdErr.String())
		return
	}

	// 得到合约地址
	contractAddr := txReceipt.ContractAddress

	// 调用java合约的 issue 方法，为自己发行1000token
	// 构建调用交易结构体
	tx = rpc.NewTransaction(key.GetAddress()).
		Invoke(contractAddr, java.EncodeJavaFunc("issue", key.GetAddress(), "1000")).
		VMType(rpc.JVM)

	// 交易签名
	tx.Sign(key)

	// 向hyperchain发起调用合约
	txReceipt, stdErr = hrpc.InvokeContract(tx)
	if stdErr != nil {
		logger.Error(stdErr.String())
		return
	}

	// 查看自己账户内token余额
	tx = rpc.NewTransaction(key.GetAddress()).
		Invoke(contractAddr, java.EncodeJavaFunc("getAccountBalance", key.GetAddress())).
		VMType(rpc.JVM)

	// 交易签名
	tx.Sign(key)

	// 向hyperchain发起调用合约
	txReceipt, stdErr = hrpc.InvokeContract(tx)
	if stdErr != nil {
		logger.Error(stdErr.String())
		return
	}

	// 解码合约返回值
	fmt.Printf("getAccountBalance返回值解码前: %s\n", txReceipt.Ret)
	fmt.Printf("getAccountBalance返回值解码后: %s\n", java.DecodeJavaResult(txReceipt.Ret))

	// 调用java合约testPostEvent方法
	// 构造调用交易结构体
	tx = rpc.NewTransaction(key.GetAddress()).
		Invoke(contractAddr, java.EncodeJavaFunc("testPostEvent")).
		VMType(rpc.JVM)

	// 签名
	tx.Sign(key)

	// 向hyperchain发起调用交易请求
	txReceipt, stdErr = hrpc.InvokeContract(tx)
	if stdErr != nil {
		logger.Error(stdErr.String())
		return
	}

	size := len(txReceipt.Log)
	for i := 0; i < size; i++ {
		fmt.Printf("java log 解码前: %s\n", txReceipt.Log[i].Data)
		// 解码java合约log
		decoded, sysErr := java.DecodeJavaLog(txReceipt.Log[i].Data)
		if sysErr != nil {
			logger.Error(sysErr)
			return
		}
		fmt.Printf("java log 解码后: %s\n", decoded)
	}
}



