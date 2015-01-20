package util

import (
	"os/exec"
	"fmt"
	"strings"
	//	"bytes"
	"os"
	"bufio"
)

type JubatusProcess struct {
	cmd *exec.Cmd
	Port int
}

func jubatus_keep(target exec.Cmd) {
	for {
		target.Wait()
		target.Start()
	}
}

type JubatusProcessError struct {
	message string
}
func (j *JubatusProcessError) Error() string {
	return fmt.Sprintf("Jubatus Process Error: %s", j.message)
}

func (j *JubatusProcess) Kill() error {
	if err := j.cmd.Process.Kill(); err != nil {
		return err
	}
	return nil
}

func NewJubatusProcess(command string, filepath string) (*JubatusProcess, error) {
	port := 9200
	for {
		cmd := exec.Command(command, "-f", filepath, "-p", fmt.Sprintf("%d", port))
	cmd.Stderr = os.Stderr

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}

		if err := cmd.Start(); err != nil {
			return nil, err
		}

		r := bufio.NewReader(stdout)
		for {
			byteline, _, err := r.ReadLine()
			if err != nil {
				return nil, err
			}
			line := string(byteline)
			if strings.Contains(line, "RPC server startup") {
				return &JubatusProcess{cmd, port}, nil
			} else if strings.Contains(line, "server failed to start") {
				port += 1
				break
			}
		}
	}
}


/*
func main() {
	proc, err := NewJubatusProcess("hoge")
	if err != nil {
		panic(err)
	}
	defer proc.Kill()
	client, err := proc.GetClient()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i += 1 {
		go func() {
			client, err := proc.GetClient()
			if err != nil {
				panic(err)
			}

			for {
				d := jubatus.NewDatum()
				d.AddNum(fmt.Sprintf("f-%d", i), float64(i))
				client.SimilarRowFromDatum(d, 10)
			}
		}()
	}
	for cnt := 0; cnt < 10; cnt += 1{
		d := jubatus.NewDatum()
		d.AddNum(fmt.Sprintf("f-%d", cnt), float64(cnt))
		client.UpdateRow(fmt.Sprintf("row-%d", cnt), d)
	}
	time.Sleep(10 * 1000 * 1000 * 1000)
}
*/
