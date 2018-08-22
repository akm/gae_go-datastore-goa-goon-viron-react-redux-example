# Admin

## Run API server locally

1. Start local server
    ```
    $ make local
    ```
1. Open http://localhost:8081/_ah/login
1. Click login
1. Open developer console of browser
1. Define methods to send request
    ```
    const showStatusCode = r => {
      console.log(r.status)
      return r
    }
    const sendReqWithBody = (method, path, data) => {
      return fetch(path, {method: method, body: JSON.stringify(data)}).then(showStatusCode).then(r => r.json()).then(d => console.log(d))
    }
    const sendReqWithoutBody = (method, path) => {
      return fetch(path, {method: method}).then(showStatusCode).then(r => r.json()).then(d => console.log(d))
    }
    const post = (path, data) => sendReqWithBody("POST", path, data)
    const put = (path, data) => sendReqWithBody("PUT", path, data)
    const get = (path) => sendReqWithoutBody("GET", path)
    const del = (path) => fetch(path, {method: "DELETE"}).then(showStatusCode).then(r => r.text()).then(d => console.log(d))
    ```
1. Send some requests
   ```
   post("/memos", {content: "Memo#1", shared: false})
   get("/memos")
   ```


## Run Viron server locally

1. Start local server
    ```
    $ make local
    ```
1. Open another terminal
1. Start Docker
1. Start Viron
    ```
    $ make run_viron
    ```
1. Open https://localhost:8082
1. Add http://localhost:8081/swagger.json
