package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"time"
)
// 每个链码都需要定义结构体
type BlockChainRealEstate struct {
}

// Init 链码初始化   执行peer chaincode instantiate实例化的时候会调用该方法
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	var accountIds = [7]string{
		"5feceb66ffc8",
		"6b86b273ff34", 
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
		"5feceb66ffc2",
	}
	var userNames = [7]string{"管理员", "①号业主", "②号业主", "③号业主", "④号业主", "⑤号业主","新用户"}
	var balances = [7]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000,5000000}
	//初始化账号数据 遍历六个用户 
	for i, val := range accountIds {
		account := &lib.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		// 写入账本 账户一个一个写入账本
		if err := utils.WriteLedger(account, stub, lib.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	// 提取本次调用的交易所指定的参数
	// 例如： {"Args":["init","a","100",b,"200"]}
	// funcName = queryAccountList  args = 为用户的二进制数字

	// ChaincodeID: ChainCodeName,
		// Fcn:         fcn,
		// Args:        args,
	switch funcName {
	case "queryAccountList":
		return routers.QueryAccountList(stub, args)
// 链码增加功能，新增账户功能
	case "createAccount":
		return routers.CreateAccount(stub, args)
	case "createRealEstate":
		return routers.CreateRealEstate(stub, args)
	case "queryRealEstateList":
		return routers.QueryRealEstateList(stub, args)
	case "createSelling":
		return routers.CreateSelling(stub, args)
	case "createSellingByBuy":
		return routers.CreateSellingByBuy(stub, args)
	case "querySellingList":
		return routers.QuerySellingList(stub, args)
	case "querySellingListByBuyer":
		return routers.QuerySellingListByBuyer(stub, args)
	case "updateSelling":
		return routers.UpdateSelling(stub, args)
	case "createDonating":
		return routers.CreateDonating(stub, args)
	case "queryDonatingList":
		return routers.QueryDonatingList(stub, args)
	case "queryDonatingListByGrantee":
		return routers.QueryDonatingListByGrantee(stub, args)
	case "updateDonating":
		return routers.UpdateDonating(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
