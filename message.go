package elasticemail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type person struct {
	name,
	address string
}

type Message struct {
	Template      string
	Substitutions map[string]string
	From          person
	To            []person
	CC            []person
	BCC           []person
	Subject       string
	ReplyTo       person
	HTML          string
	Text          string
}

func (m Message) SetTemplate(t string) Message {
	m.Template = t

	return m
}

func (m Message) SetSender(name, address string) Message {
	m.From = person{
		name:    name,
		address: address,
	}

	return m
}

func (m Message) SetReplyTo(name, address string) Message {
	m.ReplyTo = person{
		name:    name,
		address: address,
	}

	return m
}

func (m Message) AddRecipient(name, address string) Message {
	m.To = append(m.To, person{
		name:    name,
		address: address,
	})

	return m
}

func (m Message) SetRecipient(name, address string) Message {
	m.To = []person{
		{
			name:    name,
			address: address,
		},
	}

	return m
}

func (m Message) AddCC(name, address string) Message {
	m.CC = append(m.CC, person{
		name:    name,
		address: address,
	})

	return m
}

func (m Message) SetCC(name, address string) Message {
	m.CC = []person{
		{
			name:    name,
			address: address,
		},
	}

	return m
}

func (m Message) AddBCC(name, address string) Message {
	m.BCC = append(m.BCC, person{
		name:    name,
		address: address,
	})

	return m
}

func (m Message) SetBCC(name, address string) Message {
	m.BCC = []person{
		{
			name:    name,
			address: address,
		},
	}

	return m
}

func (m Message) SetSubject(s string) Message {
	m.Subject = s

	return m
}

func (m Message) SetHTML(s string) Message {
	m.HTML = s

	return m
}

func (m Message) SetText(s string) Message {
	m.Text = s

	return m
}

func (m Message) AddVariable(key, value string) Message {
	if m.Substitutions == nil {
		m.Substitutions = make(map[string]string)
	}

	m.Substitutions[key] = value

	return m
}

func peopleAsString(people []person) string {
	var a []string

	for _, p := range people {
		a = append(a, p.name+"<"+p.address+">")
	}

	return strings.Join(a, ";")
}

func (m Message) asMap() map[string]interface{} {
	payload := map[string]interface{}{
		"apikey":          os.Getenv(elasticEmailAPIKeyEmailEnvVarName),
		"isTransactional": false,
		"subject":         m.Subject,
		"sender":          m.From.address,
		"senderName":      m.From.name,
		"replyTo":         m.ReplyTo.address,
		"replyToName":     m.ReplyTo.name,
		"msgTo":           peopleAsString(m.To),
		"msgCC":           peopleAsString(m.CC),
		"msgBcc":          peopleAsString(m.BCC),
	}

	if m.HTML != "" {
		payload["bodyHtml"] = m.HTML
		payload["charsetBodyHtml"] = `utf-8`
	}

	if m.Text != "" {
		payload["bodyText"] = m.HTML
		payload["charsetBodyText"] = `utf-8`
	}

	if m.Template != "" {
		payload["template"] = m.Template

		for k, v := range m.Substitutions {
			key := `merge_` + k
			payload[key] = v
		}
	}

	return payload
}

// Send Email using ElasticEmail API
func (m Message) Send() error {
	body := m.asMap()

	byteArray, err := json.Marshal(body)

	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/send?apikey=%s", apiEndpoint, os.Getenv(elasticEmailAPIKeyEmailEnvVarName))

	req, err2 := http.NewRequest("POST", url, bytes.NewBuffer(byteArray))

	if err2 != nil {
		return err2
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	hc := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, errh := hc.Do(req)

	if errh != nil {
		log.Println(errh)
		return errh
	}

	if resp != nil {
		if resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Printf("%d %s\n", resp.StatusCode, resp.Status)
	}

	return fmt.Errorf("error while trying to send email")
}
