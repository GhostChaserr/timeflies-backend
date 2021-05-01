

# Todos.


# Making Http post requests
```go
  payload = map[string]string{ "ok": "true" }
  payloadBuf := new(bytes.Buffer)
  json.NewEncoder(payloadBuf).Encode(payload)
  req, _ := http.NewRequest("POST", url, payloadBuf)

  req, _ := http.NewRequest("POST", url, payloadBuf)

  client := &http.Client{}
  res, e := client.Do(req)
  if e != nil {
      return e
  }

  defer res.Body.Close()
```

# Story service functionality.

GET /api/stories
1. Registering new story on behalf of user


POST /api/stories
2. Retrieve stories or single story