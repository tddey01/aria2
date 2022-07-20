package service

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"github.com/tddey01/aria2/model"
	"github.com/tddey01/aria2/utils"
)

func BlockStartNewTotal3() {
	spec := "01, 01, *, *, *, *" // 每天23 点 01 分
	c := cron.New()
	if err := c.AddFunc(spec, BlockStartNew); err != nil {
		log.Error("当天出块 统计失败  func error:", err.Error())
		return
	}
	log.Debug("开启任务计划")
	c.Start()
}

func BlockStartNew() {
	if err := BlockTotalCount(); err != nil { // 出块统计当天
		log.Error("携程死掉了", err.Error())
	}
}

// 每天出块统计

func BlockTotalCount() (err error) {

	str := fmt.Sprintf(" %s", utils.TimeHMS())
	datacount, err := model.GetCount()
	if err != nil {
		return err
	}

	str += fmt.Sprintf("\n下载中 >>>>：%s \n下载完成 >>: %s ", datacount[0].Downloading, datacount[0].Downloaded) //节点：f080468  有效算力: 6.432 PiB  今日块: 3  24h幸运值：80.00% 3日内块：2
	log.Debug(" 发消息 企业微信你")
	restp := &Msg{
		Touser:  "@all",
		Msgtype: "text",
		Agentid: 1000002,
		Text: &Message{
			Content: str,
		},
		Safe: 0,
	}
	resp, err := json.Marshal(restp)

	if err != nil {
		log.Error(err.Error())
		return
	}

	if err = utils.SendMsg(resp); err != nil {
		log.Error(err.Error())
		return
	}
	return nil
}

type Msg struct {
	Touser               string   `json:"touser"`
	Msgtype              string   `json:"msgtype"`
	Agentid              int      `json:"agentid"`
	Text                 *Message `json:"text"`
	Safe                 int      `json:"safe"`
	EnableIdTrans        int      `json:"enable_id_trans"`
	EnableDuplicateCheck int      `json:"enable_duplicate_check"`
}

type Message struct {
	Content string `json:"content"`
}
