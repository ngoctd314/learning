# Describing the UI

React is a JavaScript library for rendering user interfaces (UI).

## Your first component

React applications are built from isolated pieces of UI called components. A React component is a JavaScript function that you can sprinkle with markup. Components can be as small as a button, or as large as an entire page.

```js
function App() {
  return (
    <section>
      <h1>Amazing scientists</h1>
      <Profile />
      <Profile />
      <Profile />
    </section>
  );
}

function Profile() {
  return <img src="https://i.imgur.com/MK3eW3As.jpg" alt="Katherine Johnson" />;
}
```

Components are one of the core concepts of React. They are the foundation upon which you build user interfaces, which makes them perfect place to start your React journey!

**Components: UI building blocks**

React lets you combine your markup, CSS, and JavaScript into custom "components", reusable UI elements for your app.

**Defining a component**

1. Export the component

The export default prefix is a standard JavaScript syntax (not specific to React). It lets you mark the main function in a file so that you can later import it from other files.

2. Define the function

With function Profile() {} you define a JavaScript function with the name Profile

Pitfall: React components are regular JavaScript functions, but their name must start with a capital letter or they won't work!

3. Add markup

The component returns an <img /> tag with src and alt attributes. <img /> is written like HTML, but it is actually JavaScript under the hood! This syntax is called JSX, and it lets you embedded markup inside JavaScript.

Return statements can be written all on one line, as in this component:

```js
return <img src="" alt=""/>
```

But if your markup isn't all on the same line as the return keyword, you must wrap it in a pair of parentheses:

```js
return (
    <div>
        <img src="" alt=""/>
    </div>
)
```

Pitfall: Without parentheses, any code on the lines after return will be ignored.

4. Using a component

Now that you've defined your Profile component, you can nest it inside other components. For example, you can expose a Gallery component that uses multiple Profile components:

```js
function Profile() {
  return <img src="https://i.imgur.com/MK3eW3As.jpg" alt="Katherine Johnson" />;
}

function Gallery() {
  return (
    <section>
      <h1>Amazing scientists</h1>
      <Profile />
      <Profile />
      <Profile />
    </section>
  );
}

export default App;
```

**What the browser sees**

Notice the different in casing

+ <section> is lowercase, so React knows we refer to an HTML tag.
+ <Profile/> starts with a capital P, so React knows that we want to use our component called Profile.

**Nesting and organizing components**

Components are regular JavaScript function, so you can keep multiple components in the same file.

Because the Profile components are rendered inside Gallery - even several times! - We can say that Gallery is a parent component, rendering each Profile as a "child". This is part of the magic of React.

Pitfall: Components can render other components, but you must never nest their definitions:

```js
export default function Gallery() {
    // Never define a component inside another component
    function Profile() {
        // ...
    }
}
```

The snippet above is very slow and causes bugs. Instead, define every component at the top level:

```js
export default function Gallery() {
    // ...
}

function Profile() {
    // ...
}
```

When a child component needs some data from a parent, pass it by props instead of nesting definitions.

