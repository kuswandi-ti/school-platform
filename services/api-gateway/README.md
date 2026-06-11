# API Gateway

Placeholder for the custom Go API Gateway.

The gateway exposes external REST/JSON APIs, validates tokens, extracts actor and scope context, routes requests, maps REST to internal gRPC calls, standardizes responses, and handles cross-cutting request concerns.

Do not place domain business logic or service-owned database queries here.
