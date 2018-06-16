# Zendesk-Exporter

## Docker

~~~ shell
docker run -d --name zendesk-exporter -v "<path/to/config/file>:/app/config/zendesk.yml" -p "<ip>:<port>:9146" shakapark/zendesk-exporter:tag
~~~

## Prometheus

```yml
scrape_configs:
  - job_name: 'zendesk'
    metrics_path: /zendesk
    static_configs:
    - targets: ['<ip>:<port>']
```