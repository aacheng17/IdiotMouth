package idiotmouth

import (
	"fmt"
	"log"
	"strings"

	"example.com/hello/core"
)

// declaring a struct
type IdiotmouthHub struct {

	// declaring struct variable
	core.Hub

	start rune

	end rune

	usedWords []string

	letters map[string]int

	wordsLeft int
}

func (h *IdiotmouthHub) getAssertedClients() map[*IdiotmouthClient]bool {
	ret := make(map[*IdiotmouthClient]bool)
	for k, v := range h.Clients {
		ret[k.(*IdiotmouthClient)] = v
	}
	return ret
}

// MESSAGE TYPES (SERVER TO CLIENT)
// 0: regular chat messages
// 1: scores
// 2: prompt
// 3: restart (data is inconsequential, probably empty string)

func (h *IdiotmouthHub) HandleHubMessage(m *core.Message) {
	c := (m.Client).(*IdiotmouthClient)
	switch m.MessageType {
	case byte('0'):
		for client := range h.Clients {
			h.SendData(client, byte('0'), []byte(c.Name+": "+string(m.Data)))
		}
		word := strings.TrimSpace(strings.ToLower(string(m.Data)))
		if len(word) >= 3 && word[0] == byte(h.start) && word[len(word)-1] == byte(h.end) {
			switch h.validWord(word) {
			case 0:
				worth := h.getWorth()
				bonus := len(word) - 2
				finalWorth := worth * bonus
				c.score += finalWorth
				err := h.gotIt(word)
				if err == 1 {
					for client := range h.Clients {
						h.SendData(client, byte('0'), []byte("You have passed or gotten all possible words. Type restart to restart the game."))
					}
					break
				}
				for client := range h.Clients {
					h.SendData(client, byte('0'), []byte(c.Name+" earned "+fmt.Sprint(worth)+"x"+fmt.Sprint(bonus)+"="+fmt.Sprint(finalWorth)+" points"))
					h.SendData(client, byte('0'), []byte("."))
					h.SendData(client, byte('2'), []byte(h.getPrompt()))
					h.SendData(client, byte('0'), []byte("New letters: "+h.getPrompt()))
					h.SendData(client, byte('1'), []byte(h.getScores()))
				}
			case 2:
				for client := range h.Clients {
					h.SendData(client, byte('0'), []byte("This word has already been used this game."))
				}
			}
		} else if string(m.Data) == "/pass" {
			c.pass = true
			if h.getMajorityPass() {
				err := h.pass()
				if err == 1 {
					for client := range h.Clients {
						h.SendData(client, byte('0'), []byte("You have passed or gotten all possible words. Type restart to restart the game."))
					}
					break
				}
				for client := range h.Clients {
					h.SendData(client, byte('0'), []byte("Letters passed, new letters generated"))
					h.SendData(client, byte('0'), []byte("."))
					h.SendData(client, byte('0'), []byte("New letters: "+h.getPrompt()))
					h.SendData(client, byte('2'), []byte(h.getPrompt()))
				}
			}
		} else if string(m.Data) == "/restart" {
			h.reset()
			for client := range h.Clients {
				h.SendData(client, byte('3'), []byte(""))
				h.SendData(client, byte('0'), []byte(c.Name+" restarted the game"))
				h.SendData(client, byte('1'), []byte(h.getScores()))
				h.SendData(client, byte('2'), []byte(h.getPrompt()))
				h.SendData(client, byte('0'), []byte("."))
				h.SendData(client, byte('0'), []byte("New letters: "+h.getPrompt()))
			}
		}
	case byte('1'):
		name := string(m.Data)
		if c.Name == "" {
			c.Name = name
		}
		for client := range h.Clients {
			h.SendData(client, byte('0'), []byte(name+" joined"))
		}
		for client := range h.Clients {
			h.SendData(client, byte('1'), []byte(h.getScores()))
		}
		h.SendData(c, byte('2'), []byte(h.getPrompt()))
		h.SendData(c, byte('0'), []byte("New letters: "+h.getPrompt()))
	}
}

func (h *IdiotmouthHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				h.RemoveClient(client, "Removed client that disconnected.")
			}
		case message := <-h.Messages:
			log.Println("Received message\n\tType: " + fmt.Sprint(message.MessageType) + "\n\tData: " + string(message.Data))
			h.HandleHubMessage(message)
		}
	}
}

func NewIdiotmouthHub() core.Hublike {
	h := &IdiotmouthHub{
		Hub:     *core.NewHub(),
		letters: make(map[string]int),
	}
	h.reset()
	return h
}
