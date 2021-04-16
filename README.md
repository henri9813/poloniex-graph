# Poloniex-Graph

This project permit retrieve your poloniex Balance and send it to a InfluxDB instance.

## Setup

Copy the `.env.dist` to `.env`
```
cp .env.dist .env
```

And feed the differents fields

## Usage

```
docker-compose up
```

`grafana-dashboard.json` contain a dashboard you can import into the Grafana.
