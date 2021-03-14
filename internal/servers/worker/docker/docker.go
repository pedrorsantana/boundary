package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ahmetalpbalkan/dexec"
	docker "github.com/fsouza/go-dockerclient"
	"gopkg.in/yaml.v2"
)

// const (
// 	code = `
// secret:
//   output: "/home/teste.txt"
//   code:
//    - echo $(curl --location --request POST 'https://nightly.btcore.app/oauth/token' --header 'Content-Type: application/x-www-form-urlencoded' --header 'Authorization: Basic cm9vdDoxMTE4YWQ4OS05NDIwLTQ1YjQtYjFhYi0wYmUxYTk1ZTQ5MmQ=' --header 'Cookie: __cfduid=da78d7ed0575ff2c9bd24e1427ece5dee1614091357' --data-urlencode 'grant_type=password' --data-urlencode 'username=admin@btcore.app' --data-urlencode 'password=64fdbbf3-0e1d-4a77-ba2c-b232937328cb') > /home/teste.txt
//    - ls
// `
// )

// func main() {
// 	result, err := ReadCodeAndExecute(code)
// 	fmt.Println(result, err)
// }

type SingleOrMulti struct {
	Values []string
}

func (sm *SingleOrMulti) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		sm.Values = make([]string, 1)
		sm.Values[0] = single
	} else {
		sm.Values = multi
	}
	return nil
}

type Data struct {
	Secret map[string]SingleOrMulti
}

func ReadCodeAndExecute(code string) (string, error) {
	var t Data
	yaml.Unmarshal([]byte(code), &t)

	dPullImage()

	var executionCode []string
	var output string

	m, d, err := dRun()
	for k, v := range t.Secret {
		if k == "code" {
			executionCode = v.Values
		}
		if k == "output" {
			output = v.Values[0]
		}
	}

	str := strings.Join(executionCode, ";")
	if output != "" {
		str = str + ";" + fmt.Sprintf("cat %s", output)
	}

	result, err := dExecute(m, d, str)
	if err != nil {
		return "", err
	}

	return result, nil
}

func dPullImage() *exec.Cmd {
	cmd := exec.Command("docker", "pull", "ubuntu")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	return cmd
}

func dRun() (dexec.Execution, dexec.Docker, error) {
	cl, _ := docker.NewClient("unix:///var/run/docker.sock")
	d := dexec.Docker{cl}

	m, _ := dexec.ByCreatingContainer(docker.CreateContainerOptions{
		Config: &docker.Config{Image: "ubuntu"}})
	return m, d, nil
}

func dExecute(m dexec.Execution, d dexec.Docker, command string) (string, error) {
	//split := strings.Split(command, " ")
	cmd := d.Command(m, "sh", "-c", command)
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
