package elasticemail

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

func (m Message) charset() string {
	return `utf-8`
}

send?apikey = 94DAF66E-4DF6-4E8E-AF96-D094A8D21DF3&

Template string
Substitutions map[string]string
Recipients []Recipient
Subject string
From =
&fromName
=&sender = &
senderName =

&replyTo =
&replyToName =
&to =

&msgCC =
&msgBcc=
&lists =
bodyHtml = &
bodyText =

&charsetBodyHtml =
&charsetBodyText =
&encodingType =
&template=
&headers_firstname = firstname: myValueHere
isTransactional =false

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
