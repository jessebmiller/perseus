# make it a package

import threading, os
from urllib.parse import urlencode
from urllib.request import Request, urlopen

def send(message):
    """ send a message async """
    threading.Thread(target=send_, args=(message,)).start()


def sendf(fmt_string, *args):
    """ format the message first then send that async """
    message = fmt_string.format(*args)
    send(message)


def send_(message):
    """
    post the message to the perseus server

    use the configured host and namespace

    """

    url = "{}{}".format(
        os.environ.get("PERSEUS_HOST", "http://perceus"),
        os.environ.get("PERSEUS_NS", "/default"),
    )
    values = {"message": message}
    request = Request(url, urlencode(values).encode())
    urlopen(request)
