function Gallery() {
	return (
		<section>
			<h1>The Gallery</h1>
			<Profile />
			<Profile />
			<Profile />
		</section>
	);
}

function Profile() {
	return <img src="https://i.imgur.com/MK3eW3As.jpg" alt="Katherine Johnson" />;
}

export default Gallery;
