export default function StoryTray({ stories }) {
	return (
		<ul>
			{stories.map((story) => {
				<li key={story.id}>{story.label}</li>;
			})}
		</ul>
	);
}
