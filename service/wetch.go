package service

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"github.com/tddey01/aria2/config"
	"github.com/tddey01/aria2/model"
	"github.com/tddey01/aria2/utils"
	"strconv"
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
	if config.GetConfig().Watch.Enable {
		if err := BlockTotalCount(); err != nil { // 出块统计当天
			log.Error("携程死掉了", err.Error())
		}
	}
}

// 每天出块统计

func BlockTotalCount() (err error) {
	str := fmt.Sprintf(" %s", utils.TimeHMS())
	datacount, err := model.GetCount()
	if err != nil {
		return err
	}

	down, _ := strconv.Atoi(datacount[0].Downloaded) // 正在下载中
	count, _ := strconv.Atoi(datacount[0].Total)     // 总数
	act, _ := strconv.Atoi(datacount[0].Successful)  // 封装数量

	total := (float64(down) / float64(count)) * 100 // 下载占比
	actv := (float64(act) / float64(count)) * 100   // 封装百分比
	Totals := strconv.FormatFloat(total, 'f', 2, 64)
	Actvls := strconv.FormatFloat(actv, 'f', 2, 64)

	str += fmt.Sprintf("\n正在下载中 >>>>>>> ：%s \n已完成下载 >>>>>>> : %s \n下载进度百分比>>>>> : %v%% \n已完成封装总数 >>>> : %v \n封装进度百分比 >>>> : %v%%",
		datacount[0].Downloading, datacount[0].Downloaded, Totals, datacount[0].Successful, Actvls)
	log.Debug(" 发消息 企业微信你")
	restp := &Msg{
		Touser:  "@all",
		Msgtype: "text",
		Agentid: 1000004,
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
