class Builder(object):
    def __init__(self, alias=None, url=None):
        self.alias = alias
        self.url = url
    def __getitem__(self, alias): return self.alias[alias]
    def __getitem__(self, url): return self.url[url]