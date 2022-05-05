package socket

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

const MAX_PERCENT = 100
const DONE_PERCENT = 101

type PdfGenerator struct {
	Uuid    uuid.UUID `json:"uuid"`
	Percent int       `json:"percent"`
}

type PdfGeneratorDto struct {
	PdfGenerator
	Type string `json:"type"`
}

type PdfGeneratorSendDto struct {
	msgType string
	data    []PdfGenerator
}

var clients = make(map[string][]PdfGenerator, 10)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	}}

func sendPercent(c *websocket.Conn, sendData []PdfGenerator) error {
	jsonSendData, err := json.Marshal(sendData)
	if err != nil {
		log.Println("cant marshal to Json:", err)
		return err
	}

	err = c.WriteMessage(1, jsonSendData)
	if err != nil {
		log.Println("write:", err)
		return err
	}
	return nil
}

func sendTimer(c *websocket.Conn, header string) {
	for now := range time.Tick(time.Second) {
		wg := &sync.WaitGroup{}
		for val := range clients[header] {
			wg.Add(1)

			go func(wg *sync.WaitGroup, val int) {
				var percent int

				if clients[header][val].Percent == MAX_PERCENT {
					percent = DONE_PERCENT
				} else {
					percent = clients[header][val].Percent + 1
				}

				clients[header][val] = PdfGenerator{
					Uuid:    clients[header][val].Uuid,
					Percent: percent,
				}

				wg.Done()
			}(wg, val)
		}

		wg.Wait()
		log.Println(now)
		err := sendPercent(c, clients[header])

		if err != nil {
			log.Print("upgrade:", err)
			return
		}
	}
}
func Echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	var header = r.Header.Get("Authorization")
	defer c.Close()

	if clients[header] == nil {
		clients[header] = make([]PdfGenerator, 0, 10)
	}

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		recievedMsg := &PdfGeneratorDto{}
		json.Unmarshal(message, &recievedMsg)

		if recievedMsg.Type == "startConn" {
			go sendTimer(c, header)
		}

		if recievedMsg.Type == "createPdf" {
			recievedPdf := PdfGenerator{
				Uuid:    recievedMsg.Uuid,
				Percent: recievedMsg.Percent,
			}

			clients[header] = append(clients[header], recievedPdf)
		}

		err = c.WriteMessage(mt, []byte("hi from server"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}
