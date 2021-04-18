# Poloniex-Graph

This project permit retrieve your poloniex Balance and send it to a InfluxDB instance.

## Environment variables

- POLONIEX_API_SECRET
- POLONIEX_API_KEY
- INFLUXDB_URL
- INFLUXDB_DB
- INFLUXDB_USER
- INFLUXDB_USER_PASSWORD

Copy the `.env.dist` to `.env`
```
cp .env.dist .env
```

## Usage

You can use `henri9813/poloniex-graph` docker image.

```
docker-compose up
```

> Comment the `build` line 

`grafana-dashboard.json` contain a dashboard you can import into the Grafana.
