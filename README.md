## Package
This package allows you to define a constraint to your graphql `input`.

## Requirements
- [go](https://go.dev) >= 1.18
- [99designs/gqlgen](https://github.com/99designs/gqlgen) >= 0.17 

## Setup
1. Add this library to your go project with:
```shell
go get -u github.com/jacoz/gqlgen-oneof-directive
go mod tidy
```

2. Add to your `schema.graphqls` file the directive and then use it like the example below:
```graphql
directive @oneOf on INPUT_OBJECT

# ...

input PaymentRequestInput @oneOf {
    card: CardPaymentRequestInput
    sepa: SepaPaymentRequestInput
    invoice: InvoicePaymentRequestInput
}
```

3. Regenerate the graphql schema with the command
```shell
go run github.com/99designs/gqlgen generate
```

4. Load the directive into your graphql server configuration: 
```go
c := generated.Config{}
c.Directives.OneOf = oneof.Directive
```
