# QIE Blockchain

**QIE Blockchain** is a high-performance, EVM-compatible blockchain built for scalability, interoperability, and developer-friendly decentralized application deployment.

It provides fast consensus, low transaction costs, and cross-chain capabilities to support next-generation Web3 infrastructure and applications.

---

## 🌐 Overview

QIE Blockchain offers a powerful platform for building decentralized applications using familiar Ethereum tools while leveraging a modern and scalable blockchain architecture.

### Key Features

* ⚡ Fast block finality
* 💰 Low transaction fees
* 🧠 Ethereum smart contract compatibility
* 🔗 Cross-chain interoperability support
* 🛠 Developer-friendly infrastructure
* 🌍 Open-source and community-driven

---

## 🧱 Technology Stack

* Cosmos SDK-based architecture
* CometBFT consensus engine
* Ethereum Virtual Machine (EVM) compatibility
* IBC-enabled cross-chain communication

---

## 🚀 Getting Started

### Prerequisites

Ensure the following are installed:

* Go **1.21+**
* Git
* Make

---

### Clone Repository

```bash
git clone https://github.com/qieadmin/QIE-Blockchain.git
cd QIE-Blockchain
```

---

### Build (Recommended)

```bash
make install
```

---

### Alternative Build Method

```bash
cd cmd/qied
go build
```

---

### Add Binary to Global Path

```bash
cp qied /usr/bin/
```

*(You may need `sudo` depending on your system permissions.)*

---

### Initialize Node

```bash
qied init my-node --chain-id qie_1990-1
```

---

### Start Node

```bash
qied start
```

---

## 🔧 Configuration

After initialization, node configuration files are stored in:

```
~/.qied/
```

Important configuration files:

* `config/config.toml`
* `config/app.toml`
* `genesis.json`

---

## 🧪 Testing

Run the test suite:

```bash
make test
```

---

## 📡 Network Information

| Parameter           | Value          |
| ------------------- | -------------- |
| **Chain ID**        | `qie_1990-1`   |
| **Token Symbol**    | `QIE`          |
| **Consensus**       | CometBFT       |
| **Smart Contracts** | EVM Compatible |

---

## 🤝 Contributing

We welcome contributions from the community.

1. Fork the repository
2. Create a new branch for your feature or fix
3. Commit your changes with clear messages
4. Open a Pull Request

Please ensure:

* The project builds successfully
* Tests pass
* Code follows clean standards

---

## 📜 License

This project is licensed under the **Apache 2.0 License**.
See the `LICENSE` file for full details.

---

## 📬 Contact & Resources

* GitHub: https://github.com/qieadmin/QIE-Blockchain
* Website: https://qie.digital
* Documentation: https://docs.qie.digital

---

## ⚠️ Disclaimer

QIE Blockchain is under active development and may change frequently.
Use in production environments at your own discretion.
