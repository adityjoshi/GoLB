# ![hiking](https://github.com/adityjoshi/GoLB/assets/111140014/e3ab8e71-6b9b-462d-8318-c9aa5288dbc4)  A Simple Load Balancer in Golang 

## Running it

### Start Multiple Instances of `server.py`

1. **Install Node**:
   ```bash
   brew install node(For MAC)
   ```
2. **Run Server**:
   ```bash
   node server.js
    ```

### Start Load Balancer 
  ```bash
   go run main.go
   ```
### Bombard the loadbalancer with requests
   ```bash
  for i in {1..20}; do curl 127.0.0.1:8000; done
   ```
