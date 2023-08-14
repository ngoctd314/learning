# Managing state

How your state is organized and how the data flows between your components. Redundant or duplicate state is a common source of bugs.

## Reacting to input with state

In React, you won't modify the UI from code directly. You won't write commands like "disable the button", "enable the button", "show the success message", etc. Intead, you will describe the UI you want to see for the different visual states of your component ("initial state", "typing state", "success state") and then trigger the state changes in response to user input.

React provides a declarative way to manipulate the ui instead of manipulating individual pieces of the ui directly, you describe the different states that your components can be in, and switch between them in response to the user input.

**You will learn**

- How declarative UI programming differs from imperative UI programming
- How to enumerate the different visual states your component can be in
- How to trigger the changes between the different visual states from code

### How declarative UI compares to imperative

When you design UI interactions, you probably think about how the UI changes in response to user actions.

Consider a form that lets the user submit an answer:

- When you type something into the form, the "Submit" button becomes enabled.
- When you press "Submit", both the form and the button become disabled, and a spinner appears.
- If the network request succeeds, the form gets hidden, and the "Thank you" message appears.
- If the network request fails, an error message appears, and the form becomes enabled again.

In imperative programming, the above corresponds directly to how you implement interaction. You have to write the exact instructions to manipulation the UI depending on what just happened.

Here's another way to think about this: imagine riding next to someone in a car and telling them turn by turn where to go.

They don't know where you want to go, they just follow your commands. (And if you get the directions wrong, you end up in the wrong place!). It's called imperative because you have to command each element, from the spinner to the button, telling the computer how to update the UI.

```js
async function handleFormSubmit(e) {
    e.preventDefautl();
    disable(textarea);
    disable(button);
    show(loadingMessage);
    hide(errorMessage);
    try {
        await submitForm(textare.value);
        show(successMessage);
        hide(form);
    } catch(err) {
        show(errorMessage);
        errorMessage.textContext = err.message;
    } finally {
        hide(loadingMessage);
        enable(textarea);
        enable(button);
    }
}
```

Manipulating the UI imperatively works well enough for isolated examples, but it gets exponentially more difficult to manage in more complex systems. Imagine updating a page full of different forms like this one. Adding a new UI element or a new interaction would require carefully checking all existing code to make sure you haven't introduced a bug.

In React, you don't directly manipulate the UI - meaning you dont enable, disable, show or hide components directly. Instead, you declare what you want to show, and React figures out how to update the UI. Think of getting into a taxi and telling the driver where you want to go instead of telling them exactly where to turn. It's the driver's job to get you there, and they might even know some shortcuts you haven't considered!

### Thinking about UI declaratively

You've seen how to implement a form imperatively above.

### Recap

- Declarative programming means describing the UI for each visual state rather than micromanaging the UI (imperative).
- When developing a component:

1. Identify all its visual states
2. Determine the human and computer triggers for state changes
3. Model the state change with useState
4. Remove non-essential state to avoid bugs and paradoxes
5. Connect the event handlers to set state

## Choosing the state structure

Structuring state we well can make a difference between a component that is pleasant to modify and debug, and one that is a constant source of bugs. The most important principle is that state shouldn't contain redundant or duplicated information. If there's unnecessary state, it's easy to forget to update it, and introduce bugs!

For example, this form has a redundant fullName state variable:

```js
import {useState} from 'react';

export default function Form() {
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [fullName, setFullName] = useState('');

    function handleFirstNameChange(e) {
        setFirstName(e.target.value);
        setFullName(e.target.value + ' ' + lastName);
    }

    function handleLastNameChange(e) {
         setLastName(e.target.value);
         setFullName(firstName + "  " + e.target.value);
    }
}
```

**You will learn**

- When to use a single vs multiple state variables
- What to avoid when organizing state
- How to fix common issues with the state structure

### Principles for structuring state

When you write a component that holds some state, you'll have to make choices about how many state variables to use and what the shape their data should be.

- 1. Group related state. If you always update two or more state variables at the same time, consider merging them into a single state variable.
- 2. Avoid contraditions in state. When the state is structured in a way that serveral pieces of state any contradict and disagree with each other, you leave room for mistakes. Try to avoid this.
- 3. Avoid redundant state. If you can calculate some information from the component's props or its existing state variables using during rendering, you should not put that information into that component's state.
- 4. Avoid duplication in state. When the same data is duplicated between multiple state variables, or within nested objects, it is difficult to keep them in sync. Reduce duplication when you can.
- 5. Avoid deeply nested state.

The goal behind these principles is to make state easy to update without introducing mistakes. Removing redundant and duplicate data from state helps ensure that all its pieces stay in sync. Make your state as simple as it can be - but no simpler.

**Pitfall**
If your state variable is an object, remember that you can't update only one field in it without explicity copying the other fields. For example, you can't do setPosition({x: 100}) in the above example because it would not have the y property at all! Instead, if you wanted to set X alone, you would either do setPosition({...position, x: 100}), or split them into two state variables and do setX(100).

**Don't mirror props in state**

```js
function Message({messageColor}) {
     const [color, setColor] = useState(messageColor);
}
```

Here, a color state variable is initialized to the messageColor prop. The problem is that the parent component passes a different value of messageColor later (for example, 'red' instead of 'blue'), the color state variable would not be updated! The state is only initialized during the first render.

"Mirroring" props into state only makes sense when you want to ignore all updated for a specific prop. By convention, start the prop name with initial or default to clarify that its new values are ignored:

```js
function Message({intialColor}) {
    const [color, setColor] = useState(initialColor);
}
```

**Recap**

- If two state variables always update together, consider merging them into one.
- Choose your state variables carefully to avoid creating "impossible" states.
- Structure your state is a way that reduces the chances that you will make a mistake updating it.
- Avoid redundant and duplicate state so that you don't need to keep it in sync.
- Don't put props into state unless you specifically want to prevent updates.
- For UI patterns like selection, keep ID index state instead of the object itself.
- If updating deeply nested state is complicated, try flattening it.
