package WeWorkFinanceSDK

type ChatDataResponse struct {
	Error
	ChatDataList []ChatData `json:"chatdata,omitempty"`
}

func (this ChatDataResponse) IsError() bool {
	return this.ErrCode != 0
}

type ChatData struct {
	Seq              uint64 `json:"seq,omitempty"`                // 消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1
	MsgId            string `json:"msgid,omitempty"`              // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	PublickeyVer     uint32 `json:"publickey_ver,omitempty"`      // 加密此条消息使用的公钥版本号。
	EncryptRandomKey string `json:"encrypt_random_key,omitempty"` // 使用publickey_ver指定版本的公钥进行非对称加密后base64加密的内容，需要业务方先base64 decode处理后，再使用指定版本的私钥进行解密，得出内容。
	EncryptChatMsg   string `json:"encrypt_chat_msg,omitempty"`   // 消息密文。需要业务方使用将encrypt_random_key解密得到的内容，与encrypt_chat_msg，传入sdk接口DecryptData,得到消息明文。
}
