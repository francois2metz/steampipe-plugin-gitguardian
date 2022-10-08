# Table: gitguardian_audit_log

Returns the audit log.

## Examples

### List events

```sql
select
  id,
  date,
  event_name,
  action_type
from
  gitguardian_audit_log
order by
  date
```

### List delete events

```sql
select
  id,
  date,
  event_name,
  action_type
from
  gitguardian_audit_log
where
  action_type='DELETE'
```

### List actions from a user

```sql
select
  id,
  date,
  event_name,
  action_type
from
  gitguardian_audit_log
where
  member_email='test@example.net'
```
