# Managing State

## Reacting to input with state

In React, you don't directly manipulate the UI - meaning you don't enable, disable, show or hide components directly. Instead, you declare what you want to show, and React figures out how to update the UI. Think of getting into a taxi and telling the driver where you want to go instead of telling them exactly where to turn.


## Sharing state between components

- When you want to coordinate two components, move their state to their common parent.
- Then pass the information down through props from their component parent
- Finally, pass the event handlers down so that the children can change the parent's state
- It's useful to consider components as "controlled" (driven by props) or "uncontrolled" (driven by state)

## Passing Data Deeply with Context

Usually, you will pass information from a parent component to a child component via props. But passing props can become verbose and inconvenient if you have to pass them through many components in the middle, of if many components in your app need the same information. Context lets the parent component make some information available to any component in the tree below it - no matter how deep - without passing it explicitly through props.

### The problem with passing props

Passing props is a greate way to explicitly pipe data through your UI tree to the components that use it. But passing props can become verbose and inconvenient when you need to pass some prop deeply through the tree, of if many components need the same props.

### Context:an alternative to passing props

Context lets a parent component provide data to the entire tree below it.