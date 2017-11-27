package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	 "encoding/json"
	"github.com/imroc/req"
	"crypto/sha256"
	"time"
	"encoding/hex"
	"strconv"
	"github.com/flywithbug/qcloudsms"
	"math/rand"

)

type UserInfoModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	UserName	string 		`json:"username,omitempty" form:"username,omitempty"`
	Password	string  	`json:"password,omitempty" form:"password,omitempty"`
	CreateTime	int64  		`json:"createtime,omitempty" form:"createtime,omitempty"`
	UpdateTime	int64  		`json:"updatetime,omitempty" form:"updatetime,omitempty"`
	Sex 		int 		`json:"sex,omitempty" form:"sex,omitempty"`     //0默认未设置 1男，2女
	UserId 		string 		`json:"userid,omitempty" form:"userid,omitempty"`
	DepartName	string		`json:"departname,omitempty" form:"departname,omitempty"`
	Phone 		string		`json:"phone,omitempty" form:"phone,omitempty"`
	PhonePrefix 	string		`json:"phoneprefix,omitempty" form:"phoneprefix,omitempty"`
	Mail 		string		`json:"mail,omitempty" form:"mail,omitempty"`
	OldPassword 	string		`json:"oldpassword,omitempty" form:"oldpassword,omitempty"`
	Authtoken 	string		`json:"authtoken,omitempty" form:"authtoken,omitempty"`
	State 		int 		`json:"state,omitempty" form:"state,omitempty"`
}

type SMSTXParamModel struct {
	Uid  		int64 		`json:"uid,omitempty"`
	Signature	string  	`json:"sig,omitempty"`
	Time		int64  		`json:"time,omitempty"`
	Extend 		string 		`json:"extend" `
	Ext		string		`json:"ext,omitempty"`
	TelModel	telephone	`json:"tel,omitempty"`
	Tpl_id          int		`json:"tpl_id,omitempty"`
	Sign		string		`json:"sign,omitempty"`
	Params          [1]string	`json:"params,omitempty"`
}

type telephone struct {
	Code		string 			`json:"nationcode,omitempty"`
	Mobile		string 			`json:"mobile,omitempty"`
}

type telPara struct {

}

func main()  {

	TestSend()
}

func TestSend() {
	conf := qcloudsms.NewClientConfig()
	conf.AppID = "1400048815"
	conf.AppKey = "023328a2ffd090219ead91e7262ac155"
	client, err := qcloudsms.NewClient(conf)
	if err != nil {
		return
	}
	sms, err := qcloudsms.SMSService(client)
	if err != nil {
		return
	}
	ext := qcloudsms.SmsExt{}
	ext.Type = 0
	ext.NationCode = "86"

	resp, err := sms.Send("17602198956", "您的验证码是：1234 如非本人操作，请忽略本短信.(http://www.flywithme.top)",ext)
	if err != nil {
		return
	}
	fmt.Println(resp)

}



func testGetR()  {
	var user UserInfoModel
	r,_ := req.Get("http://localhost:8081/user/1000018")
	data ,_ := r.ToBytes()
	json.Unmarshal(data,&user)
	//r.ToJSON(&user)
	fmt.Println(r,user)
}

func postRequest()  {

	hash := sha256.New()
	mobile := "137617104" +
		"30"
	strRand := fmt.Sprintf("%06d", rand.Intn(999999))

	strTime := time.Now().Unix()
	fmt.Println(strTime)
	sig :="appkey=023328a2ffd090219ead91e7262ac155"+"&random="+strRand+"&time="+ strconv.FormatInt(strTime,10)+"&mobile="+mobile
	fmt.Println(sig)

	data2 := []byte(sig)


	hash.Write(data2)
	md := hash.Sum(nil)
	mdstr := hex.EncodeToString(md)
	fmt.Println(mdstr)

	tel := telephone{}
	tel.Code = "86"
	tel.Mobile = "17602198956"
	sms := SMSTXParamModel{}
	sms.Time = time.Now().Unix()
	sms.TelModel = tel
	sms.Signature = mdstr
	sms.Extend = ""
	//sms.Sign = "我在案发现场"
	sms.Tpl_id = 56229
	sms.Params[0] = "1234"
	fmt.Println(sms.Params)
	header := make(http.Header)

	header.Set("Content-Type", "application/json")

	resp ,_ := req.Post("https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=1400048815&random="+strRand, req.BodyJSON(&sms),header)                  // body => id=imroc&pwd=roc
	fmt.Println(resp)
}




func loginTest()  {

	//authHeader := req.Header{
	//	"Accept":        "application/json",
	//}
	var use UserInfoModel

	use.UserName =  "Jack008"
	use.Password =  "admin123"
	header := make(http.Header)

	header.Set("Content-Type", "application/json")

	resp ,_ := req.Post("http://localhost:8081/user/login", req.BodyJSON(&use),header)
	fmt.Println(resp)
}




func getRequest()  {
	client := &http.Client{}
	request ,_ := http.NewRequest("GET","http://localhost:8081/user/1000018",nil)
	request.Header.Set("Connection", "keep-alive")
	response, _ := client.Do(request)
	defer  response.Body.Close()

	var user UserInfoModel

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body,&user)

		fmt.Println(user)
	}
}