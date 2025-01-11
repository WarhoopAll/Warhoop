package soap

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
	"warhoop/app/config"
	"warhoop/app/log"
	"warhoop/app/utils"
)

type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Xmlns   string   `xml:"xmlns:soap,attr"`
	Body    Body     `xml:"soap:Body"`
}

type Body struct {
	ExecuteCommand ExecuteCommand `xml:"executeCommand"`
}

type ExecuteCommand struct {
	Xmlns   string `xml:"xmlns,attr"`
	Command string `xml:"command"`
}

var cfg = config.Get()

func (svc *SoapService) ExecuteCommand(command string) (string, error) {
	if !cfg.Soap.Enable {
		svc.logger.Warn("SOAP is disabled in the configuration")
		return "", utils.ErrDisable
	}

	soapEnvelope := Envelope{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: Body{
			ExecuteCommand: ExecuteCommand{
				Xmlns:   "urn:TC",
				Command: command,
			},
		},
	}

	var soapRequestBuffer bytes.Buffer
	encoder := xml.NewEncoder(&soapRequestBuffer)
	if err := encoder.Encode(soapEnvelope); err != nil {
		svc.logger.Error("Failed to encode SOAP request",
			log.String("command", command),
			log.String("error", err.Error()),
		)
		return "", err
	}

	soapRequest := soapRequestBuffer.String()

	svc.logger.Debug("SOAP request generated",
		log.String("command", command),
		log.String("soapRequest", soapRequest),
	)

	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(cfg.Soap.Login+":"+cfg.Soap.Password))

	req, err := http.NewRequest("POST", cfg.Soap.Host, bytes.NewBufferString(soapRequest))
	if err != nil {
		svc.logger.Error("Failed to create HTTP request",
			log.String("command", command),
			log.String("error", err.Error()),
		)
		return "", err
	}

	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		svc.logger.Error("SOAP request failed",
			log.String("command", command),
			log.String("error", err.Error()),
		)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		svc.logger.Error("SOAP request returned non-OK status",
			log.String("command", command),
			log.Int("status", resp.StatusCode),
			log.String("response", string(body)),
		)
		return "", fmt.Errorf("non-OK status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		svc.logger.Error("Failed to read SOAP response",
			log.String("command", command),
			log.String("error", err.Error()),
		)
		return "", err
	}

	svc.logger.Debug("SOAP command executed successfully",
		log.String("command", command),
	)

	return string(body), nil
}
