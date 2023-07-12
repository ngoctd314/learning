# Managing State

## Reacting to input with state

In React, you don't directly manipulate the UI - meaning you don't enable, disable, show or hide components directly. Instead, you declare what you want to show, and React figures out how to update the UI. Think of getting into a taxi and telling the driver where you want to go instead of telling them exactly where to turn.


## Sharing state between components

- When you want to coordinate two components, move their state to their common parent.
- Then pass the information down through props from their component parent
- Finally, pass the event handlers down so that the children can change the parent's state
- It's useful to consider components as "controlled" (driven by props) or "uncontrolled" (driven by state)