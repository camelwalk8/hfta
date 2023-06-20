package main

import (
	"errors"
	"os"
	"testing"
)

func Test_ReadConfig(t *testing.T) {

	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"

	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Fatal(e)
	}

	got := c.TransferFor
	want := "emailJson"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func Test_ExistanceOfOneFileFilter(t *testing.T) {

	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"

	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Fatal(e)
	}

	got := c.Spec.FileTransferSpec.FileFilterRegEx[0]
	want := "^[0-9a-f]{12}4[0-9a-f]{3}[89ab][0-9a-f]{15}.md5"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func Test_ExistanceOfLogFileDir(t *testing.T) {
	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"

	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Errorf(e.Error())
	}

	want := c.Spec.LogFileDir
	if got, e := Exists(want); !got {
		t.Errorf(e.Error())
	}

}

func Test_ExistanceOfSrcFileDir(t *testing.T) {
	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"

	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Errorf(e.Error())
	}

	want := c.Spec.FileTransferSpec.SourceDir
	if got, e := Exists(want); !got {
		t.Errorf(e.Error())
	}

}

func Test_ExistanceOfDstFileDir(t *testing.T) {
	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"

	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Errorf(e.Error())
	}

	want := c.Spec.FileTransferSpec.DestinationDir
	if got, e := Exists(want); !got {
		t.Errorf(e.Error())
	}

}

func Test_MissingLogFileDir(t *testing.T) {

	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"

	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Errorf(e.Error())
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("Cannot get hold of current working dir, %v", err)
	}

	want := pwd + "/logs"
	got := c.Spec.LogFileDir
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
		t.Fail()
	}

	_, err = os.Stat(want)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			t.Errorf("%v does not exists", want)
			t.Fail()
		} else {
			t.Errorf("Missing default dir, %v", err)
			t.Fail()
		}
	}

}

func Test_MissingDLQFDir(t *testing.T) {

	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"

	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Fatal(e)
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Cannot get hold of current working dir, %v", err)
	}

	want := pwd + "/dlq"
	got := c.Spec.DlqDir
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
		t.Fail()
	}

	_, err = os.Stat(want)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			t.Errorf("%v does not exists", want)
			t.Fail()
		} else {
			t.Fatalf("Missing default dir, %v", err)
		}
	}

}

func Test_DefaultCheckSumWhenMissing(t *testing.T) {
	ymlconfig := "/home/auser/go/src/hfta/hfta-email-json/configs/hfta_app_config.yml"
	c, e := getApplicationConfiguration(ymlconfig)
	if e != nil {
		t.Fatal(e)
	}
	want := "MD5"
	got := c.Spec.FileTransferSpec.NewCheckSum.CheckSumType

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

// package main

// import (
// "fmt"
// "gopkg.in/yaml.v2"
// )

// func main() {
// // Read the YAML file.
// data, err := yaml.ReadFile("config.yml")
// if err != nil {
// fmt.Println(err)
// return
// }

// // Print the configuration.
// fmt.Println("configVer:", data["configVer"])
// fmt.Println("appName:", data["appName"])
// fmt.Println("appDesc:", data["appDesc"])
// fmt.Println("appPod:", data["appPod"])
// fmt.Println("appType:", data["appType"])
// fmt.Println("transferFor:", data["transferFor"])
// fmt.Println("batchSize:", data["spec"].(map[string]interface{})["batchSize"])
// fmt.Println("maxReprocessTime:", data["spec"].(map[string]interface{})["maxReprocessTime"])
// fmt.Println("archive:", data["spec"].(map[string]interface{})["archive"])
// fmt.Println("logFileDir:", data["spec"].(map[string]interface{})["logFileDir"])
// fmt.Println("dlqDir:", data["spec"].(map[string]interface{})["dlqDir"])
// fmt.Println("archiveDir:", data["spec"].(map[string]interface{})["archiveDir"])
// fmt.Println("fileTransferSpec:")
// fmt.Println("  fileFilterRegEx:")
// for _, v := range data["spec"].(map[string]interface{})["fileTransferSpec"].([]interface{}) {
//     fmt.Println("    -", v.(map[string]interface{})["fileFilterRegEx"])
// }
// fmt.Println("  sourceDir:", data["spec"].(map[string]interface{})["fileTransferSpec"].(map[string]interface{})["sourceDir"])
// fmt.Println("  destinationDir:", data["spec"].(map[string]interface{})["fileTransferSpec"].(map[string]interface{})["destinationDir"])
// fmt.Println("  newCheckSum:")
// fmt.Println("    checkSumType:", data["spec"].(map[string]interface{})["fileTransferSpec"].(map[string]interface{})["newCheckSum"].(map[string]interface{})["checkSumType"])
