package main


//food数据结构体
type FoodInfo struct{
	FoodID string `json:FoodID`                             //食品ID
	FoodProInfo ProInfo `json:FoodProInfo`                  //生产信息
	FoodIngInfo []IngInfo `json:FoodIngInfo`                  //配料信息
	FoodLogInfo LogInfo `json:FoodLogInfo`                  //物流信息
}
type FoodAllInfo struct{
	FoodID string `json:FoodId`
	FoodProInfo ProInfo `json:FoodProInfo`
	FoodIngInfo []IngInfo `json:FoodIngInfo`
	FoodLogInfo []LogInfo `json:FoodLogInfo`
}

//生产信息
type ProInfo struct{
	FoodName string `json:FoodName`                         //食品名称
	FoodSpec string `json:FoodSpec`                         //食品规格
	FoodMFGDate string `json:FoodMFGDate`                   //食品出产日期
	FoodEXPDate string `json:FoodEXPDate`                   //食品保质期
	FoodLOT string `json:FoodLOT`                           //食品批次号
	FoodQSID string `json:FoodQSID`                         //食品生产许可证编号
	FoodMFRSName string `json:FoodMFRSName`                 //食品生产商名称
	FoodProPrice string `json:FoodProPrice`                 //食品生产价格
	FoodProPlace string `json:FoodProPlace`                 //食品生产所在地

}

type IngInfo struct{
	IngID string `json:IngID`                               //配料ID
	IngName string `json:IngName`                           //配料名称
}

type LogInfo struct {
	LogDepartureTm string `json:LogDepartureTm` //出发时间
	LogArrivalTm   string `json:LogArrivalTm`   //到达时间
	LogMission     string `json:LogMission`     //处理业务(储存or运输)
	LogDeparturePl string `json:LogDeparturePl` //出发地
	LogDest        string `json:LogDest`        //目的地
	LogToSeller    string `json:LogToSeller`    //销售商
	LogStorageTm   string `json:LogStorageTm`   //存储时间
	LogMOT         string `json:LogMOT`         //运送方式
	LogCopName     string `json:LogCopName`     //物流公司名称
	LogCost        string `json:LogCost`
	FoodNum        string `json:FoodNum`     //HS商品编码

	//费用
}
