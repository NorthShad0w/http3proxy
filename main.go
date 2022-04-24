package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"syscall"

	h3pclient "git.cyru1s.com/cyru1s/http3proxy/client"
)

func main() {
	file, _ := os.Open("./h3p.conf")
	defer file.Close()

	scanner := bufio.NewScanner(file)
    scanner.Scan()
    secret := scanner.Text()

    //Use a secret to enhence opsec
    if secret != "C{ZM2<%4H!)$kQ8cuaV?"{
        printErr()
        os.Exit(0)
    }

    scanner.Scan()
    cdn_ip := scanner.Text()

    scanner.Scan()
    sni_name := scanner.Text()

    scanner.Scan()
    host_name := scanner.Text()

    //To do 
    //delete the config file

	//fork and run enhence opsec
	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if id == 0 {
		//child process
		k := h3pclient.StartClient(cdn_ip,sni_name,host_name)
		fmt.Println(k)

	} else {
		os.Exit(0)
	}

}

func printErr() {
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("'%s' 不是内部或外部命令，也不是可运行的程序或批处理文件。\r\n", os.Args[0])
	case "linux":
		fmt.Printf("%s: line 6: ��e�WDT# : No such file or directory\r\n", os.Args[0])
		fmt.Printf("%s: line 15: syntax error near unexpected token `('\r\n", os.Args[0])
	default:
	}
}