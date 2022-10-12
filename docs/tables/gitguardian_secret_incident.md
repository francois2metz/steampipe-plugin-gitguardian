# Table: gitguardian_secret_incident

Returns the leaked incidents detected.

## Examples

### List all incidents

```sql
select
  id,
  date,
  status
from
  gitguardian_secret_incident;
```

### List open incidents

```sql
select
  id,
  date,
  status
from
  gitguardian_secret_incident
where
  status in ('TRIGGERED', 'ASSIGNED');
```

### List shared incidents

```sql
select
  id,
  date,
  status,
  share_url
from
  gitguardian_secret_incident
where
  share_url is not null;
```

### Get incidents sorted by the number of occurrences

```sql
select
  id,
  date,
  status,
  occurrences_count
from
  gitguardian_secret_incident
order by
  occurrences_count desc;
```

### Get last month incidents

```sql
select
  id,
  date,
  status
from
  gitguardian_secret_incident
where
  date>(current_date - interval '1' month);
```
