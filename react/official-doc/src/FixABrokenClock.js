export function Clock({ time }) {
	let hours = 6;
	let className;
	if (hours >= 0 && hours <= 6) {
		className = "night";
	} else {
		className = "day";
	}

	return <h1 id={className}>{}</h1>;
}
