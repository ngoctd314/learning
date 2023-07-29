import { LevelContext } from "./App";

export function Section({ level, children }) {
	return (
		<section className="section">
			<LevelContext.Provider value={level}>{children}</LevelContext.Provider>
		</section>
	);
}
