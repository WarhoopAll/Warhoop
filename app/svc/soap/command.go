package soap

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
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
		svc.logger.Warn("service.SoapService.ExecuteCommand",
			log.Bool("cfg.Soap.enable", cfg.Soap.Enable),
		)
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
		svc.logger.Error("service.SoapService.ExecuteCommand",
			log.String("command", command),
			log.String("err", err.Error()),
		)
		return "", err
	}

	soapRequest := soapRequestBuffer.String()

	svc.logger.Debug("service.SoapService.ExecuteCommand",
		log.String("command", command),
		log.String("soapRequest", soapRequest),
	)

	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(cfg.Soap.Login+":"+cfg.Soap.Password))

	req, err := http.NewRequest("POST", cfg.Soap.Host, bytes.NewBufferString(soapRequest))
	if err != nil {
		svc.logger.Error("service.SoapService.ExecuteCommand",
			log.String("err", err.Error()),
			log.String("command", command),
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
		svc.logger.Error("service.SoapService.ExecuteCommand",
			log.String("err", err.Error()),
			log.String("command", command),
		)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return "", readErr
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		svc.logger.Error("service.SoapService.ExecuteCommand",
			log.String("err", err.Error()),
			log.String("command", command),
		)
		return "", err
	}

	svc.logger.Debug("service.SoapService.ExecuteCommand",
		log.String("command", command),
	)

	return string(body), nil
}
