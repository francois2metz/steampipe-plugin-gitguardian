## v0.0.5 [2023-07-23]

_What's new?_

- Update steampipe sdk to v5
- Update go-gitguardian to v0.3.6

## v0.0.4 [2022-11-04]

_What's new?_

- Add `gitguardian_member` table (this require a new permission scope: members:read)

## v0.0.3 [2022-10-20]

_What's new?_

- Fix typo
- Add examples
- Update GitGuardian case

## v0.0.2 [2022-10-08]

_What's new?_

- Limit PerPage setting when using a LIMIT clause
- Add `gitguardian_audit_log` table (this require a new permission scope: audit_logs:read)

## v0.0.1 [2022-10-07]

_What's new?_

- Initial release with 2 tables:
  - `gitguardian_secret_incident`
  - `gitguardian_source`
