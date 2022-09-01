package utils

import (
	"bytes"
	"encoding/json"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
)

const (
	//发送消息使用导的url
	sendurl = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`
	//获取token使用导的url
	get_token = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=`
)

var requestError = xerrors.New("request error,check url or network")

type access_token struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

func Get_token(corpid, corpsecret string) (at access_token, err error) {
	resp, err := http.Get(get_token + corpid + "&corpsecret=" + corpsecret)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = requestError
		return
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf, &at)
	if at.Access_token == "" {
		err = xerrors.New("corpid or corpsecret error.")
	}
	return
}

type send_msg_error struct {
	Errcode int    `json:"errcode`
	Errmsg  string `json:"errmsg"`
}

func SendMsg(msgbody []byte, CorPid, CorpSecret string) error {

	token, err := Get_token(CorPid, CorpSecret)
	if err != nil {
		println(err.Error())
	}

	body := bytes.NewBuffer(msgbody)
	resp, err := http.Post(sendurl+token.Access_token, "application/json", body)
	if resp.StatusCode != 200 {
		return requestError
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var e send_msg_error
	err = json.Unmarshal(buf, &e)
	if err != nil {
		return err
	}
	if e.Errcode != 0 && e.Errmsg != "ok" {
		return xerrors.New(string(buf))
	}
	return nil
}
