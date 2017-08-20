## Setup

```
CREATE KEYSPACE browsers WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE browsers.browser_counts (
  counter counter,
  os varchar,
  PRIMARY KEY (os)
);

```
