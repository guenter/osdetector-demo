## Setup Cassandra

```
CREATE KEYSPACE browsers WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE browsers.browser_counts (
  counter counter,
  os varchar,
  PRIMARY KEY (os)
);
```

## Usage

```
Run the server

Usage:
  osdetector-demo serve [flags]

Flags:
      --address string               The address to run the HTTP server on (default ":8080")
      --cassandra-host stringArray   Cassandra hosts to connect to (default [127.0.0.1])
      --cassandra-keyspace string    The keyspace in Cassandra (default "browsers")
  -h, --help                         help for serve
      --template string              The filename of the template (default "index.html")

Global Flags:
      --config string   config file (default is $HOME/.osdetector-demo.yaml)
```
