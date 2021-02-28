package mysys

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func MyExec() {
	// 解释:lsCmd.Output()返回字节数组,grepCmd.StdinPipe()返回一个reader,cmd.Stdin = in把buffer赋值给输入
	// writeString()
	// printDate()
	// grepData()
	lsData()
}

func writeString() {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("sh")
	cmd.Stdin = in
	go func() {
		in.WriteString("echo hello world > test.txt\n")
		in.WriteString("exit\n")
	}()
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}

func printDate() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
}

func grepData() {
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
}

func lsData() {
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
