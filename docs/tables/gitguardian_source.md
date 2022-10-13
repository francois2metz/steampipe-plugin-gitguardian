# Table: gitguardian_source

Returns the sources monitored.

## Examples

### List sources

```sql
select
  id,
  url
from
  gitguardian_source;
```

### List at risk sources

```sql
select
  id,
  url,
  health
from
  gitguardian_source
where
  health='at_risk';
```

### List unscanned sources

```sql
select
  id,
  url,
  health
from
  gitguardian_source
where
  last_scan_date is null;
```

### List sources where the last scan failed

```sql
select
  id,
  url,
  health
from
  gitguardian_source
where
  last_scan_status='failed';
```

### List source with open incidents

```sql
select
  id,
  url,
  open_incidents_count
from
  gitguardian_source
where
  open_incidents_count > 0;
```
