export function Toolbar() {
	return (
		<div
			onClick={() => alert("You clicked on the toolbar")}
			style={{ height: 300, width: 500, backgroundColor: "#000" }}
		>
			<button
				onClick={(e) => {
					alert("Playing!");
					e.stopPropagation();
				}}
			>
				Play Movie
			</button>
			<button onClick={() => alert("Uploading!")}>Upload image</button>
		</div>
	);
}
