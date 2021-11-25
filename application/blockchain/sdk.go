package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// 配置信息
var (
	SDK           *fabsdk.FabricSDK          // Fabric提供的SDK
	ChannelName   = "assetschannel"          // 通道名称
	ChainCodeName = "blockchain-real-estate" // 链码名称
	Org           = "org1"                   // 组织名称
	User          = "Admin"                  // 用户
	ConfigPath    = "conf/config.yaml"       // 配置文件路径
)

// Init 初始化
func Init() {
	var err error
	// 通过配置文件初始化SDK
	SDK, err = fabsdk.New(config.FromFile(ConfigPath))
	if err != nil {
		panic(err)
	}
}

// ChannelExecute 区块链交互
func ChannelExecute(fcn string, args [][]byte) (channel.Response, error) {
	// 创建客户端，表明在通道的身份
	// args表示用户的ID   指定org1上的 管理员身份查询通道上某个组织,即建立客户端连接
	// 创建CC 可以不适用管理员哦
	ctx := SDK.ChannelContext(ChannelName, fabsdk.WithOrg(Org), fabsdk.WithUser(User))
	// 创建连接 RC 必须使用管理员
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	// 对区块链增删改的操作（调用了链码的invoke）
	resp, err := cli.Execute(channel.Request{
		ChaincodeID: ChainCodeName,
		// fcn表示执行的链码函数
		Fcn:         fcn,
		// 表示对哪个用户
		Args:        args,
		// 为链码请求配置  访问端节点
	}, channel.WithTargetEndpoints("peer0.org1.blockchainrealestate.com"))
	if err != nil {
		return channel.Response{}, err
	}
	//返回链码执行后的结果
	return resp, nil
}

// ChannelQuery 区块链查询
func ChannelQuery(fcn string, args [][]byte) (channel.Response, error) {
	// RC 资源管理客户端
	// CC 通道客户端
	// 创建客户端，表明在通道的身份
	// fcn 为要执行的方法，args是传入的参数
	ctx := SDK.ChannelContext(ChannelName, fabsdk.WithOrg(Org), fabsdk.WithUser(User))
	// fmt.Println("ctx=",ctx)
	cli, err := channel.New(ctx)
	// args是用户的二进制数字
	fmt.Println("args=",args)
	if err != nil {
		return channel.Response{}, err
	}
	// 对区块链查询的操作（调用了链码的invoke），将结果返回
	resp, err := cli.Query(channel.Request{
		ChaincodeID: ChainCodeName,
		Fcn:         fcn,
		Args:        args,
		// 为链码请求配置  访问端节点
	}, channel.WithTargetEndpoints("peer0.org1.blockchainrealestate.com"))
	if err != nil {
		return channel.Response{}, err
	}
	//返回链码执行后的结果
	return resp, nil
}
