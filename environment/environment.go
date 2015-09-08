package environment

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var Env Environment

func init() {
	etcdmachines := flag.String("etcd", "", "The etcd machines.")
	sitesdirectory := flag.String("directory", "", "Directory for sites.")
	flag.Parse()
	if *etcdmachines == "" {
		var err error
		etcdmachines, err = getEnvValue("MONGOLAR_ETCD_MACHINES")
		if err != nil {
			log.Fatal(err)
		}
	}
	if *sitesdirectory == "" {
		var err error
		sitesdirectory, err = getEnvValue("MONGOLAR_SITES_DIRECTORY")
		if err != nil {
			log.Fatal(err)
		}
	}
	Env = Environment{
		EtcdMachines:   strings.Split(*etcdmachines, "|"),
		SitesDirectory: *sitesdirectory,
	}
	Env.refresh()
}

type Environment struct {
	EtcdMachines   []string
	SitesDirectory string
}

func (e *Environment) refresh() {
	go func() {
		for _ = range time.Tick(10 * time.Second) {
			etcdmachines, err := getEnvValue("MONGOLAR_ETCD_MACHINES")
			if err != nil || *etcdmachines != "" {
				//TODO: ERROR handling needs to be added
				fmt.Println(err)
			} else {
				e.EtcdMachines = strings.Split(*etcdmachines, "|")
			}
		}
	}()
}

func getEnvValue(name string) (*string, error) {
	value := os.Getenv(name)
	if value == "" {
		return &value, fmt.Errorf("%v is not set, %v environment value is required.", name, name)
	}
	return &value, nil
}
