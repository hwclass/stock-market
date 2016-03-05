#!/usr/bin/python

import httplib
import csv
from config import urls
import urlDispatcher
import json
import redis
import sched, time
from datetime import datetime

# generating the redis client
redisClient = redis.StrictRedis(host=urls.redis['host'], port=urls.redis['port'], db=urls.redis['db'])

# the urls to fetch
firmData = urlDispatcher.urlDataAsObjects

# creating the connection for the urls
conn = httplib.HTTPConnection(urls.stockUrl)

# the scheduler reference
s = sched.scheduler(time.time, time.sleep)

# scheduler method
def get_data(sc): 
    print "Getting data..." + datetime.time(datetime.now())
    tempListOfData = []
    for currentFirm in firmData:
    	if currentFirm:
    		conn.request('GET', currentFirm.alias['url'])
    		csvFile = conn.getresponse().read().split()
    		reader = csv.reader(csvFile)
    		for row in reader:
    			if row:
    				#Rows as Date, Open, High, Low, Close, Volume, AdjClose
    				tempListOfData.append(row)
    		redisClient.set(currentFirm.alias['alias'], json.dumps(tempListOfData, separators=(',',':')))
    		tempListOfData = []

    conn.close()
    sc.enter(60, 1, get_data, (sc,))

s.enter(60, 1, get_data, (s,))
s.run()