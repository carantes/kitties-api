# GO Kitties API
Sample restful golang API

## Setup

- Requirements
- Install
- Run
- Tests

## Capabilites

- Gorilla MUX HTTP Router
- GO DotEnv (.env)
- UUID
- ORM


## Resources

GET /api/v1/kitties
```
# Return all kities
# 200 OK

[
    {
        "ID": "904bd6f4-8a13-4bcf-a77b-6f7a3c1b5dfd",
        "Name": "Gaspart",
        "Breed": "British",
        "BirthDate": "2016-07-05",
        "Errors": null
    },
    {
        "ID": "c412d936-a69b-4c11-9d04-327e07d57a4f",
        "Name": "Marcel",
        "Breed": "European",
        "BirthDate": "2014-05-02",
        "Errors": null
    }
]
```

GET /api/v1/kitties/904bd6f4-8a13-4bcf-a77b-6f7a3c1b5dfd
```
# Return a single kitty
# 200 OK

{
    "ID": "904bd6f4-8a13-4bcf-a77b-6f7a3c1b5dfd",
    "Name": "Gaspart",
    "Breed": "British",
    "BirthDate": "2016-07-05",
    "Errors": null
}
```

POST /api/v1/kitties
```
# Create a kitty

#body
{
	"Name": "Chucky",
  	"Breed": "Killer Toy",
    "BirthDate": "2013-07-05"
}

# 200 OK
{
    "ID": "7b23d9a8-3f87-4982-b42e-c1272d8ca1cf"
    "Name": "Chucky",
    "Breed": "Killer Toy",
    "BirthDate": "2013-07-05",
    "Errors": null
}

# 400 Bad Request
# 500 Internal error
```


DELETE /api/v1/kitties/7b23d9a8-3f87-4982-b42e-c1272d8ca1cf
```
# Delete a kitty

# 204 No Content (success)
# 404 Not Found
# 500 Internal error
```