package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type BasicAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type StaticConfigs struct {
	Targets []string `yaml:"targets"`
}
type Job struct {
	JobName       string        `yaml:"job_name"`
	BasicAuth     BasicAuth     `yaml:"basic_auth"`
	StaticConfigs StaticConfigs `yaml:"static_configs"`
}

func main() {
	var job = Job{
		JobName:   "prometheus",
		BasicAuth: BasicAuth{"admin", "123@567"},
		StaticConfigs: StaticConfigs{
			Targets: []string{"localhost:9999"},
		},
	}
	bytes, err := yaml.Marshal(job)
	fmt.Println(string(bytes), err)

	ctx := `
job_name: prometheus1
basic_auth:
  username: admin1
  password: 123@abc
static_configs:
  x: xxx
  targets:
    - localhost:99991
  xxxx: yyyy
x: abc
`

	var job2 Job

	err = yaml.Unmarshal([]byte(ctx), &job2)
	fmt.Println(job2, err)

	err = yaml.UnmarshalStrict([]byte(ctx), &job2)
	fmt.Println(job2, err)
}
