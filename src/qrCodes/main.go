package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	rQrCode "github.com/tuotoo/qrcode"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", IndexAction)
	http.HandleFunc("/qrCode", ReadQrCode)
	log.Println("请打开页面: http://127.0.0.1:2345")
	http.ListenAndServe(":2345", nil)

	// readQrCode2();
	// writeQrCode();
	// readQrCode();
}

// 主页
func IndexAction(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("template/page/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(writer, nil)
}

type QrCode struct {
	QrContent string
}

// 读取二维码
func ReadQrCode(writer http.ResponseWriter, request *http.Request) {
	//判断请求方式
	if request.Method == "POST" {
		//设置内存大小
		request.ParseMultipartForm(64 << 20)
		//获取上传的第一个文件
		file, _, _ := request.FormFile("qrFile")
		// 读取文件
		qrmatrix, err := rQrCode.Decode(file)
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		log.Println("获取到二维码内容：", qrmatrix.Content)

		t, err := template.ParseFiles("template/page/qrCode.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(writer, QrCode{QrContent: qrmatrix.Content})
	} else {
		//解析模板文件
		t, _ := template.ParseFiles("template/page/qrCode.html")
		//输出文件数据
		t.Execute(writer, nil)
	}
}

// 写二维码
func writeQrCode() {
	// 写二维码
	err := qrcode.WriteFile("https://funnyzpc.cnblogs.com", qrcode.Medium, 256, "D:/tmp/cnblogs.png")
	if err != nil {
		fmt.Println(err)
	}
}

// 读二维码
func readQrCode() {
	file, error := os.Open("D:/tmp/cnblogs.png")
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	defer file.Close()
	qrmatrix, err := rQrCode.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}
func readQrCode2() {
	/*
		results, err := r2QrCode.GetDataFromPNG("C:/Users/Administrator/Desktop/my.png")
		if err != nil {
			panic(err)
		}
		for _, result := range results{
			fmt.Printf("Symbol Type: %s, Data %s", result.SymbolType, result.Data )
		}
	*/
}
