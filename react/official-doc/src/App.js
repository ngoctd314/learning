import "./App.css";
import { Form } from "./Form";
import { TrafficLight } from "./TrafficLight";
import { Snapshot } from "./UpdateSnapshot";

// App root
function App() {
	return (
		<>
			<Snapshot />
			<Form />
			<TrafficLight />
		</>
	);
}

export default App;
