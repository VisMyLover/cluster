package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

//filepath: 要编译的文件的路径
func build(filepath string) {
	_ = os.Setenv("CGO_ENABLED", "0")
	_ = os.Setenv("GOARCH", "amd64")
	_ = os.Setenv("GOOS", "linux")

	arg := []string{"build", filepath}
	if err := exec.Command("go", arg...).Run(); err != nil {
		fmt.Println("编译失败:", err)
	} else {
		fmt.Println("编译成功")
	}
}

func main() {

	//build(`./main.go`)
	//build(`C:/Users/王雄/go/src/git.ucloudadmin.com/uaek/hamster/main.go`)
	//build(`C:/Users/王雄/go/src/git.ucloudadmin.com/kun/operator-server/main.go`)
	build(`C:/Users/王雄/go/src/git.ucloudadmin.com/leesin/cluster/main.go`)
	getTime()
}

func getTime() {
	nowTime := time.Now()

	t := nowTime.Format("2006-01-02 15:04:05")

	fmt.Println(t)
}
