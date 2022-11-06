# arch-hexagonal
A Golang project with Hexagonal Arch.
Using [Cobra-cli](https://github.com/spf13/cobra)

# Input adapter: CLI interface

Creating new product:
```bash
go run main.go cli -a=create -n=Product CLI -p=33.99
```

Getting product by ID:
```bash
go run main.go cli -a=get -i="35c68dee-e321-44cb-b1b1-4703ea71db12"
```

# Input adapter: Webserver
Starting server:
```bash
go run main.go http
```

Getting a product:
```bash
curl http://localhost:8080/product/38092168-d995-40ad-8e96-1419b2a6ddae
```

# Output adapter: SQLite