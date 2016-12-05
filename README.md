# perseus

* Centralizes distributed debug messages for killing monsters.
* An all-around badass.
* With a polished shield so your many debug messages don't turn you to stone.

# concept

Percius provides libraries and a server to conviniently wrangle debug messages

The library is essentially an alternative to print statement debugging. Send
debug messages to the server and view them in a browser instead of finding and
watching pring statements.

The server accepts messages sent by the library and puts them into a nice web
page.

# Use

Install a library

    // in go
    import "github.com/oaodev/perseus/libs/go/perseus"

    # for python
    $ pip install git+https://github.com/OAODEV/perseus.git
    $ python
    >>> from pyrseus import send, sendf

and call Send

    // in go
    perseus.Send("message")

    # in python
    >>> send("message")

    # formatted message
    >>> sendf("X is: {}", x)

The default configuration assumes that a perseus server is running on
`http://perseus` and will send messages to the namespace `/default`.

Setting environment variables can change this behavior. For example:

    PERSEUS_URL=http://myperseus.com:2120
    PERSEUS_NS=/my-custom-namespace