# Table: gitguardian_member

Retrieve details about workspace members.

## Examples

### List all members

```sql
select
  id,
  name,
  email
from
  gitguardian_member;
```

### List owner members

```sql
select
  id,
  name,
  email
from
  gitguardian_member
where
  role='owner';
```

### List manager members

```sql
select
  id,
  name,
  email
from
  gitguardian_member
where
  role='manager';
```
