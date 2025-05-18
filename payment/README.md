# Payment Service

> A modular and extensible payment gateway service built with **NestJS**, **TypeScript**, **MySQL**, and **AWS**, designed to support both **international** and **local** payment methods.

---

## 🚀 Features

- 🧱 **Built with NestJS** – Scalable architecture with dependency injection
- 🌐 **Supports International Payment Methods** – Easily integrate providers like Stripe, PayPal, etc.
- 🇻🇳 **Supports Local Payment Methods** – VNPay, Momo, and more
- ☁️ **Cloud-Ready** – Designed for AWS deployment (Lambda, SQS, SES, RDS, etc.)
- 🛡️ **Secure** – Follows best practices for payment handling and data protection
- 📦 **Modular** – Easy to plug in or remove payment providers
- 🧪 **Testable** – Unit and integration test support

---

## 🏗️ Tech Stack

| Layer         | Tech                     |
|---------------|--------------------------|
| Backend       | [NestJS](https://nestjs.com/) |
| Language      | TypeScript               |
| Database      | MySQL (Amazon RDS)       |
| Cloud         | AWS (Lambda, SQS, SES, etc.) |
| ORM           | TypeORM                  |
| Auth (optional) | JWT or API Key          |

---

## 📁 Project Structure (Simplified)

```

src/
├── payment/
│   ├── providers/        # Each payment provider (Stripe, Momo, etc.)
│   ├── dto/              # Request/response data contracts
│   ├── payment.service.ts
│   ├── payment.module.ts
├── common/
│   └── utils/
├── config/               # Environment and AWS configs
├── main.ts

````

---

## 🧑‍💻 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/phankieuphu/ecom-payment.git
cd ecom-payment
````

### 2. Install Dependencies

```bash
npm install
```

### 3. Environment Configuration

Create a `.env` file from `.env.example`:

```env
DATABASE_URL=mysql://user:password@host:port/dbname
AWS_REGION=ap-southeast-1
AWS_ACCESS_KEY_ID=your_key
AWS_SECRET_ACCESS_KEY=your_secret
PAYMENT_MODE=sandbox
```

### 4. Run the Application

```bash
npm run start:dev
```

---

## 📦 Supported Payment Providers

| Provider | Type            | Status        |
| -------- | --------------- | ------------- |
| Stripe   | International   | ✅ Implemented |
| PayPal   | International   | ⏳ Planned     |
| VNPay    | Local (Vietnam) | ✅ Implemented |
| Momo     | Local (Vietnam) | ✅ Implemented |
| ZaloPay  | Local (Vietnam) | ⏳ Planned     |

---

## 🧪 Testing

```bash
npm run test
```

---

## ☁️ Deployment (AWS)

* Ready for deployment on:

  * AWS Lambda (with Serverless Framework)
  * EC2 or ECS with Docker
  * Amazon RDS for MySQL

---

## 📖 Documentation

* Swagger or Postman collection (planned)
* Provider-specific docs in `docs/` folder (planned)

---

## 🛠 TODO

* [ ] Add unit tests for payment flows
* [ ] Add Swagger API documentation
* [ ] Add PayPal and ZaloPay integration
* [ ] CI/CD pipeline with GitHub Actions

---

## 🤝 Contributing

Pull requests are welcome. Please open an issue first to discuss changes.

---

## 🪪 License

MIT

---

## 👤 Author

* [Phan Kieu Phu](https://github.com/phankieuphu)
