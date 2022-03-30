# Go prometheus exporter + grafana 

This is a simple Go project to exemplify the use of metrics with Prometheus exporter and observability with Grafana.

## Run

1) Start the containers with `docker-compose up`
2) Go's default metrics will be exported at http://localhost:1234/metrics
3) Prometheus will be running at http://localhost:9090/
4) Grafana will be running at http://localhost:3000 and set up your dashboards
5) To increment the custom counter, access http://localhost:1234/count
