Plotly Go API
=============

The Go API is heavily based on the work by https://github.com/baruchlubinsky/go-plotly

This package provides wrapper functions for [Plotly](https://plot.ly)'s HTTP API.

Install the library with:

    go get github.com/plotly/go-api/plotly


Authentication
==============

In order to use this package you require API credentials from Plotly. These may
be stored in:

1. `.plotly_credentials.json`
2. `plotly_credentials.json`
3. `$HOME/plotly_credentials.json`
4. `$HOME/.plotly_credentials.json`
5. `/etc/plotly/.plotly_credentials.json`
6. `/etc/plotly/plotly_credentials.json`
7. Environment variables named `PLOTLY_USERNAME` and `PLOTLY_APIKEY`

If more than one of these are available, the highest one in the list takes preference.

The `.json` files should contain the following:

    {"Username":"yourname","Apikey":"yourkey"}


Usage
==

An example program is provided in this repository.

Limitations
==

This is a work in process.

One important thing to be aware of is that the plotly API always returns 200,
so checking for an error from the request is not suitable, rather look at the
`error` field in the response.
