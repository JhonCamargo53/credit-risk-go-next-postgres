package logger

import (
	"encoding/json"
	"log"
	"os"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
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
