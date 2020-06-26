package main

import (
	"encoding/json"
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)
type FoodChainCode struct{	
}
func (a *FoodChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (a *FoodChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fn,args := stub.GetFunctionAndParameters()
    if fn == "addProInfo"{
        return a.addProInfo(stub,args)
    } else if fn == "addIngInfo"{
        return a.addIngInfo(stub,args)
    } else if fn == "getFoodInfo"{
        return a.getFoodInfo(stub,args)
    }else if fn == "addLogInfo"{
        return a.addLogInfo(stub,args)
    }else if fn == "getProInfo"{
        return a.getProInfo(stub,args)
    }else if fn == "getLogInfo"{
        return a.getLogInfo(stub,args)
    }else if fn == "getIngInfo"{
        return a.getIngInfo(stub,args)
    }else if fn == "getLogInfo_l"{
        return a.getLogInfo_l(stub,args)
    }else if fn=="updateLogInfo"{
        return a.updateLogInfo(stub,args)

    }


    return shim.Error("Recevied unkown function invocation")
}

func (a *FoodChainCode) addProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error 
    var FoodInfos FoodInfo

    if len(args)!=10{//如果参数没有10个，报错
        return shim.Error("Incorrect number of arguments.")
    }
    FoodInfos.FoodID = args[0]
    if FoodInfos.FoodID == ""{
        return shim.Error("FoodID can not be empty.")
    }
    
    
    FoodInfos.FoodProInfo.FoodName = args[1]
    FoodInfos.FoodProInfo.FoodSpec = args[2]
    FoodInfos.FoodProInfo.FoodMFGDate = args[3]
    FoodInfos.FoodProInfo.FoodEXPDate = args[4]
    FoodInfos.FoodProInfo.FoodLOT = args[5]
    FoodInfos.FoodProInfo.FoodQSID = args[6]
    FoodInfos.FoodProInfo.FoodMFRSName = args[7]
    FoodInfos.FoodProInfo.FoodProPrice = args[8]
    FoodInfos.FoodProInfo.FoodProPlace = args[9]
    ProInfosJSONasBytes,err := json.Marshal(FoodInfos)
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(FoodInfos.FoodID,ProInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}

func(a *FoodChainCode) addIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
        
    var FoodInfos FoodInfo
    var IngInfoitem IngInfo

    if  (len(args)-1)%2 != 0 || len(args) == 1{
        return shim.Error("Incorrect number of arguments")
    }

    FoodID := args[0]
    for i :=1;i < len(args);{
        IngInfoitem.IngID = args[i]
        IngInfoitem.IngName = args[i+1]
        FoodInfos.FoodIngInfo = append(FoodInfos.FoodIngInfo,IngInfoitem)
        i = i+2
    }
    
    
    FoodInfos.FoodID = FoodID
  /*  FoodInfos.FoodIngInfo = foodIngInfo*/
    IngInfoJsonAsBytes,err := json.Marshal(FoodInfos)
    if err != nil {
    return shim.Error(err.Error())
    }

    err = stub.PutState(FoodInfos.FoodID,IngInfoJsonAsBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
        
}

func(a *FoodChainCode) addLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    var err error
    var FoodInfos FoodInfo

    if len(args)!=11{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodInfos.FoodID = args[0]
    if FoodInfos.FoodID == ""{
        return shim.Error("FoodID can not be empty.")
    }
    FoodInfos.FoodLogInfo.LogDepartureTm = args[1]
    FoodInfos.FoodLogInfo.LogArrivalTm = args[2]
    FoodInfos.FoodLogInfo.LogMission = args[3]
    FoodInfos.FoodLogInfo.LogDeparturePl = args[4]
    FoodInfos.FoodLogInfo.LogDest = args[5]
    FoodInfos.FoodLogInfo.LogToSeller = args[6]
    FoodInfos.FoodLogInfo.LogStorageTm = args[7]
    FoodInfos.FoodLogInfo.LogMOT = args[8]
    FoodInfos.FoodLogInfo.LogCopName = args[9]
    FoodInfos.FoodLogInfo.LogCost = args[10]
    FoodInfos.FoodLogInfo.FoodNum = args[11]
    
    LogInfosJSONasBytes,err := json.Marshal(FoodInfos)
    if err != nil{
        return shim.Error(err.Error())
    } 
    err = stub.PutState(FoodInfos.FoodID,LogInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}

func(a *FoodChainCode) getFoodInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(FoodID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    var foodAllinfo FoodAllInfo

    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err :=resultsIterator.Next()
        if err != nil {
             return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodProInfo.FoodName !=""{
            foodAllinfo.FoodProInfo = FoodInfos.FoodProInfo
        }else if FoodInfos.FoodIngInfo != nil{
            foodAllinfo.FoodIngInfo = FoodInfos.FoodIngInfo
        }else if FoodInfos.FoodLogInfo.LogMission !=""{
            foodAllinfo.FoodLogInfo = append(foodAllinfo.FoodLogInfo,FoodInfos.FoodLogInfo)
        }

    }
    
    jsonsAsBytes,err := json.Marshal(foodAllinfo)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(jsonsAsBytes)
}
func(a *FoodChainCode) getProInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(FoodID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var foodProInfo ProInfo

    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err :=resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodProInfo.FoodName != ""{
            foodProInfo = FoodInfos.FoodProInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(foodProInfo)   
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    if len(args) !=1{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(FoodID)

    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var foodIngInfo []IngInfo
    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err := resultsIterator.Next()
        if err != nil{
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodIngInfo != nil{
            foodIngInfo = FoodInfos.FoodIngInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(foodIngInfo)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

    var LogInfos []LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    FoodID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(FoodID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

   
    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodLogInfo.LogMission != ""{
            LogInfos = append(LogInfos,FoodInfos.FoodLogInfo)
        }
    }
    jsonsAsBytes,err := json.Marshal(LogInfos)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}
func getLogInfo(stub shim.ChaincodeStubInterface, FoodNum string) (LogInfo, bool) {
    var Log LogInfo
    b,err:=stub.GetState(FoodNum)
    if err!=nil{
        return Log,false
    }
    if b==nil{
        return Log,false
    }
    err=json.Unmarshal(b,&Log)
    if err!=nil{
        return Log,false
    }
    return Log,true
}
func PutLog(stub shim.ChaincodeStubInterface,Log LogInfo)([]byte,bool){
    b,err:=json.Marshal(Log)
    if err!=nil{
        return nil,false
    }
    err=stub.PutState(Log.FoodNum,b)
    if err!=nil {
        return nil,false
    }
    return b,true
}
func(a *FoodChainCode)updateLogInfo(stub shim.ChaincodeStubInterface,args [] string) pb.Response{

    if len(args)!=2{
        return shim.Error("给定的参数不符合要求")
    }
    var loginfo LogInfo

    //err:=json.Unmarshal([]byte(args[0],&loginfo))
    //if err!=nil{
   //     return shim.Error("反序列化信息失败")
   // }
    result,b1:=getLogInfo(stub,loginfo.FoodNum)
    if !b1{
        return shim.Error("根据货物编号查询信息发生错误")
    }
    result.LogDepartureTm=loginfo.LogDepartureTm
    result.LogArrivalTm=loginfo.LogArrivalTm
    result.LogMission=loginfo.LogMission
    result.LogDeparturePl=loginfo.LogDeparturePl
    result.LogDest=loginfo.LogDest
    result.LogToSeller=loginfo.LogToSeller
    result.LogStorageTm=loginfo.LogStorageTm
    result.LogMOT=loginfo.LogMOT
    result.LogCopName=loginfo.LogCopName
    result.LogCost=loginfo.LogCost
    result.FoodNum=loginfo.FoodNum

    _,b1=PutLog(stub,result)
    if !b1{
        return shim.Error("保存信息时发生错误")
    }
    err:=stub.SetEvent(args[1],[]byte{})
    if err!=nil{
        return shim.Error(err.Error())
    }
    return shim.Success([]byte("信息更新成功"))

}

func(a *FoodChainCode) getLogInfo_l(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    var Loginfo LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    FoodID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(FoodID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodLogInfo.LogMission != ""{
           Loginfo = FoodInfos.FoodLogInfo
           continue 
       }
    }
    jsonsAsBytes,err := json.Marshal(Loginfo)
    if err != nil{
        return shim.Error(err.Error ())
    }
    return shim.Success(jsonsAsBytes)
}


func main(){
     err := shim.Start(new(FoodChainCode))
     if err != nil {
         fmt.Printf("Error starting Food chaincode: %s ",err)
     }
}
