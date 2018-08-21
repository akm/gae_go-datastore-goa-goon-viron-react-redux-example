# Server

## Development at local

1. Start local server
    ```
    $ make local
    ```
1. Open http://localhost:8080/_ah/login
1. Click login
1. Open developer console of browser
1. Define methods to send request
    ```
    const sendReqWithBody = (method, path, data) => {
      return fetch(path, {method: method, body: JSON.stringify(data)}).then(r => r.json()).then(d => console.log(d))
    }
    const sendReqWithoutBody = (method, path) => {
      return fetch(path, {method: method}).then(r => r.json()).then(d => console.log(d))
    }
    const post = (path, data) => sendReqWithBody("POST", path, data)
    const put = (path, data) => sendReqWithBody("PUT", path, data)
    const del = (path) => sendReqWithoutBody("DELETE", path)
    const get = (path) => sendReqWithoutBody("GET", path)
    ```
1. Send some requests
   ```
   post("/memos", {content: "Memo#1", shared: false})
   get("/memos")
   ```
