# Proto Contracts

Internal gRPC protobuf contracts live here.

Sprint 0 only prepares contract-first folders. Do not define final service RPCs until the related sprint task is active and the service boundary/API contract has been reviewed.

## Package Naming

Use this pattern:

```text
schoolplatform.<domain>.v1
```

Examples:

```text
schoolplatform.identity.v1
schoolplatform.finance.v1
schoolplatform.common.v1
```

## Folder Structure

```text
identity/v1/
schoolcore/v1/
admission/v1/
academic/v1/
finance/v1/
communication/v1/
reporting/v1/
common/v1/
```

## Compatibility Rules

- Do not reuse deleted field numbers.
- Use `reserved` for removed fields and names.
- Keep request and response messages explicit.
- Do not include tokens, passwords, or Confidential details.
- Update service docs when proto contracts change.
