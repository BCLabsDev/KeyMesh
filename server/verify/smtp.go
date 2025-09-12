package verify

import (
	"crypto/tls"
	"fmt"
	"io"
	"keymesh/utils/config"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type email struct{}

func sendEmail(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", config.SMTP_Host, config.SMTP_Prot)

	// 配置 TLS
	tlsConfig := &tls.Config{ServerName: config.SMTP_Host}

	// 建立 TLS 连接
	c, err := smtp.DialTLS(addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("DialTLS error: %w", err)
	}
	defer c.Close()

	// 认证
	auth := sasl.NewPlainClient("", config.SMTP_Email, config.SMTP_Pass)
	if err := c.Auth(auth); err != nil {
		return fmt.Errorf("Auth failed: %w", err)
	}

	// 设置发件人
	if err := c.Mail(config.SMTP_Email, nil); err != nil {
		return fmt.Errorf("MAIL FROM error: %w", err)
	}

	// 设置收件人
	if err := c.Rcpt(to, nil); err != nil {
		return fmt.Errorf("RCPT TO error: %w", err)
	}

	// 写入邮件内容
	wc, err := c.Data()
	if err != nil {
		return fmt.Errorf("DATA command error: %w", err)
	}

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		config.SMTP_Email, to, subject, body)

	if _, err = io.Copy(wc, strings.NewReader(msg)); err != nil {
		return fmt.Errorf("write message error: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("closing writer error: %w", err)
	}

	// 结束
	if err := c.Quit(); err != nil {
		return fmt.Errorf("QUIT error: %w", err)
	}

	return nil
}
func (email) SendCode(code, to string) error {
	subject := "KeyMesh验证码"
	body := "您刚刚发送了一个验证码。请验证您的邮箱地址，让我们知道您是此账户的合法所有人。您的验证码是：" + code
	return sendEmail(to, subject, body)
}
