package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type ApplicationConfig struct {
	ConfigVer   string `yaml:"configVer" validate:"required"`
	AppName     string `yaml:"appName"`
	AppDesc     string `yaml:"appDesc"`
	AppPod      string `yaml:"appPod"`
	AppType     string `yaml:"appType"`
	TransferFor string `yaml:"transferFor"`
	Spec        struct {
		BatchSize        int    `yaml:"batchSize"`
		MaxReprocessTime int    `yaml:"maxReprocessTime"`
		Archive          bool   `yaml:"archive"`
		LogFileDir       string `yaml:"logFileDir"`
		DlqDir           string `yaml:"dlqDir"`
		ArchiveDir       string `yaml:"archiveDir"`
		FileTransferSpec struct {
			FileFilterRegEx []string `yaml:"fileFilterRegEx"`
			SourceDir       string   `yaml:"sourceDir"`
			DestinationDir  string   `yaml:"destinationDir"`
			NewCheckSum     struct {
				CheckSumType string `yaml:"checkSumType"`
			} `yaml:"newCheckSum"`
		} `yaml:"fileTransferSpec"`
	} `yaml:"spec"`
}

func (a *ApplicationConfig) IsValid() (bool, error) {
	if a.ConfigVer == "" {
		return false, errors.New("missing configVer from yml configuration")
	}

	if a.AppName == "" {
		return false, errors.New("missing appName from yml configuration")
	}

	if a.TransferFor == "" {
		return false, errors.New("missing transferFor from yml configuration")
	}

	if a.Spec.BatchSize == 0 {
		a.Spec.BatchSize = 200
	}

	if a.Spec.MaxReprocessTime == 0 {
		a.Spec.MaxReprocessTime = 5
	}

	if a.Spec.FileTransferSpec.NewCheckSum.CheckSumType == "" {
		a.Spec.FileTransferSpec.NewCheckSum.CheckSumType = "MD5"
	}

	pwd, err := os.Getwd()
	if err != nil {
		return false, fmt.Errorf("cannot get hold of current working dir, %w", err)
	}

	if b, err := Exists(a.Spec.FileTransferSpec.SourceDir); !b {
		return false, fmt.Errorf("cannot get hold of source dir, %w", err)
	}

	if b, err := Exists(a.Spec.FileTransferSpec.DestinationDir); !b {
		return false, fmt.Errorf("cannot get hold of destination dir, %w", err)
	}

	if len(a.Spec.FileTransferSpec.FileFilterRegEx) == 0 {
		return false, fmt.Errorf("missing file filter regex, i need atlease one exp., to continue")
	}

	if a.Spec.LogFileDir == "" {
		a.Spec.LogFileDir = pwd + "/logs"
	}
	if b, _ := Exists(a.Spec.LogFileDir); !b {
		os.Mkdir(a.Spec.LogFileDir, os.ModePerm)
	}

	if a.Spec.DlqDir == "" {
		a.Spec.DlqDir = pwd + "/dlq"
	}
	if b, _ := Exists(a.Spec.DlqDir); !b {
		os.Mkdir(a.Spec.DlqDir, os.ModePerm)
	}

	if a.Spec.Archive {
		if a.Spec.ArchiveDir == "" {
			a.Spec.ArchiveDir = pwd + "/arc"
		}
		if b, _ := Exists(a.Spec.ArchiveDir); !b {
			os.Mkdir(a.Spec.ArchiveDir, os.ModePerm)
		}
	}

	return true, nil

}

func getApplicationConfiguration(ymlAppConfigPath string) (*ApplicationConfig, error) {
	if ymlAppConfigPath == "" {
		return nil, fmt.Errorf("application configuration path is empty")
	}

	if b, err := Exists(ymlAppConfigPath); !b {
		return nil, err
	}

	data, e := os.ReadFile(ymlAppConfigPath)
	if e != nil {
		return nil, e
	}

	a := ApplicationConfig{}

	// Unmarshal the YAML data into the AppConfig struct
	e = yaml.Unmarshal(data, &a)
	if e != nil {
		return nil, e
	}

	if b, e := a.IsValid(); !b {
		return nil, e
	}

	return &a, nil

}
