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
