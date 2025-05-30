# saga-cart

saga-cart is a full-stack e-commerce application designed to provide a seamless shopping experience. Built with modern technologies, it features a robust backend implementing the SAGA pattern, specifically using choreography for effective distributed transaction management. The application is scalable, maintainable, and optimized for performance, making it an excellent choice for both developers and users.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [License](#license)

## Architecture

![Architecture Diagram](https://github.com/pepega90/GoStore/blob/main/gostore_saga.png)

go_store follows a microservices architecture, ensuring scalability and maintainability. The backend is built using Go with the Gin framework, handling RESTful APIs and business logic. GORM is utilized for ORM with PostgreSQL as the primary database. Kafka is integrated for handling asynchronous communication between services, implementing the SAGA pattern with choreography to manage distributed transactions effectively.

The frontend is developed with React.js, styled using Tailwind CSS for a modern and responsive user interface. Docker is used for containerizing the application, simplifying deployment and ensuring consistency across different environments.

## Features

- **API Gateway**: Centralized gateway managing all incoming requests and routing them to appropriate services.
- **Product Management:** Add, update, delete, and view products with categories and inventory management.
- **Order Processing:** Users can place orders and view their order history.
- **Product Listing:** Browse and search through a comprehensive list of available products.
- **Order Listing:** View all orders with details such as order status and items included.
- **Responsive Design:** Fully responsive frontend ensuring a seamless experience across devices.
- **Scalable Architecture:** Backend designed with the SAGA pattern using choreography for managing distributed transactions.
- **Real-time Updates:** Leveraging Kafka for real-time data streaming and event-driven communication.

## Tech Stack

### Backend

- **Language:** Go
- **Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL
- **Messaging:** Kafka
- **Containerization:** Docker

### Frontend

- **Library:** React.js
- **Styling:** Tailwind CSS

## Getting Started

Follow these instructions to set up the project locally.

### Prerequisites

- **Go:** [Download and Install Go](https://golang.org/dl/)
- **Node.js & npm:** [Download and Install Node.js](https://nodejs.org/)
- **Docker:** [Download and Install Docker](https://www.docker.com/get-started)
- **Kafka:** [Download and Install Kafka](https://kafka.apache.org/downloads)
- **PostgreSQL:** [Download and Install PostgreSQL](https://www.postgresql.org/download/)

## License

This project is licensed under the [MIT License](LICENSE).

Feel free to customize this README to better fit your project's specifics and to add any additional sections that might be relevant.
