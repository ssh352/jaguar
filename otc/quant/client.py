

# pip install msgpack-python
import zmq
import time
import msgpack
# 'tcp://10.20.10.187:7000'
# 'tcp://10.20.38.191:7000'
url = 'tcp://10.20.10.187:7000'
ctx = zmq.Context.instance()
req = ctx.socket(zmq.REQ)
req.connect(url)

# newStrategy
r = {'TO': 'PMS', 'FROM': 'MONITOR', 'CMD': 'newStrategy'}
params = ["DeltaHedge", "0.1", "0.003", "1000000", "000333.SZ", "UFX", "1007", "10072"]
r['PARAMS'] = params
req.send(msgpack.packb(r))
msgpack.unpackb(req.recv())


# delRiskRules
# r = {'TO': 'RMS', 'FROM': 'MONITOR', 'CMD': 'delRiskRules'}
# params = ["R0020"]
# r['PARAMS'] = params
# req.send(msgpack.packb(r))
# msgpack.unpackb(req.recv())

# modifyRiskRules
# r = {'TO': 'RMS', 'FROM': 'MONITOR', 'CMD': 'modifyRiskRules'}
# params = ["INSTRUMENTID","000334.SZ","10072","DeltaHedge","trade_pct_vol_30","greater","0.25","1","operator01", "", "R0003"]
# r['PARAMS'] = params
# req.send(msgpack.packb(r))
# msgpack.unpackb(req.recv())

# addRiskRules
# r = {'TO': 'RMS', 'FROM': 'MONITOR', 'CMD': 'addRiskRules'}
# params = ["INSTRUMENTID","000333.SZ","10072","DeltaHedge","trade_pct_vol_30","greater","0.25","1","operator01", ""]
# r['PARAMS'] = params
# req.send(msgpack.packb(r))
# msgpack.unpackb(req.recv())


# getRiskRules
# req.send(msgpack.packb({'TO': 'RMS', 'FROM': 'MONITOR', 'CMD': 'getRiskRules'}))
# ret = msgpack.unpackb(req.recv())
# msgpack.unpackb(ret[b'DAT'])


