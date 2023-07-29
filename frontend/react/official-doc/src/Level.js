import { useContext } from "react";
import { LevelContext } from "./App";

export function Level() {
	const level = useContext(LevelContext);
	return <div>{level}</div>;
}
