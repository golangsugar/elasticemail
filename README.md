# ElasticEmail
GOLang Helper for [ElasticEmail](https://elasticemail.com/) integration
<br>
The default environment variable with API key is **ELASTICEMAIL_APIKEY**, but you can replace with yours
<hr>

## Example

Start with
```Bash
    go get -u github.com/golangsugar/elasticemail
```

Then add
```Go
import "github.com/golangsugar/elasticemail"
```
to your source code
<br>

```Go
var msg elasticemail.Message

msg.SetSender("John Smith", "smith@email.com")
msg.AddRecipient("Gene", "Simons")
msg.SetSubject("Here is that message you're waiting for")

html := `<main><p style="color:#555;font-family:Arial,Helvetica,sans-serif;font-size:1em;">GolangSugar/ElasticEmail Example</p></main>`

msg.SetHTML(html)

return msg.Send()
```
