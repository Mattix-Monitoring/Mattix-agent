# Mattix-Agent

> Um agente de monitoramento de sistema leve escrito em Go

O Mattix é um agente de monitoramento simples, rápido e para plataformas Linux, projetado para coletar métricas do sistema e expô-las através de uma API HTTP.

Ele foi criado como um projeto de aprendizado inspirado em ferramentas como o **Node Exporter** e o **Zabbix Agent**, focando em simplicidade, e baixo uso de recursos.

## Funcionalidades

- Uso de CPU
- Uso de memória
- Uso de disco
- Estatísticas de rede
- Nome do host (Hostname)
- Tempo de atividade do sistema (Uptime)
- API HTTP JSON
- Leve
- Zero dependências externas (exceto gopsutil/v4)

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

## Executando

```bash
git clone https://github.com/matesu777/Mattix-agent

cd Mattix-agent

go run .
```

A API estará disponível em

```
http://localhost:8080/metrics
```

---

## Roteiro (Roadmap)

- [x] Monitoramento de CPU
- [x] Monitoramento de memória
- [x] Monitoramento de disco
- [x] Monitoramento de rede
- [x] Nome do host (Hostname)
- [x] Tempo de atividade (Uptime)

### Próximos passos

- [X] Goroutines
- [X] Temperatura
- [X] Uso de CPU por núcleo
- [X] Múltiplos discos
- [ ] Múltiplas interfaces de rede
- [ ] Arquivo de configuração

---

### Languages 
- [Portuguese](/README.pt-BR.md)
- [English](/README.md)
