let guest = 0;

function Cup() {
	// Bad: changing a preexisting variable!
	guest = guest + 1;
	return <h2>Tea cup for guest #{guest}</h2>;
}

export function PureComponent() {
	return (
		<>
			<Cup />
			<Cup />
			<Cup />
		</>
	);
}
