package parser

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	testDir := "/tmp"
	tmpfile, err := ioutil.TempFile(testDir, "config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.WriteString(`
		{
			"controller": {
		
			},
		
			"logger": {
				
			},
			
			"main": {
				"TLSCert": ".private/cert.pem",
				"TLSKey": ".private/key.pem"
			},
			"model": {
				
					},
				
			"view": {
						
			}
		}
	`)
	if err != nil {
		log.Fatal(err)
	}

	TLSCert := ".private/cert.pem"
	testConfig := Config(tmpfile.Name())
	if testConfig.Main.TLSCert != TLSCert {
		t.Errorf("TLSCert not set. got: %s, want: %s", testConfig.Main.TLSCert, TLSCert)
	}
}
