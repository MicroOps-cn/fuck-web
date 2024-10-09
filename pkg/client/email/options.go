package email

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/gogo/protobuf/proto"

	"github.com/MicroOps-cn/fuck-web/api"
)

//go:embed template
var templFs embed.FS

var innerTmpl []Template

type Template struct {
	*OriginalTemplate
	tmpl *template.Template
}

func init() {
	innerTmpl = append(innerTmpl, Template{
		OriginalTemplate: &OriginalTemplate{
			Subject:      "Reset Password",
			Topic:        "User:ResetPassword",
			TemplateFile: "template/reset_password.html",
		},
		tmpl: template.Must(template.ParseFS(templFs, "template/reset_password.html")),
	}, Template{
		OriginalTemplate: &OriginalTemplate{
			Subject:      "Activate Account",
			Topic:        "User:ActivateAccount",
			TemplateFile: "template/activate_account.html",
		},
		tmpl: template.Must(template.ParseFS(templFs, "template/activate_account.html")),
	}, Template{
		OriginalTemplate: &OriginalTemplate{
			Subject:      "One-time Password",
			Topic:        "User:SendLoginCaptcha",
			TemplateFile: "template/send_login_code.html",
		},
		tmpl: template.Must(template.ParseFS(templFs, "template/send_login_code.html")),
	})
}

var _ api.CustomType = &Template{}

func (t Template) Marshal() ([]byte, error) {
	return proto.Marshal(t.OriginalTemplate)
}

func (t *Template) Unmarshal(data []byte) (err error) {
	if t.OriginalTemplate == nil {
		t.OriginalTemplate = &OriginalTemplate{}
	}
	if err = proto.Unmarshal(data, t.OriginalTemplate); err != nil {
		return err
	}
	t.tmpl, err = template.ParseFiles(t.TemplateFile)
	return err
}

func (t Template) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.OriginalTemplate)
}

func (t *Template) UnmarshalJSON(data []byte) (err error) {
	if t.OriginalTemplate == nil {
		t.OriginalTemplate = &OriginalTemplate{}
	}
	if err = json.Unmarshal(data, &t.OriginalTemplate); err != nil {
		return err
	}
	t.tmpl, err = template.ParseFiles(t.TemplateFile)
	if err != nil {
		return err
	}
	return
}

func (m *SmtpOptions) getTemplate(topic string, sets ...string) *Template {
	var tmpls []Template
	tmpls = append(tmpls, m.Template...)
	tmpls = append(tmpls, innerTmpl...)
	for _, set := range sets {
		for _, tmpl := range tmpls {
			if tmpl.Topic == topic && tmpl.Set == set {
				return &tmpl
			}
		}
	}
	for _, tmpl := range tmpls {
		if tmpl.Topic == topic && (tmpl.Set == "" || tmpl.Set == "__default__") {
			return &tmpl
		}
	}
	return nil
}

func (m *SmtpOptions) GetSubjectAndBody(data interface{}, topic string, sets ...string) (subject, body string, err error) {
	t := m.getTemplate(topic, sets...)
	buffer := new(bytes.Buffer)
	if t == nil || t.tmpl == nil {
		return "", "", fmt.Errorf("template is nil")
	}
	if err = t.tmpl.Execute(buffer, data); err != nil {
		return "", "", err
	}
	return t.GetSubject(), buffer.String(), nil
}
