package test

import (
	"fmt"
	_ "glory-golang/routers"
	"glory-golang/utils"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {

	//logs.Info("this %s cat is %v years old", "yellow", 3)
	//logs.Debug("this %s cat is %v years old", "yellow", 3)
	//logs.Error("this %s cat is %v years old", "yellow", 3)
	//fmt.Print("1111")
	//
	//uuid := utils.GetUUID()
	//fmt.Print(uuid)

	time := utils.GetCurrentTime()
	fmt.Println(time)
	timeString := utils.GetCurrentTimeString()
	fmt.Println(timeString)
	//r, _ := http.NewRequest("GET", "/", nil)
	//w := httptest.NewRecorder()
	//beego.BeeApp.Handlers.ServeHTTP(w, r)
	//
	//beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
	//
	//Convey("Subject: Test Station Endpoint\n", t, func() {
	//	Convey("Status Code Should Be 200", func() {
	//		So(w.Code, ShouldEqual, 200)
	//	})
	//	Convey("The Result Should Not Be Empty", func() {
	//		So(w.Body.Len(), ShouldBeGreaterThan, 0)
	//	})
	//})
}
