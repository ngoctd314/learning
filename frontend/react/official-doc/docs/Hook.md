# React Hook

## useMemo

useMemo is a React Hook that lets you cache the result of a calculation between re-render

```js
const cachedValue = useMemo(calculateValue, dependencies)
```

**Usage**

- Skipping expensive recalculation
- Skipping re-rendering