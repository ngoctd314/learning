type ButtonProps = {
	handleClick: (event: React.MouseEvent<HTMLButtonElement>) => void;
	styles: React.CSSProperties;
};

export const Button = (props: ButtonProps) => {
	return (
		<button
			onClick={(e) => {
				props.handleClick(e);
			}}
			style={props.styles}
		>
			Click
		</button>
	);
};
