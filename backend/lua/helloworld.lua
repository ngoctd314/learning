num = 42 -- All numbers
print(num)

s = 'walternate' -- Immutable strings
t = "double quotes are also fine"
u = [[ Double brakets
    start and end
    multi-line strings.]]
t = nil -- Undefines t; Lua has garbage collection

while num < 50 do
    num = num + 1 -- No ++ or += type operators.
end

if num > 40 then
    print('over 40')
elseif s ~= 'walternate' then -- ~= is not equals
    io.write('not over 40\n')
else
    -- variables are global by default.
    thisIsGlobal = 5                    -- camelCase is common
    local line = io.Read()              -- Reads next stdin line.
    print('Writer is coming, ' .. line) -- string concatenation uses the .. operator
end

-- undefined variables return nil
-- this is not an error
foo = anUnknownVariable -- now foo = nil
aBoolValue = false

-- only nil and false are falsy; 0 and '' are true!
if not aBoolValue then print('it was false') end

-- 'or' and 'and' are short circuited
-- this is similar to the a?b:c operator in C/js:
ans = aBoolValue and 'yes' or 'no' --> 'no'

karlSum = 0
for i = 1, 100 do -- the range includes both ends
    karlSum = karlSum + i
end

-- use 100, 1, -1  as the range to count down
fredSum = 0
for j = 100, 1, -1 do fredSum = fredSum + j end

-- in general, the range is begin, end, [, step]

-- Another loop construct:
repeat
    print('the way of the structure')
    num = num - 1
until num == 0
