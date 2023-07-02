
temp readme, i'l rewrite it later

### commands

import spells from db
```shell
UNSAVED_SPELL_CREATING=true APP_ENV=develop go run cmd/fill_db.go
```

start server
```shell
APP_ENV=develop go run cmd/main.go
```

swagger documentation: `localhost:8080/docs`

global libs for development: swag, go-migrate

keys

`UNSAVED_SPELL_CREATING=true` - creating spells without any validation

