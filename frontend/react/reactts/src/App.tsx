import { useState } from "react";
import "./App.css";

function App() {
	const [answer, setAnswer] = useState("");
	const [error, setError] = useState(null);
	const [status, setStatus] = useState("typing");

	if (status === "success") {
		return <h1>That's right</h1>;
	}

	return <div className="App"></div>;
}

export default App;
