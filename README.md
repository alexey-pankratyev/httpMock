# httpmock -- This module with simple http server to procces and logging to file
---
## Launching of httpmock
- You can start httpmock just from the console, start the binary file without parameters and it will be available on port 4000
```shell
▶ ./httpmock 
listening on 4000
Logging to development.log
```
---
## Using curl examples: 
```shell
  get - curl -si "http://localhost:4000/?foo=1&bar=2"
  post - curl -si -X POST -d "test for you" http://localhost:4000/
```





