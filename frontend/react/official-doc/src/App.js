import { createContext } from "react";
import "./App.css";
import { Heading } from "./Heading";
import { Section } from "./Section";

export const LevelContext = createContext(1);

// App root
function App() {
	return (
		<div>
			<h1>Hello</h1>
			<Section level={5}>
				<Heading />
			</Section>
		</div>
	);
}

export default App;
