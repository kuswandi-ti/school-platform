# Pull Request

## Summary

<!-- Jelaskan perubahan secara singkat dan jelas. -->

## Related Issue

<!-- Link issue terkait. Gunakan format: Closes #123 / Fixes #123 / Related #123 -->

Closes #

## Type of Change

Pilih satu atau lebih:

- [ ] `type: feature` — fitur baru
- [ ] `type: bug` — perbaikan bug
- [ ] `type: chore` — pekerjaan maintenance
- [ ] `type: docs` — dokumentasi
- [ ] `type: refactor` — refactor tanpa perubahan behavior
- [ ] `type: test` — test baru/perubahan test
- [ ] `type: security` — security hardening/fix
- [ ] `type: infra` — infrastructure/CI/CD/deployment
- [ ] `type: hotfix` — hotfix production

## Affected Area

Pilih area yang terdampak:

- [ ] `area: api-gateway`
- [ ] `area: identity`
- [ ] `area: school-core`
- [ ] `area: admission`
- [ ] `area: academic`
- [ ] `area: finance`
- [ ] `area: communication`
- [ ] `area: reporting`
- [ ] `area: web-admin`
- [ ] `area: mobile`
- [ ] `area: infra`
- [ ] `area: docs`
- [ ] `area: security`
- [ ] `area: observability`
- [ ] `area: ci-cd`
- [ ] `area: file-management`

## Sprint / Milestone

<!-- Contoh: Sprint 0 — Project Foundation -->

Sprint/Milestone:

## Branch Flow

Target branch:

- [ ] `feature/*` → `develop`
- [ ] `fix/*` → `develop`
- [ ] `docs/*` → `develop`
- [ ] `chore/*` → `develop`
- [ ] `develop` → `staging`
- [ ] `staging` → `main`
- [ ] `hotfix/*` → `main`

## Changes

Tuliskan perubahan utama:

- 
- 
- 

## Scope

Yang termasuk dalam PR ini:

- 
- 
- 

## Out of Scope

Yang tidak termasuk dalam PR ini:

- 
- 
- 

## Screenshots / Recording

<!-- Wajib jika ada perubahan UI. Bisa kosong jika tidak ada UI changes. -->

| Before | After |
|---|---|
|  |  |

## API / Contract Changes

- [ ] Tidak ada perubahan API/gRPC/event contract
- [ ] REST/OpenAPI berubah
- [ ] gRPC/proto berubah
- [ ] RabbitMQ event contract berubah
- [ ] Request/response schema berubah
- [ ] Backward compatible
- [ ] Breaking change

Jika berubah, jelaskan:

```text
Endpoint/proto/event:
Change:
Impact:
Migration/compatibility note:
```

## Database / Migration Changes

- [ ] Tidak ada perubahan database
- [ ] Migration baru ditambahkan
- [ ] Query sqlc berubah
- [ ] Seed data berubah
- [ ] Migration backward-compatible
- [ ] Membutuhkan backup sebelum deploy
- [ ] Membutuhkan manual data migration

Jelaskan jika ada:

```text
Service DB:
Migration:
Rollback consideration:
```

## Event / Queue Impact

- [ ] Tidak ada event/queue impact
- [ ] Event baru dipublish
- [ ] Event consumer baru ditambahkan
- [ ] Retry/DLQ behavior berubah
- [ ] Idempotency sudah dipastikan

Jelaskan jika ada:

```text
Exchange/queue:
Event name:
Producer:
Consumer:
Idempotency key:
```

## Security, Permission, and Scope Checklist

- [ ] Permission/RBAC dicek
- [ ] ABAC/scope dicek
- [ ] Object-level authorization dicek untuk resource by ID
- [ ] Parent/student/teacher/school/foundation scope aman
- [ ] Tidak ada data lintas `foundation_id` / `school_id`
- [ ] Tidak ada direct query ke database service lain
- [ ] API Gateway tidak berisi business logic
- [ ] Tidak ada token/password/secret di log
- [ ] Tidak ada secret/file `.env` yang ikut commit
- [ ] Rate limit ditambahkan jika endpoint sensitif
- [ ] Input validation ditambahkan
- [ ] SQL query parameterized / sqlc-safe

## Audit Checklist

- [ ] Tidak membutuhkan audit log
- [ ] Audit log ditambahkan untuk aksi sensitif
- [ ] Audit menyimpan actor/context/request_id/correlation_id
- [ ] Audit payload tidak menyimpan data Confidential mentah
- [ ] Download/export Restricted/Confidential diaudit jika relevan

Aksi yang diaudit:

```text
-
-
```

## File and Privacy Checklist

- [ ] Tidak ada file handling
- [ ] File private by default
- [ ] File access melalui backend authorization
- [ ] Signed URL digunakan jika diperlukan
- [ ] MIME/extension/size validation ditambahkan
- [ ] Klasifikasi data ditentukan: Public / Internal / Restricted / Confidential
- [ ] Restricted/Confidential data tidak masuk notification body/log/event payload mentah

## Finance Checklist

Isi jika PR menyentuh Finance:

- [ ] Tidak menyentuh Finance
- [ ] Menggunakan decimal/NUMERIC, bukan float
- [ ] Bill generation idempotent
- [ ] Bill menyimpan snapshot amount/policy
- [ ] Payment verification diaudit
- [ ] Void payment butuh approval jika relevan
- [ ] Parent hanya bisa akses bill anak yang terhubung

## AI Agent Usage

- [ ] Tidak menggunakan AI Agent
- [ ] Menggunakan AI Agent untuk draft/implementasi
- [ ] Output AI Agent sudah direview manual
- [ ] AI Agent tidak menangani production secrets
- [ ] AI Agent tidak membuat keputusan final security/legal/compliance
- [ ] AI Agent tidak mengubah arsitektur tanpa instruksi

Dokumen/prompt yang digunakan:

```text
AGENTS.md
SKILLS.md
docs/README.md
docs/09-ai-agent-rules.md
docs/08-coding-standard.md
docs/[active-sprint-plan].md
docs/[active-sprint-task-prompts].md
```

## Tests

Tuliskan test yang sudah dijalankan:

```bash
# contoh
go test ./...
pnpm test
pnpm lint
flutter test
flutter analyze
docker compose config
```

Checklist:

- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] API tests added/updated
- [ ] Permission/scope tests added/updated
- [ ] Event tests added/updated
- [ ] Audit tests added/updated
- [ ] Frontend tests added/updated
- [ ] Mobile tests added/updated
- [ ] Regression tested manually
- [ ] Tidak ada test yang relevan untuk perubahan ini

## QA / UAT Notes

- [ ] QA tidak diperlukan
- [ ] QA diperlukan
- [ ] UAT diperlukan
- [ ] Critical/High bug tidak ada pada core flow
- [ ] Regression checklist sudah dijalankan

QA scenario:

```text
1.
2.
3.
```

## Documentation Checklist

- [ ] Tidak perlu update dokumentasi
- [ ] `docs/README.md` updated
- [ ] `docs/01-technical-architecture.md` updated
- [ ] `docs/02-service-boundary.md` updated
- [ ] `docs/03-data-model-mvp.md` updated
- [ ] `docs/04-api-contract.md` updated
- [ ] `docs/05-event-contract.md` updated
- [ ] `docs/06-ui-screen-user-flow.md` updated
- [ ] `docs/07-test-plan-acceptance-criteria.md` updated
- [ ] `docs/08-coding-standard.md` updated
- [ ] `docs/09-ai-agent-rules.md` updated
- [ ] `docs/10-sprint-backlog-mvp.md` updated
- [ ] `docs/11-github-repository-rules.md` updated
- [ ] `docs/24-local-development-guide.md` updated
- [ ] Sprint plan/task prompt updated if relevant

## Risk Level

- [ ] `risk: low`
- [ ] `risk: medium`
- [ ] `risk: high`
- [ ] `risk: breaking-change`
- [ ] `risk: migration`
- [ ] `risk: data-sensitive`

Risk explanation:

```text

```

## Rollback Plan

Jelaskan cara rollback jika ada masalah:

```text
Rollback steps:
1.
2.
3.

Database rollback consideration:

Feature flag/config rollback:
```

## Deployment Notes

- [ ] Tidak ada deployment note khusus
- [ ] Membutuhkan environment variable baru
- [ ] Membutuhkan secret baru
- [ ] Membutuhkan migration sebelum service start
- [ ] Membutuhkan queue/worker restart
- [ ] Membutuhkan cache clear
- [ ] Membutuhkan manual approval production
- [ ] Membutuhkan backup sebelum deploy

Detail:

```text

```

## Final Checklist

- [ ] Scope sesuai issue
- [ ] Out of scope tidak ikut dikerjakan
- [ ] CI pass
- [ ] Tests pass
- [ ] Lint/format pass
- [ ] Tidak ada secret di commit
- [ ] Tidak ada sensitive data di log
- [ ] Permission/scope sudah dicek
- [ ] Audit/event/file/privacy dicek jika relevan
- [ ] Dokumentasi diupdate jika relevan
- [ ] PR siap direview

## Commit and PR Title Convention

PR title must follow:

```text
type(scope): short description
```

Examples:

```text
feat(identity): add refresh token rotation
fix(finance): prevent duplicate bill generation
docs(workflow): add git commit convention
chore(ci): add repository validation workflow
test(academic): add attendance scope tests
```

Checklist:

- [ ] PR title follows `type(scope): short description`
- [ ] Squash merge title will follow the same convention
- [ ] Commit history does not contain vague messages such as `update`, `fix bug`, or `changes`
- [ ] Migration/security/breaking-change context is documented if relevant
