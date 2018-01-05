#!/usr/bin/env python
"""Simple example of publish/subscribe illustrating topics.

Publisher and subscriber can be started in any order, though if publisher
starts first, any messages sent before subscriber starts are lost.  More than
one subscriber can listen, and they can listen to  different topics.

Topic filtering is done simply on the start of the string, e.g. listening to
's' will catch 'sports...' and 'stocks'  while listening to 'w' is enough to
catch 'weather'.
"""

#-----------------------------------------------------------------------------
#  Copyright (c) 2010 Brian Granger, Fernando Perez
#
#  Distributed under the terms of the New BSD License.  The full license is in
#  the file COPYING.BSD, distributed as part of this software.
#-----------------------------------------------------------------------------

import sys
import time
import zmq
import numpy

ctx = zmq.Context()
s = ctx.socket(zmq.SUB)
s.connect("tcp://10.20.38.191:7003")

topics = ["addRiskRules", "delRiskRules", "modifyRiskRules"] 
# manage subscriptions

# print("Receiving messages on ALL topics...")
# s.setsockopt_string(zmq.SUBSCRIBE,'')

print("Receiving messages on topics: %s ..." % topics)
for t in topics:
	s.setsockopt_string(zmq.SUBSCRIBE,t)
	

while True:
	topic, msg = s.recv_multipart()
	print('   Topic: %s, msg:%s' % (topic, msg))


