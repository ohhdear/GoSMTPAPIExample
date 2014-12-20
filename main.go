package main

import (
	"github.com/sendgrid/sendgrid-go"
	"fmt"
)

func main() {
	sgclient := sendgrid.NewSendGridClient("USERNAME", "PASSWORD")
	msg := sendgrid.NewMail()
	msg.AddTo("xxxx@sendgrid.com")
	msg.AddToName("John Doe")
	msg.SetSubject("SendGrid Go Examples")
	msg.SetText("Text!")
	msg.SetHTML("<h1>HTML!</h1>")
	msg.SetFrom("sales@yourfromdomain.com")

	//
	// SMTPAPI Examples below
	// Every App has enabled 0 : 1 flag but DKIM (cannot turn off)
	// ---------------------------
	tos := []string{"test1@sink.sendgrid.net", "test2@sink.sendgrid.net", "test3@sink.sendgrid.net"}
	msg.SMTPAPIHeader.AddTos(tos)
	
	// Add BCCs to SMTPAPI Header
	msg.AddFilter("bcc", "email", "AllBCCsGoHere@sink.sendgrid.net")

	// Bypass List Management for Whitelisting known good addresses
	msg.AddFilter("bypass_list_management", "enabled", "1")

	// Manages Click Tracking app
	msg.AddFilter("clicktrack", "enabled", "1")

	// Change DKIM domain to "From" domain
	msg.AddFilter("dkim", "use_from", "1")
	msg.AddFilter("dkim", "domain", "yourfromdomain.com")

	// Add Footer to all emails sent
	msg.AddFilter("footer", "text/plain", "plain footer")
	msg.AddFilter("footer", "text/html", "<h3>Footer</h3>")

	// Forward spam rulings to email address
	msg.AddFilter("forwardspam", "email", "SpamAddress@yourfromdomain.com")

	// Google Analytics app. Takes campaign source, medium, term, content, name
	msg.AddFilter("ganalytics", "utm_source", "sendgrid.com")
	msg.AddFilter("ganalytics", "utm_medium", "email")
	msg.AddFilter("ganalytics", "utm_term", "adwords")
	msg.AddFilter("ganalytics", "utm_content", "banner ad")
	msg.AddFilter("ganalytics", "utm_campaign", "Go library promo")

	// Manages Open Tracking 
	msg.AddFilter("opentrack", "enabled", "1")

	// Spam Check posts spam verdicts on SpamAssassin to a URL endpoint 
	// maxscore : "-10" to "10"
	msg.AddFilter("spamcheck", "enabled", "1")
	msg.AddFilter("spamcheck", "maxscore", "5")
	msg.AddFilter("spamcheck", "url", "https://urlendpoint.sendgrideventkit.herokuapp.com/")

	// Subscription Tracking App Management
	// text/plain : "If you would like to unsubscribe and stop receiving these emails click here: <% %>."
	// text/html : "<p>If you would like to unsubscribe and stop receiving these emails <% click here %>.</p>"
	msg.AddFilter("subscriptiontrack", "enabled", "1")
	msg.AddFilter("subscriptiontrack", "text/plain", "If you would like to unsubscribe and stop receiving these emails click here: <% %>.")
	msg.AddFilter("subscriptiontrack", "text/html", "<p>If you would like to unsubscribe and stop receiving these emails <% click here %>.</p>")

	// Template engine for transactional templates
	// uses SetHTML() for <%body%> tag substitution
	// template_id : "1234-5678" (Must me active template)
	msg.AddFilter("templates", "template_id", "adbc-12345-8675-309")

	
	// Send mail and check success
	if r := sgclient.Send(msg); r == nil {
		fmt.Println("Email sent!")
	} else {
		fmt.Println(r)
	}
}