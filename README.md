# rpi-radio
YouTube music streaming for Raspberry Pi

## Radio server
#### Running the server locally
First, move to the server folder `cd radio`. Create a link to the development version of the config file.
```bash
ln -sf dev.env.toml config/env.toml
```
Vendor all modules.
```bash
go mod vendor
```
Compile and run the binary. Some packages are using _cgo_, which means _gcc_ is required for compiling.
```bash
go build
./radio
```
Or simply run with `go run main.go`. After running, the server listens for requests on _localhost:8000_.

The app will be started in development mode. To start the app in production mode, set the environment variable `RADIO_ENV=prod`.
 
#### Cross-compiling and deploying to Raspberry Pi
We are targeting _arm64_ processor architecture. First, build the docker image with all dependencies for cross-compiling.
```bash
make builder
```
Then, using this image, compile the radio server (make sure the modules are vendored).
```bash
make build_server
```
The compiled binary will be outputted to the `artifacts` folder.

Build the client.
```bash
make build_client
```
If you want, you can also compile the server and the client with a single `make build`.

Deploy the app.
```bash
make deploy
```


#### Database
We are using [SQLite](https://github.com/mattn/go-sqlite3) for the database. 

#### Configuration
You can change configurations such as the port or database file name by editing the `config/env.toml` file.
