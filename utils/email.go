package utils

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"gee/global"
	"net/smtp"
	"strings"
	"github.com/jordan-wright/email"
)

func MD5VByte(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5VString (str string) string{
	s :=[]byte(str)
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

func Email(subject,body string) error {
	to := strings.Split(global.Config.Email.To, ",")
	return send(to,subject,body)
}

func ErrorToEmail(subject,body string) error {
	to := strings.Split(global.Config.Email.To,",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to,subject,body)
}

func send(to []string,subject,body string) (err error) {
	from := global.Config.From
	nickName := global.Config.Nickname
	secret := global.Config.Secret
	host := global.Config.Email.Host
	port := global.Config.Email.Port
	isSSL := global.Config.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickName != ""{
		e.From = fmt.Sprintf("%s <%s>", nickName, from)
	}else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	hostPort := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostPort,auth,&tls.Config{ServerName: host})
	}else {
		err = e.Send(hostPort,auth)
	}
	return
}