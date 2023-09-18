# Symphonic Skeleton

## Developer Guide

### Installation

```shell
// Download the skeleton
git clone git@github.com:fwidjaya20/symphonic-skeleton.git && rm -rf symphonic-skeleton/.git*

// Installation dependencies
cd symphonic-skeleton && go mod tidy

// Create .env environment configuration file
cp .env.example .env
```

## Architecture

```text
./src/post
├── application
│ ├── command
│ ├── handler
│ └── query
├── domain
│ ├── entity
│ ├── model
│ ├── repository
│ └── value-object
└── infrastructure
    └── repository
```

- **Application**: is responsible for coordinating and orchestrating the execution of use cases or application-specific workflows. It acts as a bridge between the Domain Layer and the Infrastructure Layer.
- **Domain**: represents the core business logic of the application. It contains entities, value objects, aggregates, and domain services that model the business domain's concepts and rules.
- **Infrastructure**: provides the technical foundation and support for the application. It handles infrastructure concerns such as data storage, communication with external systems, and user interfaces.