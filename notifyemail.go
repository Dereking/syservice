package syservice

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
	"time"
)

func SendMail(address []string, subject string, body string, files []string) (err error) {

	mail := &MailSmtp{
		user:     yamlConfig.Smtp.User,
		password: yamlConfig.Smtp.Password,
		host:     yamlConfig.Smtp.Host, port: "25",
	}
	message := Message{
		from:        yamlConfig.Smtp.User,
		to:          address,
		cc:          []string{},
		bcc:         []string{},
		subject:     subject, //邮件标题
		body:        body,    //正文内容
		contentType: "text/plain;charset=utf-8",
		attachment: Attachment{
			name:        files, //可以放入多张图片
			contentType: "image/png",
			withFile:    true,
		},
	}

	return mail.Send(message)

	// // 通常身份应该是空字符串，填充用户名.
	// auth := smtp.PlainAuth("", yamlConfig.Smtp.User, yamlConfig.Smtp.Password, yamlConfig.Smtp.Host)

	// contentType := "Content-Type: text/html; charset=UTF-8"
	// for _, v := range address {
	// 	s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",
	// 		v, yamlConfig.Smtp.Nickname, yamlConfig.Smtp.User, subject, contentType, body)
	// 	msg := []byte(s)
	// 	addr := fmt.Sprintf("%s:%s", yamlConfig.Smtp.Host, yamlConfig.Smtp.Port)
	// 	err = smtp.MailSmtp(addr, auth, yamlConfig.Smtp.User, []string{v}, msg)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// return
}

type MailSmtp struct {
	user     string
	password string
	host     string
	port     string
	auth     smtp.Auth
}

type Attachment struct {
	name        []string
	contentType string
	withFile    bool
}

type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  Attachment
}

func (mail *MailSmtp) Auth() {
	mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
}

func (mail *MailSmtp) Send(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/related;boundary=" + boundary
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	var imgsrc string
	if message.attachment.withFile {
		//多图片发送
		for _, graphname := range message.attachment.name {
			attachment := "\r\n--" + boundary + "\r\n"
			attachment += "Content-Transfer-Encoding:base64\r\n"
			attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + graphname + "\"\r\n"
			attachment += "Content-ID: <" + graphname + "> \r\n\r\n"
			buffer.WriteString(attachment)

			//拼接成html
			imgsrc += "<p><img src=\"cid:" + graphname + "\" height=200 width=300></p><br>\r\n\t\t\t"

			defer func() {
				if err := recover(); err != nil {
					fmt.Printf(err.(string))
				}
			}()
			mail.writeFile(buffer, graphname)
		}
	}

	//需要在正文中显示的html格式
	var template = `
    <html>
        <body>
            <p>text:%s</p><br>
            %s          
        </body>
    </html>
    `
	var content = fmt.Sprintf(template, message.body, imgsrc)
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type: text/html; charset=UTF-8 \r\n"
	body += content
	buffer.WriteString(body)

	buffer.WriteString("\r\n--" + boundary + "--")
	//	fmt.Println(buffer.String())
	smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, message.to, buffer.Bytes())
	return nil
}

func (mail *MailSmtp) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

func (mail *MailSmtp) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
