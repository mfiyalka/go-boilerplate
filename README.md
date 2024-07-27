
## go-boilerplate

Create a project folder and navigate into it:
```shell
mkdir my-project && cd my-project
```

Clone the repository:
```shell
git clone https://github.com/mfiyalka/go-boilerplate.git .
```

Remove the cloned remote repository:
```shell
git remote remove origin
```

Create .env file:
```shell
cp .env.example .env
```

An example of creating a migration:
```shell
cd migrations && ../bin/goose create create_notes_table sql
```

### Included
* gRPC
* PostgreSQL
* migrations