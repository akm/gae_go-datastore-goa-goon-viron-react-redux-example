# Admin

## Prerequisite

- [Golang](https://golang.org/) ~> 1.9.x
- [Google AppEngine SDK for Go with goapp](https://cloud.google.com/appengine/docs/standard/go/download)
    - Download `the original App Engine SDK for Go` instead of using `gcloud components install app-engine-go`
- [Node.js](https://nodejs.org/)
- [MySQL](https://www.mysql.com/)

See [API server README](../api/README.md) to install Golang and Node.js.

## Setup

```bash
$ cd $(go env GOPATH)/src/github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example
$ cd admin
$ make install
```

Modify `viron_local.mysql.env` to connect your local mysql server.

## Run servers

```bash
$ make dev
```

1. Open https://localhost:8082/
2. Click `+ 追加`
3. Set http://localhost:8081/swagger.json and click `追加`
4. And click the box added
   - You can see the admin UI

## Test APIs in browser console

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
