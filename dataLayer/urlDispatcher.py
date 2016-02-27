#!/usr/bin/python

from models import Firm
from config import urls

print Firm

urlDataAsObjects = []

for newFirm in urls.dataUrlsOfFirms:
    urlDataAsObjects.append(Firm.Builder(newFirm))