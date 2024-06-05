# ![hiking](https://github.com/adityjoshi/GoLB/assets/111140014/e3ab8e71-6b9b-462d-8318-c9aa5288dbc4)  A Simple Load Balancer in Golang 

## Running it

### Start Multiple Instances of `server.py`

1. **Install Flask**:
   ```bash
   pip install flask
   ```
2. **Run Python Server**:
   ```bash
   python3 server.py "server-1" "8001"
   python3 server.py "server-2" "8002"
   python3 server.py "server-3" "8003"
   python3 server.py "server-4" "8004"
   python3 server.py "server-5" "8005"
    ```

### Start Load Balancer 
  ```bash
   go run *.go
   ```
### Bombard the loadbalancer with requests
   ```bash
  for i in {1..20}; do curl 127.0.0.1:8000; done
   ```
