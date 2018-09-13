package mmq

const resend = `
redis.replicate_commands()
local jsons = redis.call('LRANGE', KEYS[2],0,5) 
local ret={}
local tn=redis.call('TIME')
for k,v in ipairs(jsons) do
    local obj=cjson.decode(v)
    local ts=obj['ts']
    if tn[1]-ts>0
        then
            redis.call('RPOPLPUSH',KEYS[1],KEYS[2])
        else
        end
end
return ret
`
