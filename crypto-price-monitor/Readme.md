
## Crypto-price-monitor

Crypto Price Monitor is a golang cronjob that: 

- Gets the price of the `BTC/USD` pair from bitstamp.net
- Compares the `ASK` price with the `purchasePrice` defined on config file
- If the `ASK` price is lower or equals than the `purchasePrice` its sends you an email

### Configuration and dev environment

Prerequisites: 
- a gmail account to send mails
- golang installed

1. `gedit crypto-price-monitor.env` and set:

`TO`: the receiver of the email
`FROM`: your gmail account
`PASSWORD`: the senders password (please use a fake gmail account, not your personal one)
`ALERT_PRICE`: A float value that represents the amount that you want to receive a notification

2. `go run main.go`


### Build and Linux instalation

1. `go build`
2. `sudo cp crypto-price-monitor /usr/local/bin`
3. modify `cypto-price-monitir.env` with the price and the gmail credentials
4. `sudo cp crypto-price-monitor.env /`
5. create a new system service 
    a. `sudo gedit /etc/systemd/system/crpto-price-monitor.service`
    b. Fill it with: 
    
```

    [Unit]
    Description=Tests systemd to daemonize a Go binary
    Wants=network.target
    After=network.target

    [Service]
    Type=simple
    DynamicUser=yes
    ExecStart=/usr/local/bin/crypto-price-monitor 
    Restart=on-failure
    RestartSec=20

    [Install]
    WantedBy=multi-user.target

```

6. run `sudo-systemctl daemon-reload`
7. check status (should be inactive): `systemctl status crypto-price-monitor`
8. start service: `sudo systemctl start crypto-price-monitor`


See logs at: `journalctl -u crypto-price-monitor.service -e `



