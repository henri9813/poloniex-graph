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

### Grafana

You need to fill the datasource in the Grafana interface.

You can import the `grafana-dashboard.json` which contain the dashboard.

## Developpement

You can use the `docker-compose.dev.yaml` for development purpose and test

```
docker-compose -f docker-compose.dev.yaml up
```
