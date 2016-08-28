package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("domainfinder/lib/synonyms"),
	exec.Command("domainfinder/lib/sprinkle"),
	exec.Command("domainfinder/lib/coolify"),
	exec.Command("domainfinder/lib/domainify"),
	exec.Command("domainfinder/lib/available"),
}

func main() {
	// 標準入力
	cmdChain[0].Stdin = os.Stdin

	// 標準出力
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}
		nextCmd.Stdin = stdout
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {

			// log.Fatallnだとdeferの処理はコールされない.
			log.Panicln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Panicln(err)
		}
	}

}
