# Event Schemas

Domain event schemas live here.

Events use RabbitMQ through the `domain.events` topic exchange. Event names should follow:

```text
domain.entity.action_past_tense
```

Examples:

```text
finance.payment.verified
academic.report_card.published
school.student.created
```

## Update Rules

- Add or update JSON schemas when a domain event changes.
- Keep event payloads typed and versioned.
- Keep consumers idempotent.
- Do not include tokens, passwords, raw document content, or Confidential details.
- Include tenant context when relevant.
- Coordinate event schema changes with producers, consumers, reporting, communication, and QA.

## Structure

```text
envelope.schema.json
identity/
school/
admission/
academic/
finance/
communication/
reporting/
```
