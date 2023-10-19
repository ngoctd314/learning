function App() {
  function Profile() {
    return (
      <img src="https://i.imgur.com/MK3eW3As.jpg" alt="Katherine Johnson" />
    );
  }
  function Gallery() {
    return (
      <section>
        <h1>Amazing scientists</h1>
        <Profile />
        <Profile />
        <Profile />
        <Profile />
      </section>
    );
  }
  return <Gallery />;
}

export default App;
