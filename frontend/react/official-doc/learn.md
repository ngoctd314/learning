# Learn React

JS lib for rendering UI.

## Your first component

React applications are built from isolated pieces of UI called components.

Components are core concept of React

React componets are regular JS functions, but their names must start with a capital letter or they won't work!

**Pitfall**

React components are regular JS functions, but their names must start with a capital letter or they won't work!

**Pitfall**

Components can render other components, but you must  never nest their definitions:

```jsx
function Gallery() {
    // Never define a component inside another component!
    // slow and cause bug
    function Profile() {
        return (
        <img
            src="https://i.imgur.com/MK3eW3As.jpg"
            alt="Katherine Johnson"
        />
        )
    }
    return Profile()
}
```

React app begins at a "root" component. It is created automatically when you start a new project.

## Import and Exporting Components 

Two type of export: export, export default

## Writing Markup with JSX

JSX is syntax extension for JS 

JSX: Putting markup into JavaScript

**The Rules of JSX**

1. Return a single root element

JSX looks like HTML, but under the hood it is transformed into plain JS objects. You can't return two objects from a function without wrapping them into an array. This explains why you also can't return two JSX tags without wrapping them into another tag or a Fragment.

2. Close all the tags

3. camelCase most of the things!

JSX turns into JavaScript and attributes written in JSX become keys of JavaScript objects.

## JavaScript in JSX with Curly Braces

You can only use curly braces in two ways inside JSX:

1. As text directly inside a JSX tag: <h1>{name}'s To Do List</h1> works, but <{tag}>Gregorio Y. Zara's To Do List</tag> will not.
2. As attributes immediately following the = sign: src={avatar}

Using "double curlies": CSS and other objects in JSX

In addition to strings, numbers, and other JS expressions, you can even pass objects  in JSX. Objects are also denoted with curly braces, like {name: "NgocTD"}. Therefore, to pass a JS object in JSX, you must wrap the object in another pair of curly braces: person={{name: "NgocTD"}}.

The next time you see {{ and }} in JSX, know that it's nothing more than an object inside the JSX curlies!

## Passing Props to a Component

React components use props to communicate with each other. Every parent component can pass some information to its child components by giving them props. Props might remind you of HTML attributes but you can pass any JavaScript value through them including objects, arrays and functions.

Props are immutable - a term from computer science meaning "unchangeable". When a component needs to change its props (for example, in response to a user interaction or new data), it will have to "ask" its parent component to pass it different props - a new object! Its old props will then be cast aside, and eventually the JS engine will reclaim the memory taken by them.

**Recap**

- To pass props, add them to the JSX, just like you would with HTML attributes.
- To read props, use the function Avatar({ person, size }) destructuring syntax.
- You can specify a default value like size = 100, which is used for missing and undefining props.
- You can forward all props with <Avatar {...props} /> JSX spread syntax, but don't overuse it!
- Nested JSX like <Card><Avatar/></Card> will appear as Card component's children prop.
- Props are read-only snapshots in time: every render receives a new version props.
- You can't change props. When you need interactivity, you'll need to set state.

## Condition Rendering

**Conditional returing nothing with null**

In some situations, you won't want to render anything at all. A component must return something. In this case, you can return null.

```jsx
if (isPacked) {
    return null;
}

return <p>rendered<p/>
```

In practice, returning null from a component isn't common because it might surprise a developer trying to render it. You would conditionally include or exclude the component in the parent component's JSX.

**Logical AND operator (&&)**

U want to render some JSX when the condition is true, or render nothing. With && you could conditionally render the checkmark only is isPacked is true:

```js
return (
    <li>
        {name} {isPacked && "ok"}
    </li>
)
```

A JavaScript && expression returns the value of its right side if the left side is true. React considers false like null or undefined, and doesn't render anything in its place.

**Pitfall**

Don't put numbers on the life side of &&

To test the condition, JS converts the left side to a boolean. However, if the left side is 0, then the whose expression gets that value(0), and React will happily render 0 rather than nothing.

For example

```js
messageCount && <p>New messages</>
```
It's easy to assume that it renders nothing when messageCount is 0, but it really renders the 0 itself!

To fix it, make the left side a boolean 

```js
messageCount > 0 && <p>New messages</p>
```

## Rendering Lists

**Pitfall**

Arrow functions implicitly return the expression right after =>, so you didn't need to a return statement

```js
const listItems = chemists.map(person => <p></p>) // implicit return!
const listItems = chemists.map(person => { // said to have a block body.
// They let you write more than a single line of code
    return <p></p>
})
```

JSX elements directly inside a map() call always need keys!

Keys tell React which array item each component corresponds to, so that is can match them up later. This becomes important if your array items can move (sorting), get inserted, or get deleted. A well chosen key helps React infer what exactly has happend, and make the correct updates to the DOM tree.

Use ```<Fragment>``` is render several DOM nodes

```js
const listItems = people.map(el => {
    <Fragment>
        {el.name}
    </Fragment>
})
```
**Rules of keys**

- Keys must be unique among siblings. However, it's okay to use the same keys for JSX nodes in different arrays.
- Keys must not change or that defeats their purpose?

Your components won't receive key as a prop. It's only used as a hint by React itself. If your component needs an ID, you have to pass it as a separate prop: ```<Profile key={id} useId={id} />```

## Keeping Components Pure


```js
function Cup() {
	// Bad: changing a preexisting variable!
	guest = guest + 1;
	return <h2>Tea cup for guest #{guest}</h2>;
}

export function PureComponent() {
	return (
		<>
			<Cup />
			<Cup />
			<Cup />
		</>
	);
}
```

By passing pros

```js
function Cup({guess}) {
	// Bad: changing a preexisting variable!
	guest = guest + 1;
	return <h2>Tea cup for guest #{guest}</h2>;
}

export function PureComponent() {
	return (
		<>
			<Cup guess={1}/>
			<Cup guess={2}/>
			<Cup guess={3}/>
		</>
	);
}
```

**Detecting impure**

Three kinds of inputs that you can read while rendering: props, state and context. You should always treat these inputs as read only.

React offers a Strict Mode in which it calls each component's function twice during development. By calling the component functions twice, Strict Mode helps find components that break these rules.

Strict Mode has no effect in production, so it won't slow down the app for your users.

## Local mutation

Pure functions don't mutate variables outside of the function's scope or objects that were created before the call - that makes them impure!

However, it's completely fine to change variables and objects that you're just created while rendering.

```js
function Cup({guest}) {
    return <h2>Tea cup for guest #{guest}</h2>
}

export default function TeaGathering() {
    let cups = [];
    for (let i = 1; i <= 12; i++) {
        cups.push(<Cup key={i} guest={i}/>)
    }
    return cups;
}
```

It's completely fine to change variables and objects that you're just created while rendering.

If the cups variable of the [] array were created outside the TeaGathering function, this would be a huge problem! You would be changing a preexisting object by pushing items into that array.

**Where you can cause side effects**

While functional programming relies heavily on purity, at some point, somewhere, something has to change. These changes - updating the screen, starting an animation, changing the data - are called side effects. They're things that happen "on the side", not during rendering.

In React, side effects usually belong inside event handlers. Event handlers are function that React runs when you perform some action - for example, when you click a button. Even though event handlers are defined inside your component, they dont run during rendering! So event handlers don't need to be pure.

If you tried to change any of the array's existing items, you'd have to clone those items too.

It is useful to remember which operations on arrays mutate them, and which don't. For example, push, pop, reverse, and sort will mutate the original array, but slice, filter and map will create a new one.

## Responding to events

React lets you add event handlers to your JSX. Event handlers are your own functions that will be triggered in response to interactions like clicking, hovering, focusing form inputs and so on.

**Pitfall**

Functions passed to event handlers must be passed, not called.

**Event propagation**

Event handlers will also catch events from any children your component might have. We say that an event "bubbles" or "propagates" up the tree: it starts with where the event happened, and then goes up the tree.

```jsx
export function Toolbar() {
	return (
		<div
			onClick={() => alert("You clicked on the toolbar")}
			style={{ height: 300, width: 500, backgroundColor: "#000" }}
		>
			<button onClick={() => alert("Playing!")}>Play Movie</button>
			<button onClick={() => alert("Uploading!")}>Upload image</button>
		</div>
	);
}
```

**Pitfall**

All events propagate in React except onScroll

**Stopping propagation**

Event handlers receive an event object as their only argument. By convention, it's usually called e. That event object also lets you stop the propagation. If you want to prevent an event from reaching parent components, you need to call e.stopPropagation().

**Passing handlers as alternative to propagation**

**Preventing default behavior**

Don't confuse e.stopPropagation() and e.preventDefault(). They are both useful, but are unrelated:

- e.stopPropagation() stops the event handlers attached to the tags above from firing
- e.preventDefault() prevents the default browser behavior from the few events that have it

## State: A Component's Memory

To update a component with new data, two things need to happen:

1. Retain the data between renders
2. Trigger React to render the component with new data (re-rendering)

The useState Hook provides those two things:

1. A state variable to retain the data between renders
2. A state setter function to update the variable and trigger React to render the component again.

## Meet your first Hook

In React, useState, as well as any other function starting with "use" is called a Hook

Hooks are special functions that are only available while React is rendering. They let you "hook into" different React features.

**Pitfall**
You can't call Hooks inside conditions, loops or other nested functions. Hooks are functions, but it's helpful to think of them as unconditional declarations about your component's needs.

**Anotomy of useState**
When you call useState, you are telling React that you want this component to remember something

```jsx
const [index, setIndex] = useState(0)
```

1. Your component is renderd the first time. Because you passed 0 to useState as the initial value for index, it will return [0, setIndex]. React remembers 0 is the latest state value.

2. You update the state. When a user clicks the button, it calls setIndex(index + 1). Index is 0, so it's setIndex(1). This tells React to remember index is 1 now and triggers another render

3. Your component's second render. React still sees useState(0), but because React remembers that you set index to 1, it returns [1, setIndex] instead.

4. And so on!

**How does React know which state to return?**

**State is isolated and private**

State is local to a component instance on the screen. In other words, if you render the same component twice, each copy will have completely isolated state! Changing one of them will not affect the other.

State is fully private to the component declaring it. The parent component can't change it. This lets you add state to any component or remove it without impacting the rest of the components.

**Recap**

- Use a state variable when a component needs to "remember" some information between renders
- State variables are declared by calling the useState hook
- Hooks are special functions that start with use. They let you "hook into" React features like state.
- Hooks might remind you of imports: they need to be called unconditionally, Calling hooks, including useState, is only valid at the top level of a component or another Hook.
- The useState Hook returns a pair of values: the current state and the function to update it
- You can have more than one state variable. Internally, React matches them up by their order
- State is private to the component. If you render it in two places, each copy gets its own state.

**Render and Commit**

Before you components are displayed on screen, they must be rendered by React.

**1. Triggering a render**

There are two reasons for a component to render:

It's the component's initial render. When your app starts, you need to trigger the intial render. Frameworks and sandboxes sometimes hide this code, but it's done by calling createRoot with the target DOM node, and then calling its render method with your component.

The component's (or one of its accesstor's) state has been updated. Once the component has been initially rendered, you can trigger further renders by updating its state with the set function. Updaing your component's state automatically queues a render.

**2. Rendering the component**

After you trigger a render, React calls your components to figure out what to display on screen. 'Rendering" is React calling your component.

- On initial render, React will call the root component
- For subsequent renders, React will call the function component whose state update triggered the render.

**3. Commiting to the DOM**

- For the initial render, React will use the appendChild() DOM API to put all the DOM nodes it has created on screen.
- For re-renders, React will apply the minimal ncessary operations (calculated while rendering) to make the DOM match the latest rendering output.

After rendering is done and React updated the DOM, the browser will repaint the screen.

## State as a Snapshot

State variables might look like regular JS variables that you can read and write to. However, state behaves more like a snapshot. Setting it does not change the state variable you already have, but instead triggers a re-render.

**Rendering takes a snapshot in time**

"Rendering" means that React is calling your component, which is a function.

When React re-renders a component

1. React calls your function again.
2. Your function returns a new JSX snapshot.
3. React then updates the screen to match the snapshot you've returned.

State is not like a regular variable that disappears after your function returns. State actually lives in React itself - outside of your function. When React calls your component, it gives you a snapshot of the state for that particular render. Your component returns a snapshot of the UI with a fresh set of props and event handlers in its JSX, all calculated using the state values from that render.

```jsx
export function Snapshot() {
	const [number, setNumber] = useState(0);

	return (
		<div>
			<h1>{number}</h1>
			<button
				onClick={() => {
					setNumber(number + 1);
					setNumber(number + 1);
					setNumber(number + 1);
				}}
			>
				+3
			</button>
		</div>
	);
}
```

Result: 1

Setting state only changes it for the next render. During the first render, number was 0.

Here is what this button's click handler tells React to do:

1. setNumber(number + 1): number is 0 so setNumber(0 + 1)
   - React prepares to change number to 1 on the next render

2. setNumber(number + 1): number is 0 so setNumber(0 + 1)
   - React prepares to change number to 1 on the next render

3. setNumber(number + 1): number is 0 so setNumber(0 + 1)
   - React prepares to change number to 1 on the next render

Since the number state variable is 0 for this render, its event handler looks like this:

```js
<button onClick={() => {
	setNumber(0 + 1);
	setNumber(0 + 1);
	setNumber(0 + 1);
}}>+3</button>
```

A state variable's value never changes within a render, even if its event handler's code is asynchrous. 

React keeps the state values "fixed" within one render's event handlers. You don't need to worry whether the state has changed while the code is running.

**Recap**

- Setting state requests a new render
- React stores outside of your component, as if on a shelf
- When you call useState, React gives you a snapshot of the state for that render
- Variables and event handlers don't survive re-renders. Every render has its own event handlers

## Queueing a Series of State Updates

Setting a state variable will queue another render. But sometimes you might want to perform multiple operations on the value before queueing the next render. To do this, it helps to understand how React batches state updates.

**React batches state updates**

```js
return (
	<>
		<h1>{number}</h1>
		<button onClick={() => {
			setNumber(number + 1);
			setNumber(number + 1);
			setNumber(number + 1);
		}}>+3</button>
	</>
)
```
React waits until all code in the event handlers has run before processing your state updates. This is why the re-render only happens after all these setNumber() calls. 

This might remind you of a waiter taking an order at the restaurant. A waiter doesn't run to the kitchen at the mention of your first dish! Instead, they let you finish your order, let you make changes to it, and even take orders from other people at the table.

This behavior also known as batching, makes your React app run much faster. It also avoids dealing with confusing "half-finished" renders where only some of the variables have been updated.

**Updating the same state multiple times before the next render**

```js
return (
	<>
		<h1>{number}</h1>
		<button onClick={() => {
			setNumber(n => n + 1);
			setNumber(n => n + 1);
			setNumber(n => n + 1);
		}}>+3</button>
	</>
)
```

Here n => n + 1 is called an updater function. When you pass it to a state setter.

**Naming conventions**

It's common to name the updater function argument by the first letters of the corresponding state variable:

```js
setEnabled(e => !e);
setLastName(ln => ln.reverse());
```

**Recap**

- Setting state does not change the variable in the existing render, but it requests a new render.
- React processes state updates after event handlers have finished running. This is called batching.
- To update some state multiple times in one event,  you can use setNumber(n => n + 1) updater function.