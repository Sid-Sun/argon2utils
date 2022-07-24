### ARGON 2 Utils

Quickly benchmark the performance of a wide range of argon2 configuration on your system and get output in CSV

#### Benchmark 

##### In a UNIX-like environment, simply run:
```bash
git clone https://github.com/Sid-Sun/argon2utils
cd argon2utils
go build -o argon2utils .
./argon2utils > data.csv
column -s, -t < data.csv | less -#2  -S
```

#### Extend benchmarking suite:

Edit `main.go` and add the range of configurations you want to test

