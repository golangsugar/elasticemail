package elasticemail

import "os"

type Message struct {
	Template      string
	Substitutions map[string]string
	From          Person
	To            []Person
	CC            []Person
	BCC           []Person
	Subject       string
	ReplyTo       Person
	HTML          string
	Text          string
}

func (m Message) SetTemplate(t string) Message {
	m.Template = t

	return m
}

func (m Message) SetSender(name, address string) Message {
	m.From = Person{
		Name:    name,
		Address: address,
	}

	return m
}

func (m Message) SetReplyTo(name, address string) Message {
	m.ReplyTo = Person{
		Name:    name,
		Address: address,
	}

	return m
}

func (m Message) AddRecipient(name, address string) Message {
	m.To = append(m.To, Person{
		Name:    name,
		Address: address,
	})

	return m
}

func (m Message) SetRecipient(name, address string) Message {
	m.To = []Person{
		{
			Name:    name,
			Address: address,
		},
	}

	return m
}

func (m Message) AddCC(name, address string) Message {
	m.CC = append(m.CC, Person{
		Name:    name,
		Address: address,
	})

	return m
}

func (m Message) SetCC(name, address string) Message {
	m.CC = []Person{
		{
			Name:    name,
			Address: address,
		},
	}

	return m
}

func (m Message) AddBCC(name, address string) Message {
	m.BCC = append(m.BCC, Person{
		Name:    name,
		Address: address,
	})

	return m
}

func (m Message) SetBCC(name, address string) Message {
	m.BCC = []Person{
		{
			Name:    name,
			Address: address,
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

func New() Message {
	return Message{}
}

send?

func (m Message) Send() error {
	payload := map[string]interface{}{
		"apikey":  os.Getenv("ELASTICEMAIL_APIKEY"),
		"isTransactional":false,
		"subject": m.Subject,
		"sender":m.From.Address,
		"senderName":m.From.Name,
		"replyTo":m.ReplyTo.Address,
		"replyToName":m.ReplyTo.Name,
		"msgTo":,
		"msgCC":,
		"msgBcc":,
	}

	if m.HTML!="" {
		payload["bodyHtml"] = m.HTML
		payload["charsetBodyHtml"]=`utf-8`
	}

	if m.Text!="" {
		payload["bodyText"] = m.HTML
		payload["charsetBodyText"]=`utf-8`
	}

	if m.Template != "" {
		payload["template"] = m.Template

		for k, v := range m.Substitutions {
			key := `merge_` + k
			payload[key] = v
		}
	}

	return nil
}
