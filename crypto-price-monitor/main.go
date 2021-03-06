package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

// Config is the program configuration
type Config struct {
	To         string  `mapstructure:"TO"`
	AlertPrice float64 `mapstructure:"ALERT_PRICE"`
	From       string  `mapstructure:"FROM"`
	Password   string  `mapstructure:"PASSWORD"`
}

// TickerGetResponse represents a bitstamp.net/api/ticker GET response
type TickerGetResponse struct {
	Volume    string
	Last      string
	Timestamp string
	Bid       string
	Vwap      string
	High      string
	Low       string
	Ask       string
	Open      float32
}

// Gmail smtpServer data
type smtpServer struct {
	host string
	port string
}

// Address URI for smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("crypto-price-monitor")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// GetTicker makes a HTTP GET request to bitstamp to get BTC/USD transfer pair value
func GetTicker() (string, error) {
	resp, err := http.Get("https://www.bitstamp.net/api/ticker/")
	if err != nil {
		return "error", err
	}
	fmt.Println("123123123  asda")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var tickerGetResponse TickerGetResponse
	json.Unmarshal(body, &tickerGetResponse)

	return tickerGetResponse.Ask, nil
}

// sendMail sends a gmail to a destinatary
func sendMail(data string, to string, from string, password string) error {
	receivers := []string{
		to,
	}
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	message := []byte(
		"Subject: BTC/USD transfer pair value alert \r\n" +
			"\r\n" +
			data +
			"\r\n")
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	err := smtp.SendMail(smtpServer.Address(), auth, from, receivers, message)
	if err != nil {
		return err
	}
	fmt.Println("Mail sent")
	return nil
}

// Main function. Creates a new cron job every 1 hour that:
// - gets the value of btc/usd pair
// - compare it with the purchase price
// - if its lower, send an email to the recipient cli argument
func main() {
	config, err := LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config:", err)
		os.Exit(1)
	}
	c := cron.New()
	c.AddFunc("1 * * * *", func() {
		fmt.Println("Start running with config: ", config)
		tickerValue, _ := GetTicker()
		comparingValue, _ := strconv.ParseFloat(tickerValue, 32)
		if comparingValue <= config.AlertPrice {
			mailText := "Current Bitcoin value is " + tickerValue
			sendMail(mailText, config.To, config.From, config.Password)
		}

	})
	c.Run()
}
