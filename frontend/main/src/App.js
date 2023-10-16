import LabelStudioReact from "./Preview";

const labelinConfig = `
  <View>
    <Image name="img" value="$image"></Image>
    <RectangleLabels name="tag" toName="img">
      <Label value="Hello"></Label>
      <Label value="World"></Label>
    </RectangleLabels>
  </View>
`;

const task = {
  annotations: [],
  predictions: [],
  id: 1,
  data: {
    image:
      "https://htx-misc.s3.amazonaws.com/opensource/label-studio/examples/images/nick-owuor-astro-nic-visuals-wDifg5xc9Z4-unsplash.jpg",
  },
};

const App = () => {
  return (
    <div class="app-root">
      <LabelStudioReact
        config={labelinConfig}
        task={task}
        interfaces={[
          "panel",
          "update",
          "controls",
          "side-column",
          "annotations:menu",
          "annotations:add-new",
          "annotations:delete",
          "predictions:menu",
        ]}
        user={{
          pk: 1,
          firstName: "James",
          lastName: "Dean",
        }}
      />
    </div>
  );
};

export default App;
