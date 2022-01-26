package data

import "encoding/xml"

type DoctorData struct {
	Text             string   `xml:",chardata"`
	CRM              string   `xml:"crm"`
	UpdatedAt        string   `xml:"dataAtualizacao"`
	Specialties      []string `xml:"especialidade"`
	Name             string   `xml:"nome"`
	Status           string   `xml:"situacao"`
	RegistrationType string   `xml:"tipoInscricao"`
	UF               string   `xml:"uf"`
}

type CfmResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text              string `xml:",chardata"`
		ConsultarResponse struct {
			Text       string `xml:",chardata"`
			Ns2        string `xml:"ns2,attr"`
			DoctorInfo DoctorData `xml:"dadosMedico"`
		} `xml:"ConsultarResponse"`
	} `xml:"Body"`
}
