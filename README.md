# mumblestumble
An experimental Mimblewimble implementation written in Golang. Technical documentation can be found [here](doc/README.md).

## Building

```bash
./build.sh
```

## Configuration
Make a copy of `example_config.yml` and call it `config.yml`. Update it to use your public-facing IP address and network of choice (`0` for `mainnet` and `1` for `testnet`).

```yaml
LocalAddress: 123.123.123.123
LocalPort: 25519
Network: 0
```

## Running

```bash
bin/ms
```

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).