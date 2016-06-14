package.path = package.path .. ";../src/?.lua"

require 'redis'

aliases = require ('config/aliases')
--[[
  print(aliases.apple)
  -- Now we have got the aliases for the database colums
  -- prints `apple`
--]]

params = require ('config/params')
--[[
  print(params.port)
  -- Here we have the parameters as localhost and port
  -- prints `6379`
--]]

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
