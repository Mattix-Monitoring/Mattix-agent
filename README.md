# Mattix-Agent

> A lightweight system monitoring agent written in Go

Mattix is a simple, fast and linux-plataform monitoring agent designed to collect system metrics and expose them through an HTTP API.

It was created as a learning project inspired by tools like **Node Exporter** and **Zabbix Agent**, focusing on simplicity,  and low resource usage.

## Features

- CPU usage
- Memory usage
- Disk usage
- Network statistics
- Hostname
- System uptime
- JSON HTTP API
- Lightweight
- Zero external dependencies (except gopsutil/v4)

## Example

```json
{
  "system": {
    "hostname": "ubuntu",
    "uptime": 2139,
    "updated_at": "2026-07-26T14:18:27.203979246-03:00"
  },
  "hardware": {
    "cpu": {
      "usage": 1.6115351993214586,
      "cores": [
        {
          "id": 0,
          "usage": 3.061224489795918
        },
        {
          "id": 1,
          "usage": 0
        },
        {
          "id": 2,
          "usage": 3.125
        },
        {
          "id": 3,
          "usage": 0
        },
        {
          "id": 4,
          "usage": 2.083333333333333
        },
        {
          "id": 5,
          "usage": 2.0202020202020203
        },
        {
          "id": 6,
          "usage": 4.25531914893617
        },
        {
          "id": 7,
          "usage": 1.0526315789473684
        },
        {
          "id": 8,
          "usage": 1.0101010101010102
        },
        {
          "id": 9,
          "usage": 0
        },
        {
          "id": 10,
          "usage": 1.9801980198019802
        },
        {
          "id": 11,
          "usage": 0
        }
      ]
    },
    "disk": [
      {
        "mouint_point": "/",
        "total": 64424509440,
        "free": 13891784704,
        "used": 49618415616
      },
      {
        "mouint_point": "/home",
        "total": 333063389184,
        "free": 141015748608,
        "used": 189572739072
      },
      {
        "mouint_point": "/boot",
        "total": 205520896,
        "free": 41578496,
        "used": 163942400
      },
      {
        "mouint_point": "/boot/efi",
        "total": 205520896,
        "free": 41578496,
        "used": 163942400
      }
    ],
    "memory": {
      "total": 16402276352,
      "used": 4650254336,
      "free": 11752022016
    },
    "network": {
      "name": "wlp0s",
      "mac": "28:0c:50:db:5c:6d",
      "ipv4": "192.168.18.111",
      "rx_bytes": 40038472,
      "tx_bytes": 8576432,
      "rx_speed": 0,
      "tx_speed": 0
    },
    "temperature": {
      "cpu": 46000
    }
  }
}
```

## Running

```bash
git clone https://github.com/matesu777/Mattix-agent

cd Mattix-agent

go run .
```


The API will be available at

```
http://localhost:8080/metrics
```

---

## Roadmap

- [x] CPU monitoring
- [x] Memory monitoring
- [x] Disk monitoring
- [x] Network monitoring
- [x] Hostname
- [x] Uptime

### Next

- [X] Goroutines
- [X] Temperature
- [X] Per-core CPU usage
- [X] Multiple disks
- [X] Multiple network interfaces
- [ ] Configuration file

---

### Languages 
- [Portuguese](/README.pt-BR.md)
- [English](/README.md)
