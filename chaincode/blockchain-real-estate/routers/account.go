package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
)

// QueryAccountList 查询账户列表
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		// funcName = queryAccountList  args = 为用户的二进制数字
	// args是传入的参数 '{"Args":["init"]}' 除了init的后面的参数，此处为空
	var accountList []lib.Account
	// 创建一个空的用户对象列表
	fmt.Println("俺是嫩爹")
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var account lib.Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	// 序列化
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}



// CreateAccountList  用户注册并提交相关信息时创建账户
func CreateAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 创建一个空的用户对象列表
	fmt.Println("用户注册生效")
	// time.Local = timeLocal
	//初始化默认数据
// 将传入数据反序列化
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(args).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}



	var accountIds = []string{
		"5feceb66ffc8",
		"6b86b273ff34", 
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
		"5feceb66ffc2",
	}
	accountIds = append(accountIds,"5feceb66ffc81")
	
	var userNames = []string{"管理员", "①号业主", "②号业主", "③号业主", "④号业主", "⑤号业主","新用户"}
	var balances = []float64{0, 5000000, 5000000, 5000000, 5000000, 5000000,5000000}
	userNames = append(userNames, "qwe")
	balances = append(balances, 1231312)
	
	
	//初始化账号数据 遍历新增用户 
	for i, val := range accountIds {
		account := &lib.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		// 写入账本 账户一个一个写入账本  string用[]包住
		if err := utils.WriteLedger(account, stub, lib.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	return shim.Success(nil)
}