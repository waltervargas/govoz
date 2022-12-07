# govoz
A simple voice chat in Go

## Dependencies
 - portaudio
   - `brew install portaudio`
   - `apt-get install portaudio`

## Installing the cli

``` sh
go install github.com/waltervargas/govoz/cmd/govoz@latest
```

## Usage

``` sh
> govoz # starts the server
```

``` sh
> govoz -c -url "http://192.168.0.1:8080/audio"
```
