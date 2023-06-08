export function Signup() {
	return (
		<form
			onSubmit={(e) => {
				alert("Submitting!");
				e.preventDefault();
			}}
		>
			<input />
			<button>Send</button>
		</form>
	);
}
