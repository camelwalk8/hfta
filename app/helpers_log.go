package main

import "log"

func Infol(message string) {
	if IL == nil {
		if b, _ := SetupLogging(LogDirectory); !b {
			log.Println(message)
			return
		}

	}
	IL.Println(message)
}

func Debugl(message string) {
	if DL == nil {
		if b, _ := SetupLogging(LogDirectory); !b {
			log.Println(message)
			return
		}
	}
	DL.Println(message)
}

func Errorl(message string) {
	if EL == nil {
		if b, _ := SetupLogging(LogDirectory); !b {
			log.Println(message)
			return
		}
	}
	EL.Println(message)
}

func Fatall(message string) {
	if FL == nil {
		if b, _ := SetupLogging(LogDirectory); !b {
			log.Println(message)
			return
		}
	}
	FL.Println(message)
}
