package.path = package.path .. ";../src/?.lua"

require 'redis'

local params = {
	host = '127.0.0.1',
	port = '6379'
}

local aliases = {
  apple = 'apple',
  facebook = 'facebook',
  google = 'google',
  yahoo = 'yahoo'
}

print(aliases.apple)

local redis = Redis.connect(params)
--redis:select(15) -- for tests

local channels = {
  'data:'..aliases.apple..':stock_market',
  'data:'..aliases.facebook..':stock_market',
  'data:'..aliases.google..':stock_market',
  'data:'..aliases.yahoo..':stock_market'
}

for msg, abort in redis:pubsub({ subscribe = channels }) do
  if msg.kind == 'subscribe' then
    print('Subscribed to channel '..msg.channel)
  elseif msg.kind == 'message' then
    if msg.channel == 'control_channel' then      
      print('Received an command with data: '..msg.payload)
    else
      print('Received the following message from '..msg.channel.."\n  "..msg.payload.."\n")
    end
  end
end
