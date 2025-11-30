package logger

import (
	"encoding/json"
	"log"
	"os"
)

var LogFile *os.File

func InitLogger() {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		log.Fatalf("Error creando carpeta logs: %v", err)
	}

	LogFile, err = os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error abriendo logs/app.log: %v", err)
	}

	log.SetOutput(LogFile)
	log.SetFlags(0)
}

func WriteJSON(entry map[string]interface{}) {
	b, err := json.Marshal(entry)
	if err != nil {
		log.Println(`{"level":"error","msg":"error serializando log json"}`)
		return
	}

	log.Println(string(b))
}
