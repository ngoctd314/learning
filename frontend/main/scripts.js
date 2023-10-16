// import LabelStudio from 'heartexlabs@label-studio@latest';
// import 'heartexlabs@label-studio@latest/build/static/css/main.css';


const labelStudio = new LabelStudio('label-studio', {
  config: `
    <View>
      <Image name="img" value="$image"></Image>
      <RectangleLabels name="tag" toName="img">
        <Label value="Hello"></Label>
        <Label value="World"></Label>
      </RectangleLabels>
    </View>
  `,
  interfaces: [
    "panel",
    "update",
    "controls",
    "side-column",
    "annotations:menu",
    "annotations:add-new",
    "annotations:delete",
    "predictions:menu"
  ],
  user: {
    pk: 1,
    firstName: "James",
    lastName: "Dean"
  },
  task: {
    annotations: [],
    predictions: [],
    id: 1,
    data: {
      image: "https://htx-misc.s3.amazonaws.com/opensource/label-studio/examples/images/nick-owuor-astro-nic-visuals-wDifg5xc9Z4-unsplash.jpg"
    }
  }
});

labelStudio.on("labelStudioLoad", (LS) => {
  // Perform an action when Label Studio is loaded
  const c = LS.annotationStore.addAnnotation({
    userGenerate: true
  });
  LS.annotationStore.selectAnnotation(c.id);
});

labelStudio.on("submitAnnotation", (LS, annotation) => {
  // Retrieve an annotation in JSON format
  console.log(annotation.serializeAnnotation())
});