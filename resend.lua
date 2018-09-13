local jsons = redis.call('lrange', KEYS[1],0,5) 
local ret={}
local tn=redis.call('time')

for k,v in ipairs(jsons) do
    local obj=cjson.decode(v)
    local ts=obj['ts']
    if tn[1]-ts>0
        then
            redis.call('lrem',KEYS[1],v)
        else
        end
    
end
return ret
