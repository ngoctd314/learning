# Updating Objects in State

State can hold any kind of JS value, including objects. But you shoudn't change objects that you hold in the React state directly. When you want to update an object, you need to create a new one (or making a copy of an existing one), and then set the state to use that copy.

**You will learn**
- How to correctly update an object in React state
- How to update a nested object without mutating it
- What immutability is, and how not to break it
- How to make object copying less repetitive with Immer

## What's a mutation?

```js
const [x, setX] = useState(0);
```

So far you've been working with numbers, strings and booleans. These kinds of JS values are "immutable", meaning unchangeable or "read only". You can trigger a re-render to replace a value.

```js
setX(5);
```

The x state changed from 0 to 5, but the number 0 ifself did not change. It's not possible to make any changes to build-in primitive values like numbers, strings and booleans in JS.

Now consider an object in state

```js
const [position, setPosition] = useState({x: 0, y: 0})
```

Technically, it is possible to change the contents of the object itself. This is call a mutation

```js
position.x = 5
```

However, although objects in React state are technically mutable, you should treat them as if they were immutable - like numbers, booleans and strings. Instead of mutating them, you should replace them

```js
import {useState} from 'react';
export default function MovingDot() {
    const [position, setPosition] = useState({
        x: 0,
        y: 0
    });
    return (
        <div onPointerMove={e => {
            position.x = e.clientX
            position.y = e.clientY
        }}>
            This is a pointer
        </div>
    )
}
```

The problem with this is bit of code

```js
onPointerMove={e => {
    position.x = e.clientX;
    position.y = e.clientY;
}}
```
React has no idea that object has changed. You should treat the state value you have access to in a render as read-only.

To actually trigger a re-render in this case, create a new object and pass it to the state setting function:
```js
onPointerMove={e => {
    setPosition({
        x: e.clientX
        y: e.clientY
    })
}}
```

With setPosition you're telling React:

- Replace position with this new object
- And render this component again

## Copying objects with the spread syntax