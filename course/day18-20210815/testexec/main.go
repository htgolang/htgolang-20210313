package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// ping baidu
	// exec.Cmd("ping baidu")
	fmt.Println(exec.LookPath("ping"))
	fmt.Println(exec.LookPath("ping2"))

	cmd := exec.Command("ping", "baidu.com")

	fmt.Println("run before")
	err := cmd.Run() // 有Run，等待执行结束
	fmt.Println(err)

	fmt.Println("output before")
	output, err := cmd.Output() // 无Run时 默认调用Run并等待结束
	fmt.Println(string(output), err)

	// ps -ef | grep xxx|awk 'xxx'

	// bash -c "xxx"

	// systemctl reload xxxxconst

	cmd = exec.Command("ping", "baidu.com")
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()

	var result = bytes.NewBufferString("")

	if err := cmd.Start(); err == nil {
		fmt.Println("started")

		// 对于输出和输入流必须放在Start之后Wait之前
		io.Copy(result, stdout)
		io.Copy(result, stderr)

		err := cmd.Wait() //
		fmt.Println("complated", err)

		result.WriteTo(os.Stdout)

	}

}
