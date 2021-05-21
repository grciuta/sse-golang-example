# SSE Golang Example

This is an example of SSE server and client implementation using Golang.

Main idea - is to implement connections management without using any channels.

You can read more about it [here!]()

## Usage

This example is functional! Usage is very simple

- Create an executable with: `go build`.

- To launch server, call `./sse-golang-example -component=server`. Or just simply `./sse-golang-example`, because default option of component is *server*. From this point application will launch server instance, which listens on `:8000` (this can be changed at `main.go#line:33`) and waits for your input, which will be used as a message for connected clients.

- To launch client, call `./sse-golang-example -component=client` and wait for incomming messages. **Also** If you changed your servers address, respectively change it at `main.go#line:33`.

- Inside `html_listener` you can find an HTML file, which contains simple JS script, to connect to launched SSE server through `EventSource` object. After successful connection, a received message will appear. **Also** If you changed your servers address, respectively change it at `html_listener/index.html#line:9`.