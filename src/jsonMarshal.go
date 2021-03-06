package main

import(
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type PortMapping struct {
	ContainerPort int
	HostPort int
	ServicePort int
	Protocol string
}

type Docker struct {
	Image string
	Network string
	Privileged bool
	ForcePullImage bool
	PortMappings []PortMapping
}

//这一块无法解析
type Container struct  {
	Type string `json:"type,omitempty"`
	Docker `docker:"type:json",json:"-"`
}

type HealthCheck struct  {
	Protocol string
	GracePeriodSeconds int
	IntervalSeconds int
	PortIndex int
	TimeoutSeconds int
	MaxConsecutiveFailures int
}

type Application struct {
	Id string
	Cpus float64
	Instances int
	Mem int
	Ports []int
	Container map[string]interface{}
	HealthChecks []HealthCheck
}

func main() {
	var app Application
	js := `{
"id": "docker-postgresql-0100",
"cpus": 0.3,
"instances": 1,
"mem": 300,
"ports": [
        5432
    ],
"container":{
    "type": "DOCKER",
    "docker":{
        "image": "frodenas/postgresql:latest",
        "network": "BRIDGE",
        "privileged": false,
        "forcePullImage": true,
        "portMappings": [
            {
                 "containerPort": 5432,
                 "hostPort": 0,
                 "servicePort": 5432,
                 "protocol": "tcp"
            }
        ]
     }
  },
 "healthChecks": [
     {
            "protocol": "TCP",
            "gracePeriodSeconds": 3,
            "intervalSeconds": 5,
            "portIndex": 0,
            "timeoutSeconds": 5,
            "maxConsecutiveFailures": 3
        }
  ]
}`

	jsh , err := simplejson.NewJson([]byte(js))

	if err != nil {
		fmt.Println("err")
	}

	fmt.Println(jsh.Get("container").Get("type"))

	json.Unmarshal([]byte(js), &app)
	fmt.Println(app)
}
