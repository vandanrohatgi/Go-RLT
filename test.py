import requests
from threading import Thread


url="http://127.0.0.1:8000/python"
count=10100

def make_request():
    requests.get(url=url)

if __name__ == "__main__":
    for i in range(count):
        worker= Thread(target=make_request)
        worker.daemon= True
        worker.start()