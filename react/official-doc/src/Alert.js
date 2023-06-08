import { useState } from "react";

export function Alert() {
	const [number, setNumber] = useState(0);

	return (
		<button
			onClick={() => {
				setNumber(number + 5);
				setTimeout(() => {
					alert(number);
				}, 3000);
			}}
		>
			+5
		</button>
	);
}
