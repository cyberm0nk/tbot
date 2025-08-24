# 🤖 tbot — Telegram Bot CLI built with Go & Cobra

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/dl/)  
A lightweight **Telegram bot** written in Go using [Cobra CLI](https://github.com/spf13/cobra) and [Telebot v3](https://github.com/tucnak/telebot).  
This project is designed as a **learning playground** for building command-line applications and Telegram bots in Go.

---

## ✨ Features
- 🛠 Built with **Cobra** (CLI framework for Go)  
- 🤖 Telegram bot powered by **Telebot v3**  
- 🔑 Reads bot token from `TELE_TOKEN` environment variable  
- 📦 Version injection via `-ldflags` during build  
- 📜 Simple command handlers (`/start hello`, `/version`, etc.)  
- ⚡ Fast, minimal, and easy to extend  

---

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/cyberm0nk/tbot.git
cd tbot
