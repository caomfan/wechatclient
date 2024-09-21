package services

import (
	"fmt"
	"wechatclient/manager/wx_client_mgr"
)

func HeartBeat(sessionId string) (interface{}, error) {
	// 走mgr 来查询账户然后走mgr去调用功能
	wxClient := wx_client_mgr.WxClientMgr.Get(sessionId)
	// 缓存中没有找到账号, 账号未登录
	if wxClient == nil {
		fmt.Printf("没有找到 %s 会话绑定的账号\n", sessionId)
		return nil, fmt.Errorf("没有找到 %s 会话绑定的账号", sessionId)
	}

	response, err := wxClient.HeartBeat()
	if err != nil {
		return "", err
	}
	return response, nil
}
