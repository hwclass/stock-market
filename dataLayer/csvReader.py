#!/usr/bin/python

import httplib
import csv
from config import urls
import urlDispatcher
import json
import redis

redisClient = redis.StrictRedis(host=urls.redis['host'], port=urls.redis['port'], db=urls.redis['db'])

firmData = urlDispatcher.urlDataAsObjects

conn = httplib.HTTPConnection(urls.stockUrl)

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