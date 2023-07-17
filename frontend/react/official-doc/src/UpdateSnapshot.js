import { useState } from "react";

export function Snapshot() {
	const [number, setNumber] = useState(0);

	return (
		<div>
			<h1>{number}</h1>
			<button
				onClick={async () => {
					await delay(5000);
					setNumber(number + 1);
					await delay(5000);
					setNumber(10);
				}}
			>
				+3
			</button>
		</div>
	);
}

function delay(ms) {
	return new Promise((resolve) => {
		setTimeout(resolve, ms);
	});
}
