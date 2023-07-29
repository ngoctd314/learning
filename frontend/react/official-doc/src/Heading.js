import { useContext } from "react";
import { LevelContext } from "./App";

export function Heading() {
	const level = useContext(LevelContext);
	return <div>P {level}</div>;
}
