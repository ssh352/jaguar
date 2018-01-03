
import zmq
import time
from _thread import *
url = 'tcp://127.0.0.1:5555'
ctx = zmq.Context.instance()
push = ctx.socket(zmq.PUSH)
push.bind(url)

def senddata(name, delay):
	while True:
		print("sending...")
		push.send((u'%i' % (time.time())).encode('utf8'))
		time.sleep(delay)

start_new_thread(senddata, ("t1", 1, ))