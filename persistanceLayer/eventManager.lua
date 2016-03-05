package.path = package.path .. ";../src/?.lua"

require 'redis'

local params = {
	host = '127.0.0.1',
	port = '6379'
}

local redis = Redis.connect(params)
--redis:select(15) -- for tests

local channels = {'control_channel', 'notifications'}

for msg, abort in redis:pubsub({ subscribe = channels }) do
  if msg.kind == 'subscribe' then
    print('Subscribed to channel '..msg.channel)
  elseif msg.kind == 'message' then
    if msg.channel == 'control_channel' then
      if msg.payload == 'quit_loop' then
        print('Aborting pubsub loop...')
        abort()
      else
        print('Received an unrecognized command: '..msg.payload)
      end
    else
      print('Received the following message from '..msg.channel.."\n  "..msg.payload.."\n")
    end
  end
end
