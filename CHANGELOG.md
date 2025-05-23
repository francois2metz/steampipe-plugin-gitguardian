## v0.3.0 [2025-03-01]

_What's new?_

- Update go to 1.23
- Update steampipe sdk to 5.11.3

## v0.2.0 [2023-10-15]

_What's new?_

* Update SDK to v5.6.2
* Update go to 1.21

## v0.1.0 [2023-07-30]

_What's new?_

- Add `detector_name`, `detector_display_name`, `detector_nature`, `detector_family`, `detector_group_name`, `detector_group_display_name` to the `gitguardian_secret_incident` table. Thanks [@orf](https://github.com/orf).

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
