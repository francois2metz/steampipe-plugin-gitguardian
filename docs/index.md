---
organization: francois2metz
category: ["security"]
brand_color: "#081736"
display_name: "Gitguardian"
short_name: "gitguardian"
description: "Steampipe plugin for querying incidents from Gitguardian."
og_description: "Query Gitguardian with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/francois2metz/gitguardian-social-graphic.png"
icon_url: "/images/plugins/francois2metz/gitguardian.svg"
---

# Gitguardian + Steampipe

[Gitguardian](https://www.gitguardian.com/) is a secret scanner of GitHhub or GitLab commits.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  date,
  status
from
  gitguardian_secret_incident
```

```
+---------+----------------------+-----------+
| id      | date                 | status    |
+---------+----------------------+-----------+
| 4460178 | 2022-09-16T08:48:58Z | IGNORED   |
| 4117416 | 2022-08-03T09:06:36Z | IGNORED   |
| 3793634 | 2022-06-22T14:19:03Z | TRIGGERED |
| 2832751 | 2022-03-07T10:06:53Z | TRIGGERED |
| 926032  | 2021-12-14T15:14:40Z | TRIGGERED |
| 926031  | 2021-12-14T15:14:40Z | TRIGGERED |
+---------+----------------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/gitguardian/tables)**

## Get started

### Install

Download and install the latest Gitguardian plugin:

```bash
steampipe plugin install francois2metz/gitguardian
```

### Configuration

Installing the latest gitguardian plugin will create a config file (`~/.steampipe/config/gitguardian.spc`) with a single connection named `gitguardian`:

```hcl
connection "gitguardian" {
  plugin = "francois2metz/gitguardian"

  # Create a personal access token at: https://dashboard.gitguardian.com/api
  # Scope:
  #  - incidents:read
  #  - audit_logs:read
  # token = ""
}
```

You can also use environment variables:

- `GITGUARDIAN_TOKEN`: Your Gitguardian API Key

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-gitguardian
