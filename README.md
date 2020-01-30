# JWT Parser
## Usage
`go get github.com/crwgregory/jwt-parser`

```
jwt-parser eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImp0aSI6IjE2NTg2NWVhLTM5MzEtNGRlYi1hMWRmLWI0ODRhMjYxMDllNSIsImlhdCI6MTU4MDQwNjIwOSwiZXhwIjoxNTgwNDA5ODA5fQ.xbLq9kKUqH1zLGfQ5luqlyMbjo9gJsrOEdcerqIr6so

[
    {
        "alg": "HS256",
        "typ": "JWT"
    },
    {
        "admin": true,
        "exp": 1580409809,
        "iat": 1580406209,
        "jti": "165865ea-3931-4deb-a1df-b484a26109e5",
        "name": "John Doe",
        "sub": "1234567890"
    }
]
```