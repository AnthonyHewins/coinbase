package tests

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/AnthonyHewins/coinbase"
	"gopkg.in/yaml.v3"
)

const keyFile = "key.yaml"

var singleton config

type config struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

func testClient() *coinbase.Client {
	c, err := coinbase.NewClient(coinbase.DefaultProdURL, singleton.Key, singleton.Secret, &http.Client{Timeout: time.Second * 3})

	if err != nil {
		log.Fatalf("failed creating client: %s", err)
	}
	return c
}

func TestMain(m *testing.M) {
	if os.Getenv("INTEGRATION") == "" {
		return
	}
	log.Println("$INTEGRATION set, running integration tests")

	buf, err := os.ReadFile(keyFile)
	if err != nil {
		log.Fatalf(
			"Failed reading %s: make sure you have a tests/key.yaml file to run integration tests.\n"+
				"See key.template.yaml for what to fill out, it requires making coinbase API keys",
			keyFile,
		)
	}

	var c config
	if err := yaml.Unmarshal(buf, &c); err != nil {
		log.Fatalf("failed parsing YAML: %s", err)
	}

	if c.Key == "" {
		log.Fatalf("missing coinbase key variable in your config: was the empty string")
	}

	if c.Secret == "" {
		log.Fatalf("missing coinbase key secret in your config: was the empty string")
	}

	singleton = c
	m.Run()
}
