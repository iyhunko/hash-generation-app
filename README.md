## How to test this app?

1. Run `make build` command
2. Start a worker to refresh token every interval: `./bin/refresh_hash_worker`
   - interval = 5 seconds by default. Provide **HASH_GENERATION_INTERVAL** env variable to change the value( example: `HASH_GENERATION_INTERVAL=5m ./bin/refresh_hash_worker` )
   - we use a file to save cache(just for PoC, better storage should be used in real project)
3. Start HTTP API server: `./bin/http_api`
   - default port is 8001
   - see possible env variables to configure HTTP server here: **internal/config/config.go**
4. Hit http://localhost:8001/hash URL to test response.
    - response example: ```{"hash":"11cdc662-38b6-40b7-b79c-17d6d0bb1749","generated_at":"2023-03-20T13:50:24.503731+02:00"}```
5. Start GRPC API server: `./bin/grcp_api`
   - default port is 8002
   - see possible env variables to configure GRPC server here: **internal/config/config.go**
6. Run GRPC client: `./bin/grcp_client`
   - You will see next logs with info about success connection/response
```{"level":"info","timestamp":"2023-03-20T14:13:18.303726+02:00","message":"Starting GRPC api client"}```
```{"level":"info","timestamp":"2023-03-20T14:13:18.321008+02:00","message":"Uuid: f88e70cb-493a-486f-bba8-be7c63522360, Time: 2023-03-20 14:13:14.324162 +0200 EET"}```
   - see GRPC client details here: **cmd/grpc_client/main.go**
7. Check **Makefile** to see available commands(code test, checks, build, proto etc)
