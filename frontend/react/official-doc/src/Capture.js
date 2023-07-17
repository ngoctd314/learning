export function Capture() {
	return (
		<div
			onClickCapture={() => {
				/** this runs first */
				console.log("run first");
			}}
			// onClick={() => {
			// 	console.log("run first");
			// }}
			style={{ width: 500, height: 500, backgroundColor: "red" }}
		>
			<button
				onClick={(e) => {
					e.stopPropagation();
					console.log("RUN button 1");
				}}
			>
				Button 1
			</button>

			<button
				onClick={(e) => {
					e.stopPropagation();
					console.log("RUN button 2");
				}}
			>
				Button 2
			</button>
		</div>
	);
}
