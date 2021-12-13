package WeWorkFinanceSDK

type Client interface {
	// GetChatData 拉取聊天记录
	//
	// @param [in]  seq             从指定的seq开始拉取消息，注意的是返回的消息从seq+1开始返回，seq为之前接口返回的最大seq值。首次使用请使用seq:0
	// @param [in]  limit           一次拉取的消息条数，最大值1000条，超过1000条会返回错误
	// @param [in]  proxy           使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
	// @param [in]  passwd          代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
	// @param [in]  timeout         超时时间，单位秒
	//
	// @return chatDatas       返回本次拉取消息的数据，slice结构体.内容包括errcode/errmsg，以及每条消息内容。示例如下：
	// {"errcode":0,"errmsg":"ok","chatdata":[{"seq":196,"msgid":"CAQQ2fbb4QUY0On2rYSAgAMgip/yzgs=","publickey_ver":3,"encrypt_random_key":"ftJ+uz3n/z1DsxlkwxNgE+mL38H42/KCvN8T60gbbtPD+Rta1hKTuQPzUzO6Hzne97MgKs7FfdDxDck/v8cDT6gUVjA2tZ/M7euSD0L66opJ/IUeBtpAtvgVSD5qhlaQjvfKJc/zPMGNK2xCLFYqwmQBZXbNT7uA69Fflm512nZKW/piK2RKdYJhRyvQnA1ISxK097sp9WlEgDg250fM5tgwMjujdzr7ehK6gtVBUFldNSJS7ndtIf6aSBfaLktZgwHZ57ONewWq8GJe7WwQf1hwcDbCh7YMG8nsweEwhDfUz+u8rz9an+0lgrYMZFRHnmzjgmLwrR7B/32Qxqd79A==","encrypt_chat_msg":"898WSfGMnIeytTsea7Rc0WsOocs0bIAerF6de0v2cFwqo9uOxrW9wYe5rCjCHHH5bDrNvLxBE/xOoFfcwOTYX0HQxTJaH0ES9OHDZ61p8gcbfGdJKnq2UU4tAEgGb8H+Q9n8syRXIjaI3KuVCqGIi4QGHFmxWenPFfjF/vRuPd0EpzUNwmqfUxLBWLpGhv+dLnqiEOBW41Zdc0OO0St6E+JeIeHlRZAR+E13Isv9eS09xNbF0qQXWIyNUi+ucLr5VuZnPGXBrSfvwX8f0QebTwpy1tT2zvQiMM2MBugKH6NuMzzuvEsXeD+6+3VRqL"}]}
	//
	GetChatData(seq uint64, limit uint64, proxy string, passwd string, timeout int) ([]ChatData, error)

	// DecryptData 解析密文.企业微信自有解密内容
	//
	// @param [in]  encrypt_key, getchatdata返回的encrypt_random_key,使用企业自持对应版本秘钥RSA解密后的内容
	// @param [in]  encrypt_msg, getchatdata返回的encrypt_chat_msg
	// @param [out] msg, 解密的消息明文
	//
	// @return 返回是否调用成功
	//      0   - 成功
	//  	!=0 - 失败
	//
	DecryptData(encryptRandomKey string, encryptMsg string) (msg ChatMessage, err error)

	// GetMediaData 拉取媒体消息函数
	//
	// Return值=0表示该API调用成功
	// @param [in]  sdk				NewSdk返回的sdk指针
	// @param [in]  sdkFileid		从GetChatData返回的聊天消息中，媒体消息包括的sdkfileid
	// @param [in]  proxy			使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
	// @param [in]  passwd			代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
	// @param [in]  indexbuf		媒体消息分片拉取，需要填入每次拉取的索引信息。首次不需要填写，默认拉取512k，后续每次调用只需要将上次调用返回的outindexbuf填入即可。
	// @param [in]  timeout			超时时间，单位秒
	// @param [out] media_data		返回本次拉取的媒体数据.MediaData结构体.内容包括data(数据内容)/outindexbuf(下次索引)/is_finish(拉取完成标记)
	//
	// @return 返回是否调用成功
	//      0   - 成功
	//  	!=0 - 失败
	//
	GetMediaData(indexBuf string, sdkFileId string, proxy string, passwd string, timeout int) (*MediaData, error)

	// Free 释放 Client
	// 释放 C 指针,避免内存泄漏
	Free()
}
