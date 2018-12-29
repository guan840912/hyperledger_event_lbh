package main

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"time"
)

func main() {

	// Step1 讀取設定檔，以創建總客戶端
	clientSDK,err1 :=fabsdk.New(config.FromFile("config.yaml"))

	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(clientSDK)

	// Step2 使用總客戶端創建系統資源管理客戶端
	resourceManagerContext := clientSDK.Context(fabsdk.WithUser("Admin"),fabsdk.WithOrg("Org1"))

	resourceClient ,errRe := resmgmt.New(resourceManagerContext)
	if errRe != nil{
		fmt.Println(errRe)
	}
	fmt.Println(resourceClient)


	// Step3 使用總客戶端創造MSP客戶端
	// caClient
	mspClient,err2 := msp.New(clientSDK.Context(),msp.WithOrg("Org1"))
	if err2 !=nil {
		fmt.Println(err2)
	}
	println("1")
	println(mspClient)


	// 取得Ca 管理員
	adminIdentity ,err3 :=mspClient.GetSigningIdentity("admin")
	if err3 != nil {
		fmt.Println(err3)
	}
	println("2")
	println(adminIdentity)


	// Step4 使用總客戶端創建channel客戶端

	channelProvider := clientSDK.ChannelContext("mychannel",
		fabsdk.WithUser("Admin"),
		fabsdk.WithOrg("Org1"))




		//Chaincode event 監聽事件操作

	// 創建連結
	channelClient, _ := channel.New(channelProvider)
	registration,notifier ,err := channelClient.RegisterChaincodeEvent("chaincode_event","dataUpdate")
	if err != nil {
		fmt.Println("failed to register chaincode event")
	}
	defer channelClient.UnregisterChaincodeEvent(registration)

	// 將event 打印出來
	select {
	case ccEvent := <-notifier:
		fmt.Printf("received chaincode event %v\n", ccEvent)

	case <-time.After(time.Second * 40):
		fmt.Println("timeout while waiting for chaincode event")
	}





	// 專門的event客戶端

	// 監聽BlockEvent

	// 若不加上WithxxxEvents， 客戶端不會允許進行聆聽
	//eventClient, _ := event.New(channelProvider,event.WithBlockEvents())
	//
	//registration,notifier ,err := eventClient.RegisterBlockEvent()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	fmt.Println("failed to register block event")
	//}
	//defer eventClient.Unregister(registration)
	//
	//select {
	//case blockEvent := <-notifier:
	//	fmt.Printf("received  event %v\n", blockEvent)
	//
	//case <-time.After(time.Second * 40):
	//	fmt.Println("timeout while waiting for block event")
	//}






	// filterBlockEvent監聽方式

	// 若不加上WithxxxEvents， 客戶端不會允許進行聆聽
	//eventClient, _ := event.New(channelProvider,event.WithBlockEvents())
	//
	//registration,notifier ,err := eventClient.RegisterFilteredBlockEvent()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	fmt.Println("failed to register Filtered block event")
	//}
	//defer eventClient.Unregister(registration)
	//
	//select {
	//case filterBlockEvent := <-notifier:
	//	fmt.Printf("received  event %v\n", filterBlockEvent)
	//
	//case <-time.After(time.Second * 40):
	//	fmt.Println("timeout while waiting for filter block event")
	//}



}
