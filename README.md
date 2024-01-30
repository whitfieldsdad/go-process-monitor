# Process monitor

⚠️ This project is under active development ⚠️

## Features

- Track the creation of new processes on macOS without privileged access
- JSONL output

## Usage

```bash
go run main.go | jq
```

```json
...
{
  "header": {
    "id": "0bc1f3c3-983e-4473-accd-32bf10913172",
    "time": "2024-01-30T14:51:09.567093-05:00",
    "object_type": "process",
    "event_type": "started"
  },
  "data": {
    "pid": 94878,
    "ppid": 67981,
    "name": "netstat",
    "argv": [
      "netstat"
    ],
    "argc": 1,
    "command_line": "netstat",
    "create_time": "2024-01-30T14:51:09.554-05:00",
    "executable": {
      "path": "/usr/sbin/netstat",
      "filename": "netstat",
      "hashes": {
        "md5": "3a31cacef8c534211b49bbcca577ec59",
        "sha1": "909b6a3429c4e9a369bb1f7a15f68ab7dbfd42db",
        "sha256": "63bb37395f9998183d0be827b9824d590ff469c8151328308135368bceeb8c9f",
        "xxh3": 15253356458188016367
      }
    }
  }
}
{
  "header": {
    "id": "89ddb038-503c-42b7-8d54-6349ed7cca58",
    "time": "2024-01-30T14:51:10.443681-05:00",
    "object_type": "process",
    "event_type": "started"
  },
  "data": {
    "pid": 94885,
    "ppid": 67981,
    "name": "whoami",
    "create_time": "2024-01-30T14:51:10.437-05:00",
    "executable": {
      "path": "/usr/bin/whoami",
      "filename": "whoami",
      "hashes": {
        "md5": "9998866dc9ea32e4b4cff7ce737272ab",
        "sha1": "3c62dad9b22c6bf437c4eb4dd73a13175d326575",
        "sha256": "c4167b65515e95be93ecb3cdc555096bb088bccaeb7ee22cc0f817d040761b25",
        "xxh3": 6707975974401128194
      }
    }
  }
}
...
```
