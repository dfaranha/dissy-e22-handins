## Your Task

You are given the source code of a simple website that distributes secret information.
However, this secret information is encrypted with AES-CBC under a key you do not know.

Your task is to recover the secret information wihout having access to the encryption key.
This requires you to exploit the properties of the CBC mode of encryption,
together with the fact that the service outputs helpful error messages in case
something goes wrong. A very helpful tutorial to CBC padding oracle attacks can be found [here](https://research.nccgroup.com/2021/02/17/cryptopals-exploiting-cbc-padding-oracles), and additional information can be found on [Wikipedia](https://en.wikipedia.org/wiki/Padding_oracle_attack).

To get started with the task, you can host a local version on your own machine (see below) and use it as training ground.
After you have working code, you can run against the [real target](https://cbc-rsa.netsec22.dk:8000) to recover the secret.

The interaction between your code and the website is done through a cookie `authtoken` exchanged in HTTP requests.
We suggest using the `requests` library in Python (or something similar in your favorite programming language) to simplify the interaction and cookie management. A useful example illustrating the interaction can be found in file `example.py` within this repository.

For the hand-in, you should submit your source code, and a brief report (up to 2 pages) through Brightspace explaining the approach. You are free to choose your programming language for the task, so please including compile and usage instructions.

## Running the Service Locally

With the given files, you can play around with the service and test your code
locally on your own computer.  Note that all secret data has been redacted from
the code and replaced with dummy values.

If you have installed Python 3, you can install the required packages in an
isolated virtual environment:
```
$ python -m venv venv               # (1) create a virtual environment in the directory `venv`
$ . ./venv/bin/activate             # (2) activate the virtual environment
$ pip install -r requirements.txt   # (3) install the required packages into the virtual environment
```
To run the service you then simply execute the following:
```
$ FLASK_APP=main flask run          # (4) run the application
```
The next time you want to run the service, you only need to repeat step (4)
(possibly after activating the virtual environment again Step (4)).

Alternatively, we also prepared a Docker container that you can use:
```
# docker build -t cbc-padding-oracle .
# docker run -p 5000:80 cbc-padding-oracle
```

In both cases, the application is reachable at <http://localhost:5000/>.
