package controllers

import (
	"github.com/astaxie/beego"
	"common/gokits/answerdata"
	"encoding/json"
)

type ValidationController struct {
	beego.Controller
}

func (lc *ValidationController) Get() {
	data, _ := json.Marshal(answerdata.NewAnswer(answerdata.OK, ""))
	lc.Data["result"] = data

}
