# 🖧 networkspeedcheckbot

[![Go Reference](https://pkg.go.dev/badge/github.com/hakimkal/networkspeedcheckbot.svg)](https://pkg.go.dev/github.com/hakimkal/networkspeedcheckbot)
[![Telegram](https://img.shields.io/badge/Telegram-Bot-blue?logo=telegram)](https://t.me/hakeemhal)

A **Telegram Bot** written in **Go** for engineers to **check the internet speed of their servers** quickly, without SSH or logging into dashboards.

>  Note: This bot tests the **server’s network speed**, not the end user’s device.

---

##  Features

- Ping test
- Download speed test
- Upload speed test
- (Future) Interactive Telegram inline keyboard

---

##  Dependencies

- [`speedtest-go`](https://github.com/showwin/speedtest-go) – for measuring ping, download, and upload speeds
- [`go-telegram-bot-api`](https://github.com/go-telegram-bot-api/telegram-bot-api) – for Telegram interactions

---

##  Setup & Running

1. **Clone the repo**

```bash
git clone https://github.com/hakimkal/networkspeedcheckbot.git
cd networkspeedcheckbot
