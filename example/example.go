package main

import (
	"encoding/json"
	"fmt"
	"github.com/NICEXAI/WeWorkFinanceSDK"
)

func main()  {
	corpID := "xxxxxxxxxxxx"
	corpSecret := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	rsaPrivateKey := `
-----BEGIN RSA PRIVATE KEY-----
xxxxxxxxxxxxxxxxxxxxxxxxxxx
-----END RSA PRIVATE KEY-----
	`
	//初始化客户端
	client, err := WeWorkFinanceSDK.NewClient(corpID, corpSecret, rsaPrivateKey)
	if err != nil {
		fmt.Printf("SDK 初始化失败：%v \n", err)
		return
	}
	//获取消息
	chatDataList, err := client.GetChatData(0, 100, "", "", 3)
	if err != nil {
		fmt.Printf("消息同步失败：%v \n", err)
		return
	}
	for _, chatData := range chatDataList {
		//消息解密
		chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if err != nil {
			fmt.Printf("消息解密失败：%v \n", err)
			return
		}
		str, _ := json.Marshal(chatInfo)
		fmt.Println(string(str))
	}
}


