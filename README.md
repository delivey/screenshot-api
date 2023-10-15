# screenshot-api

Simple API built with Go & Fiber to get screenshots from pages.

# Hosting

```
go run main/main.go
```

or

```
cd main && go build main.go
./main
```

# Routes

## `POST /screenshot`

`Content-Type: application/json`

```json
{
    "url": "https://example.org"
}
```

Returns an image with `Content-Type: image/png`
