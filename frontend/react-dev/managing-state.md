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

In imperative programming, the above corresponds directly to how you implement interaction. You have to write the exact instructions to manipulation the UI depending on what just happended.

Here's another way to think about this: imagine riding next to someone in a car and telling them turn by turn where to go.

They don't know where you want to go, they just follow your commands. (And if you get the directions wrong, you end up in the wrong place!). It's called imperative because you have to command each element, from the spinner to the button, telling the computer how to update the uI.

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
