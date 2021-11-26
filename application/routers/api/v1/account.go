package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
	"net/http"
)

type AccountIdBody struct {
	AccountId string `json:"accountId"`
}

type AccountRequestBody struct {
	Args []AccountIdBody `json:"args"`
}

// "args":
// [{"accountId":"4b227777d4d2"}]
// 服务数据格式
// @Summary 获取账户信息
// @Param account body AccountRequestBody true "account"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryAccountList [post]
func QueryAccountList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AccountRequestBody)
	fmt.Println(c)
	//解析Body参数 body是登陆用户的账户ID  参数绑定
	// if err := c.ShouldBind(body); err != nil {
	// 	appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
	// 	return
	// }
	if err := c.ShouldBind(body); err == nil {
		fmt.Println(body.Args)
	}else{
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}


	// 正好对应账户的12位16进制,通过这个用户ID,查询账本
	var bodyBytes [][]byte
	for _, val := range body.Args {
		bodyBytes = append(bodyBytes, []byte(val.AccountId))
	}
	
	//调用智能合约
	
	/* resp是通过链码查询的所有用户,登陆后返回当前用户,返回用户的账户ID,余额,用户
	姓名  需要反序列化
	[map[accountId:4e07408562be balance:5e+06 userName:③号业主
	*/
	 resp, err := bc.ChannelQuery("queryAccountList", bodyBytes)

	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	appG.Response(http.StatusOK, "成功", data)
}


// @Summary 增加账本中键值对，即注册账户
// @Param account body AccountRequestBody true "account"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createAccount [post]
func CreateAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AccountRequestBody)
	//解析Body参数 body是登陆用户的账户ID
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return

	// 正好对应账户的12位16进制,通过这个用户ID,查询账本
	var bodyBytes [][]byte
	for _, val := range body.Args {
		bodyBytes = append(bodyBytes, []byte(val.AccountId))
	}
	
	//调用智能合约
	
	/* resp是通过链码查询的所有用户,登陆后返回当前用户,返回用户的账户ID,余额,用户
	姓名  需要反序列化
	[map[accountId:4e07408562be balance:5e+06 userName:③号业主
	*/
	 resp, err := bc.ChannelExecute("createAccount", bodyBytes)
	
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	
	// 反序列化json
	// var data []map[string]interface{}
	// if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
	// 	appG.Response(http.StatusInternalServerError, "失败", err.Error())
	// 	return
	// }
	fmt.Println(resp)
	// appG.Response(http.StatusOK, "成功", data)
	appG.Response(http.StatusOK, "成功",nil)
}
}