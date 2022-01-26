package main

import (
	"encoding/xml"
	"fmt"
	"github.com/HRKings/cfm-go-api/data"
	"github.com/HRKings/cfm-go-api/utils"
	"net/http"
	"strings"
)

const baseurl = "https://ws.cfm.org.br:8080/WebServiceConsultaMedicos/ServicoConsultaMedicos"
const cfmRequest = `
	<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://schemas.xmlsoap.org/soap/http" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:tns="http://servico.cfm.org.br/" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	 <SOAP-ENV:Body>
	   <mns1:Consultar xmlns:mns1="http://servico.cfm.org.br/">
	     <crm>%d</crm>
	     <uf>%s</uf>
	     <chave>%s</chave>
	   </mns1:Consultar>
	 </SOAP-ENV:Body>
	</SOAP-ENV:Envelope>`

func GetCRMInfo(crm int, state string, accessCode string) (data.DoctorData, error) {
	body := fmt.Sprintf(cfmRequest, crm, state, accessCode)

	resp, err := http.Post(baseurl, "text/xml", strings.NewReader(body))
	if err != nil {
		return data.DoctorData{}, err
	}
	defer resp.Body.Close()

	response := data.CfmResponse{}

	reader := utils.NewValidUTF8Reader(resp.Body)
	decoder := xml.NewDecoder(reader)

	err = decoder.Decode(&response)
	if err != nil {
		return data.DoctorData{}, err
	}

	return response.Body.ConsultarResponse.DoctorInfo, nil
}