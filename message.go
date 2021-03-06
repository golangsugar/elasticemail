package elasticemail

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

func (m *Message) SetTemplate(t string) {
	m.Template = t
}

func (m *Message) SetSender(name, address string) {
	m.From = person{
		name:    name,
		address: address,
	}
}

func (m *Message) SetReplyTo(name, address string) {
	m.ReplyTo = person{
		name:    name,
		address: address,
	}
}

func (m *Message) AddRecipient(name, address string) {
	m.To = append(m.To, person{
		name:    name,
		address: address,
	})
}

func (m *Message) SetRecipient(name, address string) {
	m.To = []person{
		{
			name:    name,
			address: address,
		},
	}
}

func (m *Message) AddCC(name, address string) {
	m.CC = append(m.CC, person{
		name:    name,
		address: address,
	})
}

func (m *Message) SetCC(name, address string) {
	m.CC = []person{
		{
			name:    name,
			address: address,
		},
	}
}

func (m *Message) AddBCC(name, address string) {
	m.BCC = append(m.BCC, person{
		name:    name,
		address: address,
	})
}

func (m *Message) SetBCC(name, address string) {
	m.BCC = []person{
		{
			name:    name,
			address: address,
		},
	}
}

func (m *Message) SetSubject(s string) {
	m.Subject = s
}

func (m *Message) SetHTML(s string) {
	m.HTML = s
}

func (m *Message) SetText(s string) {
	m.Text = s
}

func (m *Message) AddVariable(key, value string) {
	if m.Substitutions == nil {
		m.Substitutions = make(map[string]string)
	}

	m.Substitutions[key] = value
}

func peopleAsString(people []person) string {
	var a []string

	for _, p := range people {
		if p.name != "" {
			a = append(a, p.name+" <"+p.address+">")
		} else {
			a = append(a, p.address)
		}
	}

	return strings.Join(a, ";")
}

func (m Message) asMap() map[string]string {
	payload := map[string]string{
		"apikey":          os.Getenv(elasticEmailAPIKeyEmailEnvVarName),
		"isTransactional": "true",
		"subject":         m.Subject,
		"from":            m.From.address,
		"fromName":        m.From.name,
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
	msgParams := m.asMap()

	form := url.Values{}

	for k, v := range msgParams {
		form.Set(k, v)
	}

	endpointURL := fmt.Sprintf("%s/send", apiEndpoint)

	params := form.Encode()

	req, err2 := http.NewRequest("POST", endpointURL, strings.NewReader(params))

	if err2 != nil {
		return err2
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Content-Length", strconv.Itoa(len(params)))

	hc := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, errh := hc.Do(req)

	if errh != nil {
		log.Println(errh)
		return errh
	}

	if resp != nil {
		if resp.Body != nil {
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				log.Println(string(body))
			}
		}

		if resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Printf("%d %s\n", resp.StatusCode, resp.Status)
	}

	return fmt.Errorf("error while trying to send email")
}
