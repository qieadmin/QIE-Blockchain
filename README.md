# QIE Blockchain

**QIE Blockchain** is a high-performance, EVM-compatible blockchain designed for scalability, interoperability, and developer-friendly decentralized application deployment.

It delivers fast consensus, low fees, and cross-chain capabilities to support next-generation Web3 applications and infrastructure.

---

## 🌐 Overview

QIE Blockchain provides a robust environment for building decentralized applications using familiar Ethereum tooling while benefiting from a modern, scalable blockchain architecture.

### Key Features

* ⚡ Fast block finality
* 💰 Low transaction fees
* 🧠 Ethereum smart contract compatibility
* 🔗 Cross-chain interoperability support
* 🛠 Developer-friendly infrastructure
* 🌍 Open-source and community driven

---

## 🧱 Technology Stack

* Cosmos SDK-based architecture
* CometBFT consensus engine
* Ethereum Virtual Machine compatibility
* IBC-enabled cross-chain communication

---

## 🚀 Getting Started

### Prerequisites

* Go >= 1.21
* Git
* Make

### Clone Repository

```bash
git clone https://github.com/qieadmin/QIE-Blockchain.git
cd QIE-Blockchain
```

### Build

```bash
make install
```

### Initialize Node

```bash
qied init my-node --chain-id qie_1990-1
```

### Start Node

```bash
qied start
```

---

## 🔧 Configuration

After initialization, configuration files will be located at:

```
~/.qied/
```

Important files:

* `config/config.toml`
* `config/app.toml`
* `genesis.json`

---

## 🧪 Testing

Run unit tests:

```bash
make test
```

---

## 📡 Network Information

| Parameter       | Value          |
| --------------- | -------------- |
| Chain ID        | `qie_1990-1`   |
| Token Symbol    | `QIE`          |
| Consensus       | CometBFT       |
| Smart Contracts | EVM Compatible |

---

## 🤝 Contributing

We welcome community contributions.

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Submit a Pull Request

Please ensure your code builds successfully and tests pass before submitting.

---

## 📜 License

This project is licensed under the **Apache 2.0 License**.
See the LICENSE file for details.

---

## 📬 Contact & Resources

* GitHub: https://github.com/qieadmin/QIE-Blockchain
* Website: https://qie.digital
* Documentation: https://docs.qie.digital
  
---

## ⚠️ Disclaimer

QIE Blockchain is under active development.
Use in production environments at your own risk.
