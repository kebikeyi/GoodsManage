package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	fmt.Println("this is main")
	var nick = this.GetSession("uname")
	this.Data["nick"] = nick
	fmt.Println(nick)
	// cookie, err := this.cookie("id")
	// this.Data["nick"] = cookie.uname
	this.Layout = "main/index.html"
	this.TplNames = "login/login.html"
}

func (this *MainController) TestService() {
	var params = make(map[string]string)
	params["msg"] = "dlm"
	postStr := CreateSOAP11Xml("http://qhd.qichelive.com/", "Testing", params)
	fmt.Println(postStr)

	// postStr = "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	// postStr += "<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">"
	// postStr += "  <soap:Body>"
	// postStr += "    <Testing xmlns=\"http://qhd.qichelive.com/\">"
	// postStr += "      <msg>string</msg>"
	// postStr += "    </Testing>"
	// postStr += "  </soap:Body>"
	// postStr += "</soap:Envelope>"

	fmt.Println("-------post service------->")
	var rst = PostWebService("http://demo.qichevm.com/api/vmservice.asmx", "Testing", postStr)
	this.Ctx.WriteString(rst)
}

//POST到webService
func PostWebService(url string, method string, value string) string {
	res, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(value)))
	//这里随便传递了点东西
	if err != nil {
		fmt.Println("post error", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	//取出主体的内容
	if err != nil {
		fmt.Println("read error", err)
	}
	res.Body.Close()
	fmt.Printf("result----%s", data)
	//return ByteToString(data)
	return string(data)
}
func CreateSOAP11Xml(nameSpace string, methodName string, params map[string]string) string {

	soapBody := "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	soapBody += "<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">"
	soapBody += "  <soap:Body>"
	soapBody += "    <" + methodName + " xmlns=\"" + nameSpace + "\">"
	for k, v := range params {
		soapBody += "      <" + k + ">" + v + "</" + k + ">"
	}
	//soapBody += "      <" + paramName + ">" + valueStr + "</" + paramName + ">"

	soapBody += "    </" + methodName + ">"
	soapBody += "  </soap:Body>"
	soapBody += "</soap:Envelope>"

	return soapBody
}
func CreateSOAP12Xml(nameSpace string, methodName string, valueStr string) string {
	soapBody := "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	soapBody += "<soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">"
	soapBody += "<soap12:Body>"
	soapBody += "<" + methodName + " xmlns=\"" + nameSpace + "\">"
	//以下是具体参数
	soapBody += "<Data>" + valueStr + "</Data>"
	soapBody += "</" + methodName + "></soap12:Body></soap12:Envelope>"
	return soapBody
}
