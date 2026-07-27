# Mattix-Agent

> Um agente leve de monitoramento de sistemas escrito em Go

O Mattix é um agente de monitoramento simples, rápido e voltado para Linux, projetado para coletar métricas do sistema e disponibilizá-las por meio de uma API HTTP.

Ele foi criado como um projeto de aprendizado inspirado em ferramentas como **Node Exporter** e **Zabbix Agent**, com foco em simplicidade, desempenho e baixo consumo de recursos.

## Funcionalidades

- Uso da CPU
- Uso da memória
- Uso de disco
- Temperatura da CPU
- Estatísticas de rede
- Hostname
- Uso da CPU por núcleo
- Suporte a múltiplos discos
- Suporte a múltiplas interfaces de rede
- Arquivo de configuração
- Tempo de atividade (uptime)
- API HTTP em JSON
- Leve e eficiente

## Exemplo

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
        "mount_point": "/",
        "total": 64424509440,
        "free": 13891784704,
        "used": 49618415616
      },
      {
        "mount_point": "/home",
        "total": 333063389184,
        "free": 141015748608,
        "used": 189572739072
      },
      {
        "mount_point": "/boot",
        "total": 205520896,
        "free": 41578496,
        "used": 163942400
      },
      {
        "mount_point": "/boot/efi",
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

## Instalação

```bash
curl -fsSL https://raw.githubusercontent.com/Mattix-Monitoring/Mattix-agent/main/scripts/install.sh | bash
```

A API ficará disponível em:

```
/metrics
```

---

### Idiomas

- [Português](/README.pt-BR.md)
- [English](/README.md)
