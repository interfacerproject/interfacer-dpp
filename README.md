# Interfacer Digital Product Passport API

![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)
![MongoDB](https://img.shields.io/badge/MongoDB-5.0+-green.svg)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)

<br>
<a href="https://www.interfacerproject.eu/">
  <img alt="Interfacer project" src="https://dyne.org/images/projects/Interfacer_logo_color.png" width="350" />
</a>
<br>

A specialized microservice for managing Digital Product Passports (DPP) within the Interfacer Project ecosystem. Built with Go, Gin framework, and MongoDB, this service is designed to serve the interfacer-gui through interfacer-proxy and provides DPP EECC (Extended Environmental and Circularity Certificate) type functionality.

> ⚠️ **Early Development Stage**: This microservice is currently in early stage development and is being actively developed as part of the broader Interfacer Project infrastructure.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Data Model](#data-model)
- [Usage Examples](#usage-examples)
- [Docker Deployment](#docker-deployment)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## 🌟 Overview

The Interfacer Digital Product Passport microservice is a core component of the [Interfacer Project](https://www.interfacerproject.eu/) ecosystem, specifically designed to handle DPP EECC (Extended Environmental and Circularity Certificate) type data. This microservice operates within a distributed architecture where:

- **interfacer-gui** (frontend) communicates with this service
- **interfacer-proxy** acts as the API gateway and routing layer
- **interfacer-dpp** (this service) manages DPP data and business logic

The service supports circular economy initiatives by providing specialized DPP EECC functionality, enabling comprehensive tracking and management of product environmental and circularity information throughout the product lifecycle.

### What is a Digital Product Passport (DPP) EECC Type?

A Digital Product Passport EECC (Extended Environmental and Circularity Certificate) type is a specialized digital record focused on environmental and circularity aspects of products throughout their lifecycle. The EECC type specifically emphasizes:

- **Environmental Impact Assessment**: Detailed carbon footprint, water consumption, and chemical usage metrics
- **Circularity Metrics**: Reparability scores, recyclability information, and material flow data
- **Extended Producer Responsibility**: Comprehensive tracking for compliance with circular economy regulations
- **Lifecycle Management**: Repair, refurbishment, and end-of-life processing information
- **Supply Chain Transparency**: Economic operator information and material sourcing details
- **Compliance Certifications**: CE marking, RoHS compliance, and environmental standards

## ✨ Features

- **Complete CRUD Operations**: Create, read, update, and delete Digital Product Passports
- **Comprehensive Data Model**: Supports all major DPP data categories
- **RESTful API**: Clean and intuitive REST endpoints
- **MongoDB Integration**: Scalable NoSQL database storage
- **Docker Support**: Containerized deployment ready
- **JSON API**: Standardized JSON request/response format
- **Error Handling**: Proper HTTP status codes and error messages
- **Performance**: Optimized database operations with timeouts

## 🏗️ Architecture

```
interfacer-dpp/
├── cmd/
│   └── main/
│       └── main.go           # Application entry point
├── internal/
│   ├── database/
│   │   └── database.go       # MongoDB connection and configuration
│   ├── handler/
│   │   └── handler.go        # HTTP request handlers
│   └── model/
│       └── model.go          # Data models and structures
├── Dockerfile                # Docker configuration
├── go.mod                    # Go module dependencies
├── go.sum                    # Go module checksums
└── README.md                 # This file
```

### Technology Stack

- **Language**: Go 1.24.0
- **Web Framework**: Gin HTTP web framework
- **Database**: MongoDB with official Go driver
- **Containerization**: Docker
- **Architecture Pattern**: Clean Architecture with separation of concerns

## 📋 Prerequisites

Before running this application, ensure you have the following installed:

- **Go**: Version 1.24.0 or higher
- **MongoDB**: Version 5.0 or higher
- **Docker** (optional): For containerized deployment

## 🚀 Installation

### Local Development Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/interfacerproject/interfacer-dpp.git
   cd interfacer-dpp
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Start MongoDB**:
   ```bash
   # Using MongoDB service
   brew services start mongodb/brew/mongodb-community
   
   # Or using Docker
   docker run -d -p 27017:27017 --name mongodb mongo:latest
   ```

4. **Run the application**:
   ```bash
   go run cmd/main/main.go
   ```

The API will be available at `http://localhost:8080`

## ⚙️ Configuration

### Environment Variables

The application uses the following default configuration:

- **MongoDB URI**: `mongodb://localhost:27017`
- **Database Name**: `dpp_db`
- **Collection Name**: `passports`
- **Server Port**: `8080`

To customize these settings, modify the constants in `internal/database/database.go` or use environment variables (implementation can be extended).

### Database Setup

The application automatically creates the necessary database and collection on first run. No manual database setup is required.

## 📚 API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/dpp` | Create a new Digital Product Passport |
| `GET` | `/dpp/{id}` | Retrieve a specific DPP by ID |
| `PUT` | `/dpp/{id}` | Update an existing DPP |
| `DELETE` | `/dpp/{id}` | Delete a DPP |
| `GET` | `/dpps` | Retrieve all DPPs |

### HTTP Status Codes

- `200 OK`: Successful GET, PUT operations
- `201 Created`: Successful POST operation
- `400 Bad Request`: Invalid request format or parameters
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server-side error

### Request/Response Format

All requests and responses use JSON format with `Content-Type: application/json`.

## 📊 Data Model

The Digital Product Passport model includes the following main sections:

### Core Structure (Provisional)

> ⚠️ **Note**: The data structure is currently provisional and subject to change during development.

| Category | Description | Example Fields |
|----------|-------------|----------------|
| **Product Overview** | Basic product identification and specifications | Brand name, model, GTIN, dimensions, warranty |
| **Environmental Impact** | Environmental footprint metrics | CO2 emissions, water consumption, energy usage |
| **Reparability** | Repair and maintenance information | Service instructions, spare parts availability |
| **Recyclability** | End-of-life and material information | Material composition, recycling instructions |
| **Energy & Efficiency** | Power and energy-related specifications | Battery type, charging time, power ratings |
| **Compliance & Standards** | Regulatory compliance information | CE marking, RoHS compliance, certifications |
| **Components** | Sub-component and part information | Component descriptions, GTINs, linked DPPs |
| **Economic Operator** | Manufacturer and supplier details | Company information, contact details, addresses |
| **Repair Information** | Historical repair records | Repair actions, materials used, dates |
| **Refurbishment Information** | Refurbishment history and details | Refurbishment actions, materials, dates |
| **Recycling Information** | End-of-life processing records | Recycling actions, processing dates |

### Key Data Categories

1. **Product Overview**: Basic product information, specifications, and identification
2. **Environmental Impact**: Carbon footprint, water consumption, chemical usage
3. **Reparability**: Service instructions and spare parts availability
4. **Recyclability**: Material composition and recycling instructions
5. **Energy & Efficiency**: Battery information, power ratings, charging specifications
6. **Compliance**: CE marking, RoHS compliance, and certifications
7. **Economic Operator**: Manufacturer and supplier information
8. **Lifecycle Information**: Repair, refurbishment, and recycling records

## 💡 Usage Examples

### Create a Digital Product Passport

```bash
curl -X POST http://localhost:8080/dpp \
  -H "Content-Type: application/json" \
  -d '{
    "productOverview": {
      "brandName": "EcoTech",
      "productName": "Sustainable Drill Pro",
      "modelName": "SDP-2024",
      "gtin": "1234567890123",
      "color": "Green",
      "warrantyDuration": "2 years"
    },
    "environmentalImpact": {
      "co2eEmissionsPerUnit": "5.2 kg CO2e",
      "energyConsumptionPerUnit": "150 kWh"
    },
    "reparability": {
      "serviceAndRepairInstructions": "Available online and via QR code",
      "availabilityOfSpareParts": "10 years from purchase date"
    }
  }'
```

### Retrieve a Specific DPP

```bash
curl -X GET http://localhost:8080/dpp/60f7b3b4e4b0c8a5f8e4b0c8
```

### Update a DPP

```bash
curl -X PUT http://localhost:8080/dpp/60f7b3b4e4b0c8a5f8e4b0c8 \
  -H "Content-Type: application/json" \
  -d '{
    "productOverview": {
      "brandName": "EcoTech Updated",
      "warrantyDuration": "3 years"
    }
  }'
```

### Get All DPPs

```bash
curl -X GET http://localhost:8080/dpps
```

### Delete a DPP

```bash
curl -X DELETE http://localhost:8080/dpp/60f7b3b4e4b0c8a5f8e4b0c8
```

## 🐳 Docker Deployment

### Build and Run with Docker

1. **Build the Docker image**:
   ```bash
   docker build -t interfacer-dpp .
   ```

2. **Run MongoDB container**:
   ```bash
   docker run -d --name mongodb -p 27017:27017 mongo:latest
   ```

3. **Run the application container**:
   ```bash
   docker run -d --name interfacer-dpp-api \
     --link mongodb:mongodb \
     -p 8080:8080 \
     interfacer-dpp
   ```

### Docker Compose (Recommended)

Create a `docker-compose.yml` file:

```yaml
version: '3.8'

services:
  mongodb:
    image: mongo:latest
    container_name: dpp-mongodb
    ports:
      - "27017:27017"
    volumes:
      - mongodb_data:/data/db

  api:
    build: .
    container_name: dpp-api
    ports:
      - "8080:8080"
    depends_on:
      - mongodb
    environment:
      - MONGODB_URI=mongodb://mongodb:27017

volumes:
  mongodb_data:
```

Run with:
```bash
docker-compose up -d
```

## 🛠️ Development

### Project Structure

The project follows Go best practices with a clean architecture:

- **`cmd/`**: Application entry points
- **`internal/`**: Private application code
  - **`database/`**: Database connection and configuration
  - **`handler/`**: HTTP request handlers (controllers)
  - **`model/`**: Data models and business logic

### Adding New Features

1. **Models**: Add new data structures in `internal/model/model.go`
2. **Handlers**: Implement business logic in `internal/handler/handler.go`
3. **Routes**: Register new endpoints in `cmd/main/main.go`
4. **Database**: Extend database operations in `internal/database/database.go`

### Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Handle errors appropriately
- Use context for database operations with timeouts

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## 🤝 Contributing

We welcome contributions to the Interfacer Digital Product Passport API! Please follow these guidelines:

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Commit your changes**: `git commit -m 'Add amazing feature'`
4. **Push to the branch**: `git push origin feature/amazing-feature`
5. **Open a Pull Request**

### Contribution Guidelines

- Follow Go best practices and coding standards
- Add tests for new functionality
- Update documentation as needed
- Ensure all tests pass before submitting
- Use clear, descriptive commit messages

---

## 😍 Acknowledgements

<a href="https://dyne.org">
  <img src="https://files.dyne.org/software_by_dyne.png" width="222">
</a>

Copyleft (ɔ) 2022 by [Dyne.org](https://www.dyne.org) foundation, Amsterdam

Designed, written and maintained by Ennio Donato, Micol Salomone, Giovanni Abbatepaolo and Puria Nafisi Azizi.

**[🔝 back to top](#toc)**

---

## 🌐 Links

https://www.interfacer.eu/

https://dyne.org/

**[🔝 back to top](#toc)**

---

## 💼 License

    Interfacer Digital Product Passport API
    Copyleft (ɔ) 2022 Dyne.org foundation

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
