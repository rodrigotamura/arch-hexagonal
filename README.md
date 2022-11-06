# arch-hexagonal
A Golang project with Hexagonal Arch.

# Input adapter: CLI interface
Using [Cobra-cli](https://github.com/spf13/cobra)

Creating new product:
```bash
go run main.go cli -a=create -n=Product CLI -p=33.99
```

Getting product by ID:
```bash
go run main.go cli -a=get -i="35c68dee-e321-44cb-b1b1-4703ea71db12"
```

# Output adapter: SQLite