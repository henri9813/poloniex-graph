# Poloniex-Graph

This project permit retrieve your poloniex Balance and send it to a InfluxDB instance.

## Setup

Copy the `.env.dist` to `.env`
```
cp .env.dist .env
```

And feed the differents fields

## Usage

You can use `henri9813/poloniex-graph` docker image.

```
docker-compose up
```

> Comment the `build` line 

`grafana-dashboard.json` contain a dashboard you can import into the Grafana.
